package storeapi

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	store "store/gen/store"
)

type StoreTestSuite struct {
	suite.Suite
	DB   *sql.DB
	mock sqlmock.Sqlmock
	s    store.Service
}

func stringPtr(s string) *string {
	return &s
}
func (suite *StoreTestSuite) SetupTest() {
	var err error
	suite.DB, suite.mock, err = sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.s = NewStore(suite.DB)
}

func (suite *StoreTestSuite) TearDownTest() {
	suite.DB.Close()
}

func TestStoreTestSuite(t *testing.T) {
	suite.Run(t, new(StoreTestSuite))
}

func (suite *StoreTestSuite) TestCreateUser() {
	testCases := []struct {
		name          string
		input         *store.NewUser
		expectedUser  *store.User
		expectedError error
		mockSetup     func()
	}{
		{
			name: "Successful user creation",
			input: &store.NewUser{
				Username:  "testuser",
				Email:     "test@example.com",
				FirstName: "Test",
				LastName:  "User",
				Password:  "password123",
			},
			expectedUser: &store.User{
				ID:        "mocked-uuid",
				Username:  "testuser",
				Email:     "test@example.com",
				FirstName: stringPtr("Test"),
				LastName:  stringPtr("User"),
			},
			expectedError: nil,
			mockSetup: func() {
				suite.mock.ExpectQuery("INSERT INTO users").WithArgs(
					sqlmock.AnyArg(),
					"testuser",
					"test@example.com",
					"Test",
					"User",
					sqlmock.AnyArg(),
				).WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email", "first_name", "last_name"}).
					AddRow("mocked-uuid", "testuser", "test@example.com", "Test", "User"))
			},
		},
		// Add more test cases here
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			tc.mockSetup()

			user, err := suite.s.CreateUser(context.Background(), tc.input)

			if tc.expectedError != nil {
				assert.EqualError(suite.T(), err, tc.expectedError.Error())
			} else {
				assert.NoError(suite.T(), err)
			}

			assert.Equal(suite.T(), tc.expectedUser, user)
			assert.NoError(suite.T(), suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *StoreTestSuite) TestLoginUser() {
	testCases := []struct {
		name          string
		input         *store.LoginUserPayload
		expectedToken string
		expectedError error
		mockSetup     func()
	}{
		{
			name: "Successful login",
			input: &store.LoginUserPayload{
				Username: "testuser",
				Password: "password123",
			},
			expectedToken: "mocked-jwt-token",
			expectedError: nil,
			mockSetup: func() {
				suite.mock.ExpectQuery("SELECT password FROM users WHERE username = ?").
					WithArgs("testuser").
					WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow("$2a$10$hashedpassword"))
			},
		},
		// Add more test cases here
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			tc.mockSetup()

			result, err := suite.s.LoginUser(context.Background(), tc.input)

			if tc.expectedError != nil {
				assert.EqualError(suite.T(), err, tc.expectedError.Error())
			} else {
				assert.NoError(suite.T(), err)
			}

			if tc.expectedToken != "" {
				assert.NotNil(suite.T(), result)
				assert.NotNil(suite.T(), result.Token)
				assert.Equal(suite.T(), tc.expectedToken, *result.Token)
			}

			assert.NoError(suite.T(), suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *StoreTestSuite) TestCreateProduct() {
	testCases := []struct {
		name            string
		input           *store.NewProduct
		expectedProduct *store.Product
		expectedError   error
		mockSetup       func()
	}{
		{
			name: "Successful product creation",
			input: &store.NewProduct{
				Name:        "Test Product",
				Description: stringPtr("This is a test product"),
				Price:       9.99,
				Inventory:   100,
			},
			expectedProduct: &store.Product{
				ID:          "mocked-product-uuid",
				Name:        "Test Product",
				Description: stringPtr("This is a test product"),
				Price:       9.99,
				Inventory:   100,
				UserID:      "mocked-user-uuid",
			},
			expectedError: nil,
			mockSetup: func() {
				suite.mock.ExpectBegin()
				suite.mock.ExpectQuery("SELECT id FROM users WHERE username = ?").
					WithArgs("testuser").
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("mocked-user-uuid"))
				suite.mock.ExpectQuery("INSERT INTO products").
					WithArgs(sqlmock.AnyArg(), "Test Product", "This is a test product", 9.99, 100, "mocked-user-uuid").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "price", "inventory", "userID"}).
						AddRow("mocked-product-uuid", "Test Product", "This is a test product", 9.99, 100, "mocked-user-uuid"))
				suite.mock.ExpectCommit()
			},
		},
		// Add more test cases here
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			tc.mockSetup()

			ctx := context.WithValue(context.Background(), "username", "testuser")
			product, err := suite.s.CreateProduct(ctx, tc.input)

			if tc.expectedError != nil {
				assert.EqualError(suite.T(), err, tc.expectedError.Error())
			} else {
				assert.NoError(suite.T(), err)
			}

			assert.Equal(suite.T(), tc.expectedProduct, product)
			assert.NoError(suite.T(), suite.mock.ExpectationsWereMet())
		})
	}
}

// Add more test methods for other API endpoints here

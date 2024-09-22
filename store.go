package storeapi

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	store "store/gen/store"
	"store/jwthelper"
)

type CustomError struct {
	StatusCode int
	Message    string
}

func (e *CustomError) Error() string {
	return e.Message
}

// NewForbiddenError creates a new CustomError with StatusForbidden
func NewForbiddenError(message string) *CustomError {
	return &CustomError{
		StatusCode: 403,
		Message:    message,
	}
}

type storesrvc struct {
	db *sql.DB
}

func NewStore(db *sql.DB) store.Service {
	return &storesrvc{db: db}
}

func (s *storesrvc) CreateUser(ctx context.Context, p *store.NewUser) (res *store.User, err error) {
	id := uuid.New().String()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %v", err)
	}

	query := `INSERT INTO users (id, username, email, first_name, last_name, password) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, username, email, first_name, last_name`
	user := &store.User{}
	err = s.db.QueryRowContext(ctx, query, id, p.Username, p.Email, p.FirstName, p.LastName, string(hashedPassword)).Scan(&user.ID, &user.Username, &user.Email, &user.FirstName, &user.LastName)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}

	return user, nil
}

func (s *storesrvc) LoginUser(ctx context.Context, p *store.LoginUserPayload) (res *store.LoginUserResult, err error) {
	var hashedPassword string
	err = s.db.QueryRowContext(ctx, "SELECT password FROM users WHERE username = $1", p.Username).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("invalid username or password")
		}
		return nil, fmt.Errorf("database error: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(p.Password))
	if err != nil {
		return nil, fmt.Errorf("invalid username or password")
	}

	token, err := jwthelper.GenerateJWT(p.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}

	return &store.LoginUserResult{Token: &token}, nil
}

func (s *storesrvc) GetUser(ctx context.Context, p *store.GetUserPayload) (res *store.User, err error) {
	query := `SELECT id, username, email, first_name, last_name FROM users WHERE id = $1`
	user := &store.User{}
	err = s.db.QueryRowContext(ctx, query, p.ID).Scan(&user.ID, &user.Username, &user.Email, &user.FirstName, &user.LastName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.MakeNotFound(fmt.Errorf("user not found"))
		}
		return nil, fmt.Errorf("error getting user: %v", err)
	}
	return user, nil
}

func (s *storesrvc) GetUserAll(ctx context.Context) (res []*store.User, err error) {
	query := `SELECT id, username, email, first_name, last_name FROM users`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error getting users: %v", err)
	}
	defer rows.Close()

	var users []*store.User
	for rows.Next() {
		user := &store.User{}
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.FirstName, &user.LastName); err != nil {
			return nil, fmt.Errorf("error scanning user: %v", err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (s *storesrvc) UpdateUser(ctx context.Context, p *store.UserUpdatePayload) (res *store.User, err error) {
	username, ok := ctx.Value("username").(string)
	if !ok {
		return nil, fmt.Errorf("unauthorized: missing username in context")
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error starting transaction: %v", err)
	}
	defer tx.Rollback()

	var userID string
	err = tx.QueryRowContext(ctx, "SELECT id FROM users WHERE username = $1", username).Scan(&userID)
	if err != nil {
		return nil, fmt.Errorf("error getting user ID: %v", err)
	}

	query := `
		UPDATE users
		SET email = $1, first_name = $2, last_name = $3
		WHERE id = $4
		RETURNING id, username, email, first_name, last_name;
	`

	updatedUser := &store.User{}
	err = tx.QueryRowContext(ctx, query, p.Email, p.FirstName, p.LastName, userID).Scan(
		&updatedUser.ID, &updatedUser.Username, &updatedUser.Email, &updatedUser.FirstName, &updatedUser.LastName)
	if err != nil {
		return nil, fmt.Errorf("error updating user: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("error committing transaction: %v", err)
	}

	return updatedUser, nil
}

func (s *storesrvc) DeleteUser(ctx context.Context) (err error) {
	username, ok := ctx.Value("username").(string)
	if !ok {
		return fmt.Errorf("unauthorized: missing username in context")
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var userID string
	err = tx.QueryRowContext(ctx, "SELECT id FROM users WHERE username = $1", username).Scan(&userID)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil
}

func (s *storesrvc) CreateProduct(ctx context.Context, p *store.NewProduct) (res *store.Product, err error) {
	username, ok := ctx.Value("username").(string)
	if !ok {
		return nil, fmt.Errorf("unauthorized: missing username in context")
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error starting transaction: %v", err)
	}
	defer tx.Rollback()

	var userID string
	err = tx.QueryRowContext(ctx, "SELECT id FROM users WHERE username = $1", username).Scan(&userID)
	if err != nil {
		return nil, fmt.Errorf("error getting user ID: %v", err)
	}

	id := uuid.New().String()
	query := `INSERT INTO products (id, name, description, price, inventory, userid) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, description, price, inventory, userID`

	product := &store.Product{}
	err = s.db.QueryRowContext(ctx, query, id, p.Name, p.Description, p.Price, p.Inventory, userID).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Inventory, &product.UserID)
	if err != nil {
		return nil, fmt.Errorf("error creating product: %v", err)
	}

	return product, nil
}

func (s *storesrvc) GetProduct(ctx context.Context, p *store.GetProductPayload) (res *store.Product, err error) {
	query := `SELECT id, name, description, price, inventory FROM products WHERE id = $1`

	product := &store.Product{}
	err = s.db.QueryRowContext(ctx, query, p.ID).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Inventory)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.MakeNotFound(fmt.Errorf("product not found"))
		}
		return nil, fmt.Errorf("error getting product: %v", err)
	}

	return product, nil
}

func (s *storesrvc) ListProducts(ctx context.Context) (res []*store.Product, err error) {
	query := `SELECT id, name, description, price, inventory,userid FROM products`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error listing products: %v", err)
	}
	defer rows.Close()

	var products []*store.Product
	for rows.Next() {
		product := &store.Product{}
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Inventory, &product.UserID); err != nil {
			return nil, fmt.Errorf("error scanning product: %v", err)
		}
		products = append(products, product)
	}

	return products, nil
}

func (s *storesrvc) AddToCart(ctx context.Context, p *store.CartItem) (res *store.Cart, err error) {
	username, ok := ctx.Value("username").(string)
	if !ok {
		return nil, fmt.Errorf("unauthorized: missing username in context")
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error starting transaction: %v", err)
	}
	defer tx.Rollback()

	var userID string
	err = tx.QueryRowContext(ctx, "SELECT id FROM users WHERE username = $1", username).Scan(&userID)
	if err != nil {
		return nil, fmt.Errorf("error getting user ID: %v", err)
	}

	// Check if the cart exists, if not create a new one
	var cartID string
	err = tx.QueryRowContext(ctx, "SELECT id FROM carts WHERE user_id = $1", userID).Scan(&cartID)
	if err == sql.ErrNoRows {
		cartID = uuid.New().String()
		_, err = tx.ExecContext(ctx, "INSERT INTO carts (id, user_id, total_amount) VALUES ($1, $2, $3)", cartID, userID, 0)
		if err != nil {
			return nil, fmt.Errorf("error creating cart: %v", err)
		}
	} else if err != nil {
		return nil, fmt.Errorf("error checking for existing cart: %v", err)
	}

	// Get the price of the product
	var price float64
	err = tx.QueryRowContext(ctx, "SELECT price FROM products WHERE id = $1", p.ProductID).Scan(&price)
	if err != nil {
		return nil, fmt.Errorf("error getting product price: %v", err)
	}

	// Add or update the item in the cart
	_, err = tx.ExecContext(ctx, `
		INSERT INTO cart_items (cart_id, product_id, quantity)
		VALUES ($1, $2, $3)
		ON CONFLICT (cart_id, product_id) DO UPDATE SET quantity = cart_items.quantity + $3
	`, cartID, p.ProductID, p.Quantity)
	if err != nil {
		return nil, fmt.Errorf("error adding item to cart: %v", err)
	}

	// Update the total amount of the cart
	_, err = tx.ExecContext(ctx, `
		UPDATE carts
		SET total_amount = total_amount + $1
		WHERE id = $2
	`, price*float64(p.Quantity), cartID)
	if err != nil {
		return nil, fmt.Errorf("error updating cart total amount: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("error committing transaction: %v", err)
	}

	return s.getCart(ctx, userID)
}
func (s *storesrvc) RemoveFromCart(ctx context.Context, p *store.RemoveFromCartPayload) (res *store.Cart, err error) {
	username, ok := ctx.Value("username").(string)
	if !ok {
		return nil, fmt.Errorf("unauthorized: missing username in context")
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error starting transaction: %v", err)
	}
	defer tx.Rollback()

	var userID string
	err = tx.QueryRowContext(ctx, "SELECT id FROM users WHERE username = $1", username).Scan(&userID)
	if err != nil {
		return nil, fmt.Errorf("error getting user ID: %v", err)
	}

	var cartID string
	err = tx.QueryRowContext(ctx, "SELECT id FROM carts WHERE user_id = $1", userID).Scan(&cartID)
	if err != nil {
		return nil, fmt.Errorf("error getting cart: %v", err)
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM cart_items WHERE cart_id = $1 AND product_id = $2", cartID, p.ProductID)
	if err != nil {
		return nil, fmt.Errorf("error removing item from cart: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("error committing transaction: %v", err)
	}

	return s.getCart(ctx, userID)
}

func (s *storesrvc) GetCart(ctx context.Context) (res *store.Cart, err error) {
	username, ok := ctx.Value("username").(string)
	if !ok {
		return nil, fmt.Errorf("unauthorized: missing username in context")
	}

	var userID string
	err = s.db.QueryRowContext(ctx, "SELECT id FROM users WHERE username = $1", username).Scan(&userID)
	if err != nil {
		return nil, fmt.Errorf("error getting user ID: %v", err)
	}

	return s.getCart(ctx, userID)
}

func (s *storesrvc) getCart(ctx context.Context, userID string) (*store.Cart, error) {
	query := `
		SELECT c.id, c.user_id, ci.product_id, ci.quantity, p.price
		FROM carts c
		LEFT JOIN cart_items ci ON c.id = ci.cart_id
		LEFT JOIN products p ON ci.product_id = p.id
		WHERE c.user_id = $1
	`

	rows, err := s.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("error getting cart: %v", err)
	}
	defer rows.Close()

	cart := &store.Cart{UserID: userID, Items: []*store.CartItem{}, TotalAmount: 0}
	for rows.Next() {
		var productID sql.NullString
		var quantity sql.NullInt32
		var price sql.NullFloat64
		err = rows.Scan(&cart.ID, &cart.UserID, &productID, &quantity, &price)
		if err != nil {
			return nil, fmt.Errorf("error scanning cart item: %v", err)
		}

		if productID.Valid && quantity.Valid {
			item := &store.CartItem{
				ProductID: productID.String,
				Quantity:  int(quantity.Int32),
			}
			cart.Items = append(cart.Items, item)
			if price.Valid {
				cart.TotalAmount += price.Float64 * float64(quantity.Int32)
			}
		}
	}

	if cart.ID == "" {
		return nil, store.MakeNotFound(fmt.Errorf("cart not found"))
	}

	return cart, nil
}
func (s *storesrvc) CreateOrder(ctx context.Context) (res *store.Order, err error) {
	username, ok := ctx.Value("username").(string)
	if !ok {
		return nil, fmt.Errorf("unauthorized: missing username in context")
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error starting transaction: %v", err)
	}
	defer tx.Rollback()

	var userID string
	err = tx.QueryRowContext(ctx, "SELECT id FROM users WHERE username = $1", username).Scan(&userID)
	if err != nil {
		return nil, fmt.Errorf("error getting user ID: %v", err)
	}

	cart, err := s.getCart(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("error getting cart: %v", err)
	}

	if len(cart.Items) == 0 {
		return nil, fmt.Errorf("cart is empty")
	}
	const defaultStatus = "pending"
	orderID := uuid.New().String()
	_, err = tx.ExecContext(ctx, "INSERT INTO orders (id, user_id, total_amount, overall_status) VALUES ($1, $2, $3, $4)",
		orderID, userID, cart.TotalAmount, defaultStatus)
	if err != nil {
		return nil, fmt.Errorf("error creating order: %v", err)
	}

	for _, item := range cart.Items {
		var price float64
		var sellerID string
		err = tx.QueryRowContext(ctx, "SELECT price, userid FROM products WHERE id = $1", item.ProductID).Scan(&price, &sellerID)
		if err != nil {
			return nil, fmt.Errorf("error getting product price and seller: %v", err)
		}

		_, err = tx.ExecContext(ctx, "INSERT INTO order_items (order_id, product_id, seller_id, quantity, price, status) VALUES ($1, $2, $3, $4, $5, $6)",
			orderID, item.ProductID, sellerID, item.Quantity, price, defaultStatus)
		if err != nil {
			return nil, fmt.Errorf("error adding order item: %v", err)
		}

		_, err = tx.ExecContext(ctx, "UPDATE products SET inventory = inventory - $1 WHERE id = $2", item.Quantity, item.ProductID)
		if err != nil {
			return nil, fmt.Errorf("error updating product inventory: %v", err)
		}
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM cart_items WHERE cart_id = $1", cart.ID)
	if err != nil {
		return nil, fmt.Errorf("error clearing cart: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("error committing transaction: %v", err)
	}

	return s.GetOrder(ctx, &store.GetOrderPayload{ID: orderID})
}

func (s *storesrvc) GetOrder(ctx context.Context, p *store.GetOrderPayload) (res *store.Order, err error) {
	query := `
		SELECT o.id, o.user_id, o.total_amount, o.overall_status, 
       oi.product_id, oi.seller_id, oi.quantity, oi.price, oi.status
FROM orders o
 LEFT JOIN order_items oi ON o.id = oi.order_id
WHERE o.id = $1`

	rows, err := s.db.QueryContext(ctx, query, p.ID)
	if err != nil {
		return nil, fmt.Errorf("error getting order: %v", err)
	}
	defer rows.Close()

	var order *store.Order
	for rows.Next() {
		if order == nil {
			order = &store.Order{Items: []*store.OrderItem{}}
			var productID, sellerID, itemStatus sql.NullString
			var quantity sql.NullInt32
			var price sql.NullFloat64
			err = rows.Scan(&order.ID, &order.UserID, &order.TotalAmount, &order.OverallStatus,
				&productID, &sellerID, &quantity, &price, &itemStatus)

			if err != nil {
				return nil, fmt.Errorf("error scanning order: %v", err)
			}
			if productID.Valid {
				item := &store.OrderItem{
					ProductID: productID.String,
					SellerID:  sellerID.String,
					Quantity:  int(quantity.Int32),
					Price:     price.Float64,
					Status:    itemStatus.String,
				}
				order.Items = append(order.Items, item)
			}
		} else {
			var productID, sellerID, itemStatus sql.NullString
			var quantity sql.NullInt32
			var price sql.NullFloat64
			err = rows.Scan(new(sql.NullString), new(sql.NullString), new(sql.NullFloat64), new(sql.NullString),
				&productID, &sellerID, &quantity, &price, &itemStatus)
			if err != nil {
				return nil, fmt.Errorf("error scanning order item: %v", err)
			}
			if productID.Valid {
				item := &store.OrderItem{
					ProductID: productID.String,
					SellerID:  sellerID.String,
					Quantity:  int(quantity.Int32),
					Price:     price.Float64,
					Status:    itemStatus.String,
				}
				order.Items = append(order.Items, item)
			}
		}
	}

	if order == nil {
		return nil, store.MakeNotFound(fmt.Errorf("order not found"))
	}

	return order, nil
}

func (s *storesrvc) GetUserOrders(ctx context.Context) (res []*store.Order, err error) {
	username, ok := ctx.Value("username").(string)
	if !ok {
		return nil, fmt.Errorf("unauthorized: missing username in context")
	}

	var userID string
	err = s.db.QueryRowContext(ctx, "SELECT id FROM users WHERE username = $1", username).Scan(&userID)
	if err != nil {
		return nil, fmt.Errorf("error getting user ID: %v", err)
	}

	query := `
		SELECT o.id, o.user_id, o.total_amount, o.overall_status, 
			   oi.product_id, oi.quantity, oi.price
		FROM orders o
		LEFT JOIN order_items oi ON o.id = oi.order_id
		WHERE o.user_id = $1
		ORDER BY o.id`

	rows, err := s.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("error getting user orders: %v", err)
	}
	defer rows.Close()

	orders := make(map[string]*store.Order)
	for rows.Next() {
		var orderID, userID, productID sql.NullString
		var totalAmount, price sql.NullFloat64
		var status sql.NullString
		var quantity sql.NullInt32

		err = rows.Scan(&orderID, &userID, &totalAmount, &status, &productID, &quantity, &price)
		if err != nil {
			return nil, fmt.Errorf("error scanning order: %v", err)
		}

		order, exists := orders[orderID.String]
		if !exists {
			order = &store.Order{
				ID:            orderID.String,
				UserID:        userID.String,
				TotalAmount:   totalAmount.Float64,
				OverallStatus: status.String,
				Items:         []*store.OrderItem{},
			} //making a new order to be able to see what products are there in the order
			orders[orderID.String] = order
		}

		if productID.Valid {
			item := &store.OrderItem{
				ProductID: productID.String,
				Quantity:  int(quantity.Int32),
				Price:     price.Float64,
			}
			order.Items = append(order.Items, item)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	var result []*store.Order
	for _, eachOrder := range orders {
		result = append(result, eachOrder)
	}

	if len(result) == 0 {
		return nil, store.MakeNotFound(fmt.Errorf("no orders found for user"))
	}

	return result, nil
}

func (s *storesrvc) DeleteOrder(ctx context.Context, p *store.DeleteOrderPayload) (err error) {
	username, ok := ctx.Value("username").(string)
	if !ok {
		return fmt.Errorf("unauthorized: missing username in context")
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error starting transaction: %v", err)
	}
	defer tx.Rollback()

	var userID string
	err = tx.QueryRowContext(ctx, "SELECT id FROM users WHERE username = $1", username).Scan(&userID)
	if err != nil {
		return fmt.Errorf("error getting user ID: %v", err)
	}
	query := `DELETE FROM order_items where order_id = $1`
	_, err = tx.ExecContext(ctx, query, p.ID)
	if err != nil {
		return fmt.Errorf("error deleting an item in the order: %v", err)
	}
	query = `DELETE FROM orders WHERE id = $1`
	_, err = tx.ExecContext(ctx, query, p.ID)
	if err != nil {
		return fmt.Errorf("error deleting order: %v", err)
	}
	return tx.Commit()

	return nil
}

func (s *storesrvc) GetProductsPostedByUser(ctx context.Context) (res []*store.Product, err error) {
	username, ok := ctx.Value("username").(string)
	if !ok {
		return nil, fmt.Errorf("unauthorized: missing username in context")
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error starting transaction: %v", err)
	}
	defer tx.Rollback()

	var userID string
	err = tx.QueryRowContext(ctx, "SELECT id FROM users WHERE username = $1", username).Scan(&userID)
	if err != nil {
		return nil, fmt.Errorf("error getting user ID: %v", err)
	}
	query := `SELECT name, description, price, inventory,userid FROM products WHERE userid = $1`
	rows, err := s.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("error getting products: %v", err)
	}
	defer rows.Close()
	var products []*store.Product

	//for rows.Next() {
	//	var productID, productName, productDescription, productPrice, productInventory, productQuantity sql.NullString
	//	err = rows.Scan(&productName, &productDescription, &productPrice, &productInventory, &productQuantity)
	//	if err != nil {
	//		return nil, fmt.Errorf("error scanning product: %v", err)
	//	}
	//	product := &store.Product{
	//		Name:        productName.String,
	//		Description: &productDescription,
	//		Inventory:   productInventory,
	//		Quantity:    productQuantity.String,
	//	}
	//	products[productID.String] = product
	//	products[productID.String] = product
	//}
	for rows.Next() {
		var product store.Product
		var description sql.NullString
		err = rows.Scan(&product.Name, &product.Description, &product.Price, &product.Inventory, &product.UserID)
		if err != nil {
			return nil, fmt.Errorf("error scanning product: %v", err)
		}
		if description.Valid {
			product.Description = &description.String
		}
		products = append(products, &product)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return products, nil
	//return products
}

func (s *storesrvc) UpdateOrderItemStatus(ctx context.Context, p *store.UpdateOrderItemStatusPayload) (res *store.Order, err error) {
	username, ok := ctx.Value("username").(string)
	if !ok {
		return nil, fmt.Errorf("unauthorized: missing username in context")
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error starting transaction: %v", err)
	}
	defer tx.Rollback()

	var userID string
	err = tx.QueryRowContext(ctx, "SELECT id FROM users WHERE username = $1", username).Scan(&userID)
	if err != nil {
		return nil, fmt.Errorf("error getting user ID: %v", err)
	}

	// Check if the user is the seller of the product
	var sellerID string
	err = tx.QueryRowContext(ctx, "SELECT seller_id FROM order_items WHERE order_id = $1 AND product_id = $2", p.OrderID, p.ProductID).Scan(&sellerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.MakeNotFound(fmt.Errorf("order item not found"))
		}
		return nil, fmt.Errorf("error checking seller: %v", err)
	}

	if sellerID != userID {
		return nil, NewForbiddenError("user is not authorized to update this order item")
	}

	// Update the status of the order item
	_, err = tx.ExecContext(ctx, "UPDATE order_items SET status = $1 WHERE order_id = $2 AND product_id = $3", p.Status, p.OrderID, p.ProductID)
	if err != nil {
		return nil, fmt.Errorf("error updating order item status: %v", err)
	}

	// Check if all items in the order have the same status
	var allSameStatus bool
	err = tx.QueryRowContext(ctx, "SELECT COUNT(DISTINCT status) = 1 FROM order_items WHERE order_id = $1", p.OrderID).Scan(&allSameStatus)
	if err != nil {
		return nil, fmt.Errorf("error checking order items status: %v", err)
	}

	// Update the overall order status if all items have the same status
	if allSameStatus {
		_, err = tx.ExecContext(ctx, "UPDATE orders SET overall_status = $1 WHERE id = $2", p.Status, p.OrderID)
		if err != nil {
			return nil, fmt.Errorf("error updating overall order status: %v", err)
		}
	} else {
		// Set overall status to "in progress" if items have different statuses
		_, err = tx.ExecContext(ctx, "UPDATE orders SET overall_status = 'in progress' WHERE id = $1", p.OrderID)
		if err != nil {
			return nil, fmt.Errorf("error updating overall order status: %v", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("error committing transaction: %v", err)
	}

	// Fetch and return the updated order
	return s.GetOrder(ctx, &store.GetOrderPayload{ID: p.OrderID})
}

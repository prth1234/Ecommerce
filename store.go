package storeapi

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	store "store/gen/store"
)

// store service example implementation.
// The example methods log the requests and return zero values.
type storesrvc struct {
	users    map[string]*store.User
	products map[string]*store.Product
	orders   map[string]*store.Order
	carts    map[string]*store.Cart
	mutex    sync.RWMutex
}

// NewStore returns the store service implementation.
func NewStore() store.Service {
	return &storesrvc{
		users:    make(map[string]*store.User),
		products: make(map[string]*store.Product),
		orders:   make(map[string]*store.Order),
		carts:    make(map[string]*store.Cart),
	}
}

// CreateUser implements createUser.
func (s *storesrvc) CreateUser(ctx context.Context, p *store.NewUser) (res *store.User, err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	id := uuid.New().String()
	user := &store.User{
		ID:        id,
		Username:  p.Username,
		Email:     p.Email,
		FirstName: p.FirstName,
		LastName:  p.LastName,
	}
	s.users[id] = user
	return user, nil
}

// GetUser implements getUser.
func (s *storesrvc) GetUser(ctx context.Context, p *store.GetUserPayload) (res *store.User, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	user, ok := s.users[p.ID]
	if !ok {
		return nil, store.MakeNotFound(fmt.Errorf("user not found"))
	}
	return user, nil
}

// GetUser implements getUser.
func (s *storesrvc) GetUserAll(ctx context.Context) (res []*store.User, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	//user, ok := s.users[p.ID]
	users := make([]*store.User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}

	return users, nil
}

// CreateProduct implements createProduct.
func (s *storesrvc) CreateProduct(ctx context.Context, p *store.NewProduct) (res *store.Product, err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	id := uuid.New().String()
	product := &store.Product{
		ID:          id,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Inventory:   p.Inventory,
	}
	s.products[id] = product
	return product, nil
}

// GetProduct implements getProduct.
func (s *storesrvc) GetProduct(ctx context.Context, p *store.GetProductPayload) (res *store.Product, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	product, ok := s.products[p.ID]
	if !ok {
		return nil, store.MakeNotFound(fmt.Errorf("product not found"))
	}
	return product, nil
}

// ListProducts implements listProducts.
func (s *storesrvc) ListProducts(ctx context.Context) (res []*store.Product, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	products := make([]*store.Product, 0, len(s.products))
	for _, product := range s.products {
		products = append(products, product)
	}
	return products, nil
}

// CreateOrder implements createOrder.
func (s *storesrvc) CreateOrder(ctx context.Context, p *store.NewOrder) (res *store.Order, err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	id := uuid.New().String()
	totalAmount := 0.0
	for _, item := range p.Items {
		totalAmount += item.Price * float64(item.Quantity)
	}
	order := &store.Order{
		ID:          id,
		UserID:      p.UserID,
		Items:       p.Items,
		TotalAmount: totalAmount,
		Status:      "pending",
	}
	s.orders[id] = order
	return order, nil
}

// GetOrder implements getOrder.
func (s *storesrvc) GetOrder(ctx context.Context, p *store.GetOrderPayload) (res *store.Order, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	order, ok := s.orders[p.ID]
	if !ok {
		return nil, store.MakeNotFound(fmt.Errorf("order not found"))
	}
	return order, nil
}

// AddToCart implements addToCart.
func (s *storesrvc) AddToCart(ctx context.Context, p *store.CartItem) (res *store.Cart, err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	cart, ok := s.carts[p.ProductID]
	if !ok {
		cart = &store.Cart{
			UserID: p.ProductID, // Using ProductID as UserID for simplicity
			Items:  []*store.CartItem{},
		}
		s.carts[p.ProductID] = cart
	}

	cart.Items = append(cart.Items, p)

	// Recalculate total amount
	cart.TotalAmount = 0
	for _, item := range cart.Items {
		product, ok := s.products[item.ProductID]
		if ok {
			cart.TotalAmount += product.Price * float64(item.Quantity)
		}
	}

	return cart, nil
}

// GetCart implements getCart.
func (s *storesrvc) GetCart(ctx context.Context) (res *store.Cart, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	// For simplicity, we're just returning the first cart we find
	for _, cart := range s.carts {
		return cart, nil
	}

	return &store.Cart{Items: []*store.CartItem{}}, nil
}

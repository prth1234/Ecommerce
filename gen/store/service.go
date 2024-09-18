// Code generated by goa v3.19.0, DO NOT EDIT.
//
// store service
//
// Command:
// $ goa gen store/design

package store

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Store service
type Service interface {
	// CreateUser implements createUser.
	CreateUser(context.Context, *NewUser) (res *User, err error)
	// GetUser implements getUser.
	GetUser(context.Context, *GetUserPayload) (res *User, err error)
	// GetUserAll implements getUserAll.
	GetUserAll(context.Context) (res []*User, err error)
	// CreateProduct implements createProduct.
	CreateProduct(context.Context, *NewProduct) (res *Product, err error)
	// GetProduct implements getProduct.
	GetProduct(context.Context, *GetProductPayload) (res *Product, err error)
	// ListProducts implements listProducts.
	ListProducts(context.Context) (res []*Product, err error)
	// CreateOrder implements createOrder.
	CreateOrder(context.Context, *NewOrder) (res *Order, err error)
	// GetOrder implements getOrder.
	GetOrder(context.Context, *GetOrderPayload) (res *Order, err error)
	// Retrieve all orders for a specific user
	GetUserOrders(context.Context, *GetUserOrdersPayload) (res []*Order, err error)
	// AddToCart implements addToCart.
	AddToCart(context.Context, *CartItem) (res *Cart, err error)
	// GetCart implements getCart.
	GetCart(context.Context, *GetCartPayload) (res *Cart, err error)
}

// APIName is the name of the API as defined in the design.
const APIName = "store"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.0.1"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "store"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [11]string{"createUser", "getUser", "getUserAll", "createProduct", "getProduct", "listProducts", "createOrder", "getOrder", "getUserOrders\t", "addToCart", "getCart"}

// Cart is the result type of the store service addToCart method.
type Cart struct {
	// Unique cart ID
	ID string
	// ID of the user who owns the cart
	UserID string
	// Items in the cart
	Items []*CartItem
	// Total amount of items in the cart
	TotalAmount float64
}

// CartItem is the payload type of the store service addToCart method.
type CartItem struct {
	// ID of the user who owns the cart
	UserID string
	// ID of the product
	ProductID string
	// Quantity of the product
	Quantity int
	// Price of the product
	Price *float64
}

// GetCartPayload is the payload type of the store service getCart method.
type GetCartPayload struct {
	// ID of the user whose cart to retrieve
	UserID string
}

// GetOrderPayload is the payload type of the store service getOrder method.
type GetOrderPayload struct {
	ID string
}

// GetProductPayload is the payload type of the store service getProduct method.
type GetProductPayload struct {
	ID string
}

// GetUserOrdersPayload is the payload type of the store service getUserOrders
// method.
type GetUserOrdersPayload struct {
	UserID string
}

// GetUserPayload is the payload type of the store service getUser method.
type GetUserPayload struct {
	ID string
}

// NewOrder is the payload type of the store service createOrder method.
type NewOrder struct {
	// ID of the user placing the order
	UserID string
	// Items in the order
	Items []*OrderItem
}

// NewProduct is the payload type of the store service createProduct method.
type NewProduct struct {
	// Product name
	Name string
	// Product description
	Description *string
	// Product price
	Price float64
	// Available inventory
	Inventory int
}

// NewUser is the payload type of the store service createUser method.
type NewUser struct {
	// User's username
	Username string
	// User's email address
	Email string
	// User's first name
	FirstName *string
	// User's last name
	LastName *string
}

// Order is the result type of the store service createOrder method.
type Order struct {
	// Unique order ID
	ID string
	// ID of the user who placed the order
	UserID string
	// Items in the order
	Items []*OrderItem
	// Total amount of the order
	TotalAmount float64
	// Order status
	Status string
}

type OrderItem struct {
	// ID of the product
	ProductID string
	// Quantity of the product
	Quantity int
	// Price of the product at the time of order
	Price float64
}

// Product is the result type of the store service createProduct method.
type Product struct {
	// Unique product ID
	ID string
	// Product name
	Name string
	// Product description
	Description *string
	// Product price
	Price float64
	// Available inventory
	Inventory int
}

// User is the result type of the store service createUser method.
type User struct {
	// Unique user ID
	ID string
	// User's username
	Username string
	// User's email address
	Email string
	// User's first name
	FirstName *string
	// User's last name
	LastName *string
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "not-found", false, false, false)
}

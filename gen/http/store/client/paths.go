// Code generated by goa v3.19.0, DO NOT EDIT.
//
// HTTP request path constructors for the store service.
//
// Command:
// $ goa gen store/design

package client

import (
	"fmt"
)

// CreateUserStorePath returns the URL path to the store service createUser HTTP endpoint.
func CreateUserStorePath() string {
	return "/users"
}

// LoginUserStorePath returns the URL path to the store service loginUser HTTP endpoint.
func LoginUserStorePath() string {
	return "/login"
}

// GetUserStorePath returns the URL path to the store service getUser HTTP endpoint.
func GetUserStorePath(id string) string {
	return fmt.Sprintf("/users/%v", id)
}

// GetUserAllStorePath returns the URL path to the store service getUserAll HTTP endpoint.
func GetUserAllStorePath() string {
	return "/users"
}

// UpdateUserStorePath returns the URL path to the store service updateUser HTTP endpoint.
func UpdateUserStorePath() string {
	return "/users/update"
}

// DeleteUserStorePath returns the URL path to the store service deleteUser HTTP endpoint.
func DeleteUserStorePath() string {
	return "/users/delete"
}

// CreateProductStorePath returns the URL path to the store service createProduct HTTP endpoint.
func CreateProductStorePath() string {
	return "/products"
}

// GetProductStorePath returns the URL path to the store service getProduct HTTP endpoint.
func GetProductStorePath(id string) string {
	return fmt.Sprintf("/products/%v", id)
}

// ListProductsStorePath returns the URL path to the store service listProducts HTTP endpoint.
func ListProductsStorePath() string {
	return "/products"
}

// AddToCartStorePath returns the URL path to the store service addToCart HTTP endpoint.
func AddToCartStorePath() string {
	return "/cart/item"
}

// RemoveFromCartStorePath returns the URL path to the store service removeFromCart HTTP endpoint.
func RemoveFromCartStorePath(productID string) string {
	return fmt.Sprintf("/cart/item/%v", productID)
}

// GetCartStorePath returns the URL path to the store service getCart HTTP endpoint.
func GetCartStorePath() string {
	return "/cart"
}

// CreateOrderStorePath returns the URL path to the store service createOrder HTTP endpoint.
func CreateOrderStorePath() string {
	return "/orders"
}

// DeleteOrderStorePath returns the URL path to the store service deleteOrder HTTP endpoint.
func DeleteOrderStorePath(id string) string {
	return fmt.Sprintf("/orders/%v", id)
}

// GetOrderStorePath returns the URL path to the store service getOrder HTTP endpoint.
func GetOrderStorePath(id string) string {
	return fmt.Sprintf("/orders/%v", id)
}

// GetUserOrdersStorePath returns the URL path to the store service getUserOrders HTTP endpoint.
func GetUserOrdersStorePath() string {
	return "/orders"
}

// GetProductsPostedByUserStorePath returns the URL path to the store service getProductsPostedByUser HTTP endpoint.
func GetProductsPostedByUserStorePath() string {
	return "/users/products"
}

// UpdateOrderItemStatusStorePath returns the URL path to the store service updateOrderItemStatus HTTP endpoint.
func UpdateOrderItemStatusStorePath(orderID string, productID string) string {
	return fmt.Sprintf("/orders/%v/items/%v", orderID, productID)
}

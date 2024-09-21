package design

import (
	. "goa.design/goa/v3/dsl"
)

var User = Type("User", func() {
	Attribute("id", String, "Unique user ID")
	Attribute("username", String, "User's username")
	Attribute("email", String, "User's email address")
	Attribute("firstName", String, "User's first name")
	Attribute("lastName", String, "User's last name")
	//Attribute("password", String, "User's password")
	Required("id", "username", "email")
})

var NewUser = Type("NewUser", func() {
	Attribute("username", String, "User's username")
	Attribute("email", String, "User's email address")
	Attribute("firstName", String, "User's first name")
	Attribute("lastName", String, "User's last name")
	Attribute("password", String, "User's password")
	Required("username", "email", "firstName", "lastName", "password")
})

var Product = Type("Product", func() {
	Attribute("userId", String, "Product's owner's user ID")
	Attribute("id", String, "Unique product ID")
	Attribute("name", String, "Product name")
	Attribute("description", String, "Product description")
	Attribute("price", Float64, "Product price")
	Attribute("inventory", Int, "Available inventory")
	Required("id", "name", "price", "inventory", "userId")
})

var NewProduct = Type("NewProduct", func() {
	Attribute("name", String, "Product name")
	Attribute("description", String, "Product description")
	Attribute("price", Float64, "Product price")
	Attribute("inventory", Int, "Available inventory")
	Required("name", "price", "inventory")
})

var CartItem = Type("CartItem", func() {
	Attribute("productID", String, "ID of the product")
	Attribute("quantity", Int, "Quantity of the product")
	Required("productID", "quantity")
})

var OrderItem = Type("OrderItem", func() {
	Attribute("productID", String, "ID of the product")
	Attribute("quantity", Int, "Quantity of the product")
	Attribute("price", Float64, "Price of the product at the time of order")
	Required("productID", "quantity", "price")
})

var Cart = Type("Cart", func() {
	Attribute("id", String, "Unique cart ID")
	Attribute("userID", String, "ID of the user who owns the cart")
	Attribute("items", ArrayOf(CartItem), "Items in the cart")
	Attribute("totalAmount", Float64, "Total amount of the cart")
	Required("id", "userID", "items", "totalAmount")
})

var Order = Type("Order", func() {
	Attribute("id", String, "Unique order ID")
	Attribute("userID", String, "ID of the user who placed the order")
	Attribute("items", ArrayOf(OrderItem), "Items in the order")
	Attribute("totalAmount", Float64, "Total amount of the order")
	Attribute("status", String, "Order status")
	Required("id", "userID", "items", "totalAmount", "status")
})

var UserUpdatePayload = Type("UserUpdatePayload", func() {
	Attribute("email", String, "User's email address")
	Attribute("firstName", String, "User's first name")
	Attribute("lastName", String, "User's last name")
	Required("email", "firstName", "lastName")
})

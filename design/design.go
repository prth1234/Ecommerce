package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("store", func() {
	Title("Store API")
	Description("Service for a complete store with users, products, orders, and cart functionality")
	Server("store", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})
})

var _ = Service("store", func() {
	Description("Store service")

	// User endpoints
	Method("createUser", func() {
		Payload(NewUser)
		Result(User)
		HTTP(func() {
			POST("/users")
			Response(StatusCreated)
		})
	})

	Method("getUser", func() {
		Payload(func() {
			Field(1, "id", String)
			Required("id")
		})
		Result(User)
		Error("not-found")
		HTTP(func() {
			GET("/users/{id}")
			Response(StatusOK)
			Response("not-found", StatusNotFound)
		})
	})

	Method("getUserAll", func() {
		Result(ArrayOf(User))
		HTTP(func() {
			GET("/users")
			Response(StatusOK)
		})
	})

	// Product endpoints
	Method("createProduct", func() {
		Payload(NewProduct)
		Result(Product)
		HTTP(func() {
			POST("/products")
			Response(StatusCreated)
		})
	})

	Method("getProduct", func() {
		Payload(func() {
			Field(1, "id", String)
			Required("id")
		})
		Result(Product)
		Error("not-found")
		HTTP(func() {
			GET("/products/{id}")
			Response(StatusOK)
			Response("not-found", StatusNotFound)
		})
	})

	Method("listProducts", func() {
		Result(ArrayOf(Product))
		HTTP(func() {
			GET("/products")
			Response(StatusOK)
		})
	})

	// Order endpoints
	Method("createOrder", func() {
		Payload(NewOrder)
		Result(Order)
		HTTP(func() {
			POST("/orders")
			Response(StatusCreated)
		})
	})

	Method("getOrder", func() {
		Payload(func() {
			Field(1, "id", String)
			Required("id")
		})
		Result(Order)
		Error("not-found")
		HTTP(func() {
			GET("/orders/{id}")
			Response(StatusOK)
			Response("not-found", StatusNotFound)
		})
	})

	// Cart endpoints
	Method("addToCart", func() {
		Payload(CartItem)
		Result(Cart)
		HTTP(func() {
			POST("/cart/items")
			Response(StatusOK)
		})
	})

	Method("getCart", func() {
		Payload(GetCartPayload)
		Result(Cart)
		Error("not-found")
		HTTP(func() {
			GET("/cart")
			Response(StatusOK)
			Response("not-found", StatusNotFound)
		})
	})
})

var User = Type("User", func() {
	Attribute("id", String, "Unique user ID")
	Attribute("username", String, "User's username")
	Attribute("email", String, "User's email address")
	Attribute("firstName", String, "User's first name")
	Attribute("lastName", String, "User's last name")
	Required("id", "username", "email")
})

var NewUser = Type("NewUser", func() {
	Attribute("username", String, "User's username")
	Attribute("email", String, "User's email address")
	Attribute("firstName", String, "User's first name")
	Attribute("lastName", String, "User's last name")
	Required("username", "email")
})

var Product = Type("Product", func() {
	Attribute("id", String, "Unique product ID")
	Attribute("name", String, "Product name")
	Attribute("description", String, "Product description")
	Attribute("price", Float64, "Product price")
	Attribute("inventory", Int, "Available inventory")
	Required("id", "name", "price", "inventory")
})

var NewProduct = Type("NewProduct", func() {
	Attribute("name", String, "Product name")
	Attribute("description", String, "Product description")
	Attribute("price", Float64, "Product price")
	Attribute("inventory", Int, "Available inventory")
	Required("name", "price", "inventory")
})

var Order = Type("Order", func() {
	Attribute("id", String, "Unique order ID")
	Attribute("userID", String, "ID of the user who placed the order")
	Attribute("items", ArrayOf(OrderItem), "Items in the order")
	Attribute("totalAmount", Float64, "Total amount of the order")
	Attribute("status", String, "Order status")
	Required("id", "userID", "items", "totalAmount", "status")
})

var NewOrder = Type("NewOrder", func() {
	Attribute("userID", String, "ID of the user placing the order")
	Attribute("items", ArrayOf(OrderItem), "Items in the order")
	Required("userID", "items")
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
	Attribute("totalAmount", Float64, "Total amount of items in the cart")
	Required("id", "userID", "items", "totalAmount")
})

var CartItem = Type("CartItem", func() {
	Attribute("userID", String, "ID of the user who owns the cart")
	Attribute("productID", String, "ID of the product")
	Attribute("quantity", Int, "Quantity of the product")
	Attribute("price", Float64, "Price of the product")
	Required("userID", "productID", "quantity")
})

var GetCartPayload = Type("GetCartPayload", func() {
	Attribute("userID", String, "ID of the user whose cart to retrieve")
	Required("userID")
})

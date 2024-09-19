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

	// Login endpoint
	Method("loginUser", func() {
		Description("Login a user and return a JWT token")
		Payload(func() {
			Field(1, "username", String, "Username for login")
			Field(2, "password", String, "Password for login")
			Required("username", "password")
		})
		Result(func() {
			Field(1, "token", String, "JWT token for the authenticated user")
		})
		Error("unauthorized", String, "Invalid username or password")
		HTTP(func() {
			POST("/login")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
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

	Method("updateUser", func() {
		Payload(UserUpdatePayload)
		Result(User)
		HTTP(func() {
			POST("/users/update")
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

	Method("getUserOrders", func() {
		Description("Retrieve all orders for a specific user")
		Payload(func() {
			Field(1, "userID", String)
			Required("userID")
		})
		Result(ArrayOf(Order))
		Error("not-found")
		HTTP(func() {
			GET("/users/{userID}/orders")
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

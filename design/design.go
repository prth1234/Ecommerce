package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("store", func() {
	Title("Store API")
	Description("Service for a complete store with users, products, carts, and orders")
	Server("store", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})
})

var _ = Service("store", func() {
	Description("Store service")

	// User endpoints (unchanged)
	Method("createUser", func() {
		Payload(NewUser)
		Result(User)
		HTTP(func() {
			POST("/users")
			Response(StatusCreated)
		})
	})

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

	Method("deleteUser", func() {
		HTTP(func() {
			POST("/users/delete")
			Response(StatusOK)
		})
	})

	// Product endpoints (unchanged)
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

	// Cart endpoints
	Method("addToCart", func() {
		Payload(CartItem)
		Result(Cart)
		HTTP(func() {
			POST("/cart/item")
			Response(StatusOK)
		})
	})

	Method("removeFromCart", func() {
		Payload(func() {
			Field(1, "productID", String)
			Required("productID")
		})
		Result(Cart)
		HTTP(func() {
			DELETE("/cart/item/{productID}")
			Response(StatusOK)
		})
	})

	Method("getCart", func() {
		Result(Cart)
		Error("not-found")
		HTTP(func() {
			GET("/cart")
			Response(StatusOK)
			Response("not-found", StatusNotFound)
		})
	})

	Method("createOrder", func() {
		Description("Create an order from the current cart")
		Result(Order)
		HTTP(func() {
			POST("/orders")
			Response(StatusCreated)
		})
	})

	Method("deleteOrder", func() {
		Payload(func() {
			Field(1, "id", String)
			Required("id")
		})
		Description("Delete an order from the current cart")
		HTTP(func() {
			DELETE("/orders/{id}")
			Response(StatusOK)
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
		Description("Retrieve all orders for the authenticated user")
		Result(ArrayOf(Order))
		HTTP(func() {
			GET("/orders")
			Response(StatusOK)
		})
	})

	Method("getProductsPostedByUser", func() {
		Description("Retrieve all products posted by the user")
		Result(ArrayOf(Product))
		HTTP(func() {
			GET("/users/products")
			Response(StatusOK)
		})
	})
	Method("updateOrderItemStatus", func() {
		Description("Update the status of an order item")
		Payload(func() {
			Field(1, "orderID", String, "ID of the order")
			Field(2, "productID", String, "ID of the product in the order")
			Field(3, "status", String, "New status for the order item")
			Required("orderID", "productID", "status")
		})
		Result(Order)
		Error("not-found", String, "Order or product not found")
		Error("forbidden", String, "User is not authorized to update this order item")
		HTTP(func() {
			PATCH("/orders/{orderID}/items/{productID}")
			Response(StatusOK)
			Response("not-found", StatusNotFound)
			Response("forbidden", StatusForbidden)
		})
	})
})

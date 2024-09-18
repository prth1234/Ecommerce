// Code generated by goa v3.19.0, DO NOT EDIT.
//
// store HTTP server types
//
// Command:
// $ goa gen store/design

package server

import (
	store "store/gen/store"

	goa "goa.design/goa/v3/pkg"
)

// CreateUserRequestBody is the type of the "store" service "createUser"
// endpoint HTTP request body.
type CreateUserRequestBody struct {
	// User's username
	Username *string `form:"username,omitempty" json:"username,omitempty" xml:"username,omitempty"`
	// User's email address
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// User's first name
	FirstName *string `form:"firstName,omitempty" json:"firstName,omitempty" xml:"firstName,omitempty"`
	// User's last name
	LastName *string `form:"lastName,omitempty" json:"lastName,omitempty" xml:"lastName,omitempty"`
}

// CreateProductRequestBody is the type of the "store" service "createProduct"
// endpoint HTTP request body.
type CreateProductRequestBody struct {
	// Product name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Product description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Product price
	Price *float64 `form:"price,omitempty" json:"price,omitempty" xml:"price,omitempty"`
	// Available inventory
	Inventory *int `form:"inventory,omitempty" json:"inventory,omitempty" xml:"inventory,omitempty"`
}

// CreateOrderRequestBody is the type of the "store" service "createOrder"
// endpoint HTTP request body.
type CreateOrderRequestBody struct {
	// ID of the user placing the order
	UserID *string `form:"userID,omitempty" json:"userID,omitempty" xml:"userID,omitempty"`
	// Items in the order
	Items []*OrderItemRequestBody `form:"items,omitempty" json:"items,omitempty" xml:"items,omitempty"`
}

// AddToCartRequestBody is the type of the "store" service "addToCart" endpoint
// HTTP request body.
type AddToCartRequestBody struct {
	// ID of the user who owns the cart
	UserID *string `form:"userID,omitempty" json:"userID,omitempty" xml:"userID,omitempty"`
	// ID of the product
	ProductID *string `form:"productID,omitempty" json:"productID,omitempty" xml:"productID,omitempty"`
	// Quantity of the product
	Quantity *int `form:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// Price of the product
	Price *float64 `form:"price,omitempty" json:"price,omitempty" xml:"price,omitempty"`
}

// GetCartRequestBody is the type of the "store" service "getCart" endpoint
// HTTP request body.
type GetCartRequestBody struct {
	// ID of the user whose cart to retrieve
	UserID *string `form:"userID,omitempty" json:"userID,omitempty" xml:"userID,omitempty"`
}

// CreateUserResponseBody is the type of the "store" service "createUser"
// endpoint HTTP response body.
type CreateUserResponseBody struct {
	// Unique user ID
	ID string `form:"id" json:"id" xml:"id"`
	// User's username
	Username string `form:"username" json:"username" xml:"username"`
	// User's email address
	Email string `form:"email" json:"email" xml:"email"`
	// User's first name
	FirstName *string `form:"firstName,omitempty" json:"firstName,omitempty" xml:"firstName,omitempty"`
	// User's last name
	LastName *string `form:"lastName,omitempty" json:"lastName,omitempty" xml:"lastName,omitempty"`
}

// GetUserResponseBody is the type of the "store" service "getUser" endpoint
// HTTP response body.
type GetUserResponseBody struct {
	// Unique user ID
	ID string `form:"id" json:"id" xml:"id"`
	// User's username
	Username string `form:"username" json:"username" xml:"username"`
	// User's email address
	Email string `form:"email" json:"email" xml:"email"`
	// User's first name
	FirstName *string `form:"firstName,omitempty" json:"firstName,omitempty" xml:"firstName,omitempty"`
	// User's last name
	LastName *string `form:"lastName,omitempty" json:"lastName,omitempty" xml:"lastName,omitempty"`
}

// GetUserAllResponseBody is the type of the "store" service "getUserAll"
// endpoint HTTP response body.
type GetUserAllResponseBody []*UserResponse

// CreateProductResponseBody is the type of the "store" service "createProduct"
// endpoint HTTP response body.
type CreateProductResponseBody struct {
	// Unique product ID
	ID string `form:"id" json:"id" xml:"id"`
	// Product name
	Name string `form:"name" json:"name" xml:"name"`
	// Product description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Product price
	Price float64 `form:"price" json:"price" xml:"price"`
	// Available inventory
	Inventory int `form:"inventory" json:"inventory" xml:"inventory"`
}

// GetProductResponseBody is the type of the "store" service "getProduct"
// endpoint HTTP response body.
type GetProductResponseBody struct {
	// Unique product ID
	ID string `form:"id" json:"id" xml:"id"`
	// Product name
	Name string `form:"name" json:"name" xml:"name"`
	// Product description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Product price
	Price float64 `form:"price" json:"price" xml:"price"`
	// Available inventory
	Inventory int `form:"inventory" json:"inventory" xml:"inventory"`
}

// ListProductsResponseBody is the type of the "store" service "listProducts"
// endpoint HTTP response body.
type ListProductsResponseBody []*ProductResponse

// CreateOrderResponseBody is the type of the "store" service "createOrder"
// endpoint HTTP response body.
type CreateOrderResponseBody struct {
	// Unique order ID
	ID string `form:"id" json:"id" xml:"id"`
	// ID of the user who placed the order
	UserID string `form:"userID" json:"userID" xml:"userID"`
	// Items in the order
	Items []*OrderItemResponseBody `form:"items" json:"items" xml:"items"`
	// Total amount of the order
	TotalAmount float64 `form:"totalAmount" json:"totalAmount" xml:"totalAmount"`
	// Order status
	Status string `form:"status" json:"status" xml:"status"`
}

// GetOrderResponseBody is the type of the "store" service "getOrder" endpoint
// HTTP response body.
type GetOrderResponseBody struct {
	// Unique order ID
	ID string `form:"id" json:"id" xml:"id"`
	// ID of the user who placed the order
	UserID string `form:"userID" json:"userID" xml:"userID"`
	// Items in the order
	Items []*OrderItemResponseBody `form:"items" json:"items" xml:"items"`
	// Total amount of the order
	TotalAmount float64 `form:"totalAmount" json:"totalAmount" xml:"totalAmount"`
	// Order status
	Status string `form:"status" json:"status" xml:"status"`
}

// AddToCartResponseBody is the type of the "store" service "addToCart"
// endpoint HTTP response body.
type AddToCartResponseBody struct {
	// Unique cart ID
	ID string `form:"id" json:"id" xml:"id"`
	// ID of the user who owns the cart
	UserID string `form:"userID" json:"userID" xml:"userID"`
	// Items in the cart
	Items []*CartItemResponseBody `form:"items" json:"items" xml:"items"`
	// Total amount of items in the cart
	TotalAmount float64 `form:"totalAmount" json:"totalAmount" xml:"totalAmount"`
}

// GetCartResponseBody is the type of the "store" service "getCart" endpoint
// HTTP response body.
type GetCartResponseBody struct {
	// Unique cart ID
	ID string `form:"id" json:"id" xml:"id"`
	// ID of the user who owns the cart
	UserID string `form:"userID" json:"userID" xml:"userID"`
	// Items in the cart
	Items []*CartItemResponseBody `form:"items" json:"items" xml:"items"`
	// Total amount of items in the cart
	TotalAmount float64 `form:"totalAmount" json:"totalAmount" xml:"totalAmount"`
}

// GetUserNotFoundResponseBody is the type of the "store" service "getUser"
// endpoint HTTP response body for the "not-found" error.
type GetUserNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetProductNotFoundResponseBody is the type of the "store" service
// "getProduct" endpoint HTTP response body for the "not-found" error.
type GetProductNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetOrderNotFoundResponseBody is the type of the "store" service "getOrder"
// endpoint HTTP response body for the "not-found" error.
type GetOrderNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetCartNotFoundResponseBody is the type of the "store" service "getCart"
// endpoint HTTP response body for the "not-found" error.
type GetCartNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UserResponse is used to define fields on response body types.
type UserResponse struct {
	// Unique user ID
	ID string `form:"id" json:"id" xml:"id"`
	// User's username
	Username string `form:"username" json:"username" xml:"username"`
	// User's email address
	Email string `form:"email" json:"email" xml:"email"`
	// User's first name
	FirstName *string `form:"firstName,omitempty" json:"firstName,omitempty" xml:"firstName,omitempty"`
	// User's last name
	LastName *string `form:"lastName,omitempty" json:"lastName,omitempty" xml:"lastName,omitempty"`
}

// ProductResponse is used to define fields on response body types.
type ProductResponse struct {
	// Unique product ID
	ID string `form:"id" json:"id" xml:"id"`
	// Product name
	Name string `form:"name" json:"name" xml:"name"`
	// Product description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Product price
	Price float64 `form:"price" json:"price" xml:"price"`
	// Available inventory
	Inventory int `form:"inventory" json:"inventory" xml:"inventory"`
}

// OrderItemResponseBody is used to define fields on response body types.
type OrderItemResponseBody struct {
	// ID of the product
	ProductID string `form:"productID" json:"productID" xml:"productID"`
	// Quantity of the product
	Quantity int `form:"quantity" json:"quantity" xml:"quantity"`
	// Price of the product at the time of order
	Price float64 `form:"price" json:"price" xml:"price"`
}

// CartItemResponseBody is used to define fields on response body types.
type CartItemResponseBody struct {
	// ID of the user who owns the cart
	UserID string `form:"userID" json:"userID" xml:"userID"`
	// ID of the product
	ProductID string `form:"productID" json:"productID" xml:"productID"`
	// Quantity of the product
	Quantity int `form:"quantity" json:"quantity" xml:"quantity"`
	// Price of the product
	Price *float64 `form:"price,omitempty" json:"price,omitempty" xml:"price,omitempty"`
}

// OrderItemRequestBody is used to define fields on request body types.
type OrderItemRequestBody struct {
	// ID of the product
	ProductID *string `form:"productID,omitempty" json:"productID,omitempty" xml:"productID,omitempty"`
	// Quantity of the product
	Quantity *int `form:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
	// Price of the product at the time of order
	Price *float64 `form:"price,omitempty" json:"price,omitempty" xml:"price,omitempty"`
}

// NewCreateUserResponseBody builds the HTTP response body from the result of
// the "createUser" endpoint of the "store" service.
func NewCreateUserResponseBody(res *store.User) *CreateUserResponseBody {
	body := &CreateUserResponseBody{
		ID:        res.ID,
		Username:  res.Username,
		Email:     res.Email,
		FirstName: res.FirstName,
		LastName:  res.LastName,
	}
	return body
}

// NewGetUserResponseBody builds the HTTP response body from the result of the
// "getUser" endpoint of the "store" service.
func NewGetUserResponseBody(res *store.User) *GetUserResponseBody {
	body := &GetUserResponseBody{
		ID:        res.ID,
		Username:  res.Username,
		Email:     res.Email,
		FirstName: res.FirstName,
		LastName:  res.LastName,
	}
	return body
}

// NewGetUserAllResponseBody builds the HTTP response body from the result of
// the "getUserAll" endpoint of the "store" service.
func NewGetUserAllResponseBody(res []*store.User) GetUserAllResponseBody {
	body := make([]*UserResponse, len(res))
	for i, val := range res {
		body[i] = marshalStoreUserToUserResponse(val)
	}
	return body
}

// NewCreateProductResponseBody builds the HTTP response body from the result
// of the "createProduct" endpoint of the "store" service.
func NewCreateProductResponseBody(res *store.Product) *CreateProductResponseBody {
	body := &CreateProductResponseBody{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
		Price:       res.Price,
		Inventory:   res.Inventory,
	}
	return body
}

// NewGetProductResponseBody builds the HTTP response body from the result of
// the "getProduct" endpoint of the "store" service.
func NewGetProductResponseBody(res *store.Product) *GetProductResponseBody {
	body := &GetProductResponseBody{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
		Price:       res.Price,
		Inventory:   res.Inventory,
	}
	return body
}

// NewListProductsResponseBody builds the HTTP response body from the result of
// the "listProducts" endpoint of the "store" service.
func NewListProductsResponseBody(res []*store.Product) ListProductsResponseBody {
	body := make([]*ProductResponse, len(res))
	for i, val := range res {
		body[i] = marshalStoreProductToProductResponse(val)
	}
	return body
}

// NewCreateOrderResponseBody builds the HTTP response body from the result of
// the "createOrder" endpoint of the "store" service.
func NewCreateOrderResponseBody(res *store.Order) *CreateOrderResponseBody {
	body := &CreateOrderResponseBody{
		ID:          res.ID,
		UserID:      res.UserID,
		TotalAmount: res.TotalAmount,
		Status:      res.Status,
	}
	if res.Items != nil {
		body.Items = make([]*OrderItemResponseBody, len(res.Items))
		for i, val := range res.Items {
			body.Items[i] = marshalStoreOrderItemToOrderItemResponseBody(val)
		}
	} else {
		body.Items = []*OrderItemResponseBody{}
	}
	return body
}

// NewGetOrderResponseBody builds the HTTP response body from the result of the
// "getOrder" endpoint of the "store" service.
func NewGetOrderResponseBody(res *store.Order) *GetOrderResponseBody {
	body := &GetOrderResponseBody{
		ID:          res.ID,
		UserID:      res.UserID,
		TotalAmount: res.TotalAmount,
		Status:      res.Status,
	}
	if res.Items != nil {
		body.Items = make([]*OrderItemResponseBody, len(res.Items))
		for i, val := range res.Items {
			body.Items[i] = marshalStoreOrderItemToOrderItemResponseBody(val)
		}
	} else {
		body.Items = []*OrderItemResponseBody{}
	}
	return body
}

// NewAddToCartResponseBody builds the HTTP response body from the result of
// the "addToCart" endpoint of the "store" service.
func NewAddToCartResponseBody(res *store.Cart) *AddToCartResponseBody {
	body := &AddToCartResponseBody{
		ID:          res.ID,
		UserID:      res.UserID,
		TotalAmount: res.TotalAmount,
	}
	if res.Items != nil {
		body.Items = make([]*CartItemResponseBody, len(res.Items))
		for i, val := range res.Items {
			body.Items[i] = marshalStoreCartItemToCartItemResponseBody(val)
		}
	} else {
		body.Items = []*CartItemResponseBody{}
	}
	return body
}

// NewGetCartResponseBody builds the HTTP response body from the result of the
// "getCart" endpoint of the "store" service.
func NewGetCartResponseBody(res *store.Cart) *GetCartResponseBody {
	body := &GetCartResponseBody{
		ID:          res.ID,
		UserID:      res.UserID,
		TotalAmount: res.TotalAmount,
	}
	if res.Items != nil {
		body.Items = make([]*CartItemResponseBody, len(res.Items))
		for i, val := range res.Items {
			body.Items[i] = marshalStoreCartItemToCartItemResponseBody(val)
		}
	} else {
		body.Items = []*CartItemResponseBody{}
	}
	return body
}

// NewGetUserNotFoundResponseBody builds the HTTP response body from the result
// of the "getUser" endpoint of the "store" service.
func NewGetUserNotFoundResponseBody(res *goa.ServiceError) *GetUserNotFoundResponseBody {
	body := &GetUserNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetProductNotFoundResponseBody builds the HTTP response body from the
// result of the "getProduct" endpoint of the "store" service.
func NewGetProductNotFoundResponseBody(res *goa.ServiceError) *GetProductNotFoundResponseBody {
	body := &GetProductNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetOrderNotFoundResponseBody builds the HTTP response body from the
// result of the "getOrder" endpoint of the "store" service.
func NewGetOrderNotFoundResponseBody(res *goa.ServiceError) *GetOrderNotFoundResponseBody {
	body := &GetOrderNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetCartNotFoundResponseBody builds the HTTP response body from the result
// of the "getCart" endpoint of the "store" service.
func NewGetCartNotFoundResponseBody(res *goa.ServiceError) *GetCartNotFoundResponseBody {
	body := &GetCartNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateUserNewUser builds a store service createUser endpoint payload.
func NewCreateUserNewUser(body *CreateUserRequestBody) *store.NewUser {
	v := &store.NewUser{
		Username:  *body.Username,
		Email:     *body.Email,
		FirstName: body.FirstName,
		LastName:  body.LastName,
	}

	return v
}

// NewGetUserPayload builds a store service getUser endpoint payload.
func NewGetUserPayload(id string) *store.GetUserPayload {
	v := &store.GetUserPayload{}
	v.ID = id

	return v
}

// NewCreateProductNewProduct builds a store service createProduct endpoint
// payload.
func NewCreateProductNewProduct(body *CreateProductRequestBody) *store.NewProduct {
	v := &store.NewProduct{
		Name:        *body.Name,
		Description: body.Description,
		Price:       *body.Price,
		Inventory:   *body.Inventory,
	}

	return v
}

// NewGetProductPayload builds a store service getProduct endpoint payload.
func NewGetProductPayload(id string) *store.GetProductPayload {
	v := &store.GetProductPayload{}
	v.ID = id

	return v
}

// NewCreateOrderNewOrder builds a store service createOrder endpoint payload.
func NewCreateOrderNewOrder(body *CreateOrderRequestBody) *store.NewOrder {
	v := &store.NewOrder{
		UserID: *body.UserID,
	}
	v.Items = make([]*store.OrderItem, len(body.Items))
	for i, val := range body.Items {
		v.Items[i] = unmarshalOrderItemRequestBodyToStoreOrderItem(val)
	}

	return v
}

// NewGetOrderPayload builds a store service getOrder endpoint payload.
func NewGetOrderPayload(id string) *store.GetOrderPayload {
	v := &store.GetOrderPayload{}
	v.ID = id

	return v
}

// NewAddToCartCartItem builds a store service addToCart endpoint payload.
func NewAddToCartCartItem(body *AddToCartRequestBody) *store.CartItem {
	v := &store.CartItem{
		UserID:    *body.UserID,
		ProductID: *body.ProductID,
		Quantity:  *body.Quantity,
		Price:     body.Price,
	}

	return v
}

// NewGetCartPayload builds a store service getCart endpoint payload.
func NewGetCartPayload(body *GetCartRequestBody) *store.GetCartPayload {
	v := &store.GetCartPayload{
		UserID: *body.UserID,
	}

	return v
}

// ValidateCreateUserRequestBody runs the validations defined on
// CreateUserRequestBody
func ValidateCreateUserRequestBody(body *CreateUserRequestBody) (err error) {
	if body.Username == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("username", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	return
}

// ValidateCreateProductRequestBody runs the validations defined on
// CreateProductRequestBody
func ValidateCreateProductRequestBody(body *CreateProductRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Price == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("price", "body"))
	}
	if body.Inventory == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("inventory", "body"))
	}
	return
}

// ValidateCreateOrderRequestBody runs the validations defined on
// CreateOrderRequestBody
func ValidateCreateOrderRequestBody(body *CreateOrderRequestBody) (err error) {
	if body.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("userID", "body"))
	}
	if body.Items == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("items", "body"))
	}
	for _, e := range body.Items {
		if e != nil {
			if err2 := ValidateOrderItemRequestBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateAddToCartRequestBody runs the validations defined on
// AddToCartRequestBody
func ValidateAddToCartRequestBody(body *AddToCartRequestBody) (err error) {
	if body.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("userID", "body"))
	}
	if body.ProductID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("productID", "body"))
	}
	if body.Quantity == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("quantity", "body"))
	}
	return
}

// ValidateGetCartRequestBody runs the validations defined on GetCartRequestBody
func ValidateGetCartRequestBody(body *GetCartRequestBody) (err error) {
	if body.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("userID", "body"))
	}
	return
}

// ValidateOrderItemRequestBody runs the validations defined on
// OrderItemRequestBody
func ValidateOrderItemRequestBody(body *OrderItemRequestBody) (err error) {
	if body.ProductID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("productID", "body"))
	}
	if body.Quantity == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("quantity", "body"))
	}
	if body.Price == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("price", "body"))
	}
	return
}

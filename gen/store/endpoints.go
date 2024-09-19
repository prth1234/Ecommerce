// Code generated by goa v3.19.0, DO NOT EDIT.
//
// store endpoints
//
// Command:
// $ goa gen store/design

package store

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "store" service endpoints.
type Endpoints struct {
	CreateUser    goa.Endpoint
	LoginUser     goa.Endpoint
	GetUser       goa.Endpoint
	GetUserAll    goa.Endpoint
	UpdateUser    goa.Endpoint
	CreateProduct goa.Endpoint
	GetProduct    goa.Endpoint
	ListProducts  goa.Endpoint
	CreateOrder   goa.Endpoint
	GetOrder      goa.Endpoint
	GetUserOrders goa.Endpoint
	AddToCart     goa.Endpoint
	GetCart       goa.Endpoint
}

// NewEndpoints wraps the methods of the "store" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		CreateUser:    NewCreateUserEndpoint(s),
		LoginUser:     NewLoginUserEndpoint(s),
		GetUser:       NewGetUserEndpoint(s),
		GetUserAll:    NewGetUserAllEndpoint(s),
		UpdateUser:    NewUpdateUserEndpoint(s),
		CreateProduct: NewCreateProductEndpoint(s),
		GetProduct:    NewGetProductEndpoint(s),
		ListProducts:  NewListProductsEndpoint(s),
		CreateOrder:   NewCreateOrderEndpoint(s),
		GetOrder:      NewGetOrderEndpoint(s),
		GetUserOrders: NewGetUserOrdersEndpoint(s),
		AddToCart:     NewAddToCartEndpoint(s),
		GetCart:       NewGetCartEndpoint(s),
	}
}

// Use applies the given middleware to all the "store" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.CreateUser = m(e.CreateUser)
	e.LoginUser = m(e.LoginUser)
	e.GetUser = m(e.GetUser)
	e.GetUserAll = m(e.GetUserAll)
	e.UpdateUser = m(e.UpdateUser)
	e.CreateProduct = m(e.CreateProduct)
	e.GetProduct = m(e.GetProduct)
	e.ListProducts = m(e.ListProducts)
	e.CreateOrder = m(e.CreateOrder)
	e.GetOrder = m(e.GetOrder)
	e.GetUserOrders = m(e.GetUserOrders)
	e.AddToCart = m(e.AddToCart)
	e.GetCart = m(e.GetCart)
}

// NewCreateUserEndpoint returns an endpoint function that calls the method
// "createUser" of service "store".
func NewCreateUserEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*NewUser)
		return s.CreateUser(ctx, p)
	}
}

// NewLoginUserEndpoint returns an endpoint function that calls the method
// "loginUser" of service "store".
func NewLoginUserEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*LoginUserPayload)
		return s.LoginUser(ctx, p)
	}
}

// NewGetUserEndpoint returns an endpoint function that calls the method
// "getUser" of service "store".
func NewGetUserEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*GetUserPayload)
		return s.GetUser(ctx, p)
	}
}

// NewGetUserAllEndpoint returns an endpoint function that calls the method
// "getUserAll" of service "store".
func NewGetUserAllEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		return s.GetUserAll(ctx)
	}
}

// NewUpdateUserEndpoint returns an endpoint function that calls the method
// "updateUser" of service "store".
func NewUpdateUserEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*UserUpdatePayload)
		return s.UpdateUser(ctx, p)
	}
}

// NewCreateProductEndpoint returns an endpoint function that calls the method
// "createProduct" of service "store".
func NewCreateProductEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*NewProduct)
		return s.CreateProduct(ctx, p)
	}
}

// NewGetProductEndpoint returns an endpoint function that calls the method
// "getProduct" of service "store".
func NewGetProductEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*GetProductPayload)
		return s.GetProduct(ctx, p)
	}
}

// NewListProductsEndpoint returns an endpoint function that calls the method
// "listProducts" of service "store".
func NewListProductsEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		return s.ListProducts(ctx)
	}
}

// NewCreateOrderEndpoint returns an endpoint function that calls the method
// "createOrder" of service "store".
func NewCreateOrderEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*NewOrder)
		return s.CreateOrder(ctx, p)
	}
}

// NewGetOrderEndpoint returns an endpoint function that calls the method
// "getOrder" of service "store".
func NewGetOrderEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*GetOrderPayload)
		return s.GetOrder(ctx, p)
	}
}

// NewGetUserOrdersEndpoint returns an endpoint function that calls the method
// "getUserOrders" of service "store".
func NewGetUserOrdersEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*GetUserOrdersPayload)
		return s.GetUserOrders(ctx, p)
	}
}

// NewAddToCartEndpoint returns an endpoint function that calls the method
// "addToCart" of service "store".
func NewAddToCartEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*CartItem)
		return s.AddToCart(ctx, p)
	}
}

// NewGetCartEndpoint returns an endpoint function that calls the method
// "getCart" of service "store".
func NewGetCartEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*GetCartPayload)
		return s.GetCart(ctx, p)
	}
}

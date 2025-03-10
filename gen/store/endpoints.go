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
	CreateUser              goa.Endpoint
	LoginUser               goa.Endpoint
	GetUser                 goa.Endpoint
	GetUserAll              goa.Endpoint
	UpdateUser              goa.Endpoint
	DeleteUser              goa.Endpoint
	CreateProduct           goa.Endpoint
	GetProduct              goa.Endpoint
	ListProducts            goa.Endpoint
	AddToCart               goa.Endpoint
	RemoveFromCart          goa.Endpoint
	GetCart                 goa.Endpoint
	CreateOrder             goa.Endpoint
	DeleteOrder             goa.Endpoint
	GetOrder                goa.Endpoint
	GetUserOrders           goa.Endpoint
	GetProductsPostedByUser goa.Endpoint
	UpdateOrderItemStatus   goa.Endpoint
}

// NewEndpoints wraps the methods of the "store" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		CreateUser:              NewCreateUserEndpoint(s),
		LoginUser:               NewLoginUserEndpoint(s),
		GetUser:                 NewGetUserEndpoint(s),
		GetUserAll:              NewGetUserAllEndpoint(s),
		UpdateUser:              NewUpdateUserEndpoint(s),
		DeleteUser:              NewDeleteUserEndpoint(s),
		CreateProduct:           NewCreateProductEndpoint(s),
		GetProduct:              NewGetProductEndpoint(s),
		ListProducts:            NewListProductsEndpoint(s),
		AddToCart:               NewAddToCartEndpoint(s),
		RemoveFromCart:          NewRemoveFromCartEndpoint(s),
		GetCart:                 NewGetCartEndpoint(s),
		CreateOrder:             NewCreateOrderEndpoint(s),
		DeleteOrder:             NewDeleteOrderEndpoint(s),
		GetOrder:                NewGetOrderEndpoint(s),
		GetUserOrders:           NewGetUserOrdersEndpoint(s),
		GetProductsPostedByUser: NewGetProductsPostedByUserEndpoint(s),
		UpdateOrderItemStatus:   NewUpdateOrderItemStatusEndpoint(s),
	}
}

// Use applies the given middleware to all the "store" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.CreateUser = m(e.CreateUser)
	e.LoginUser = m(e.LoginUser)
	e.GetUser = m(e.GetUser)
	e.GetUserAll = m(e.GetUserAll)
	e.UpdateUser = m(e.UpdateUser)
	e.DeleteUser = m(e.DeleteUser)
	e.CreateProduct = m(e.CreateProduct)
	e.GetProduct = m(e.GetProduct)
	e.ListProducts = m(e.ListProducts)
	e.AddToCart = m(e.AddToCart)
	e.RemoveFromCart = m(e.RemoveFromCart)
	e.GetCart = m(e.GetCart)
	e.CreateOrder = m(e.CreateOrder)
	e.DeleteOrder = m(e.DeleteOrder)
	e.GetOrder = m(e.GetOrder)
	e.GetUserOrders = m(e.GetUserOrders)
	e.GetProductsPostedByUser = m(e.GetProductsPostedByUser)
	e.UpdateOrderItemStatus = m(e.UpdateOrderItemStatus)
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

// NewDeleteUserEndpoint returns an endpoint function that calls the method
// "deleteUser" of service "store".
func NewDeleteUserEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		return nil, s.DeleteUser(ctx)
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
		p := req.(*ListProductsPayload)
		return s.ListProducts(ctx, p)
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

// NewRemoveFromCartEndpoint returns an endpoint function that calls the method
// "removeFromCart" of service "store".
func NewRemoveFromCartEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*RemoveFromCartPayload)
		return s.RemoveFromCart(ctx, p)
	}
}

// NewGetCartEndpoint returns an endpoint function that calls the method
// "getCart" of service "store".
func NewGetCartEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		return s.GetCart(ctx)
	}
}

// NewCreateOrderEndpoint returns an endpoint function that calls the method
// "createOrder" of service "store".
func NewCreateOrderEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		return s.CreateOrder(ctx)
	}
}

// NewDeleteOrderEndpoint returns an endpoint function that calls the method
// "deleteOrder" of service "store".
func NewDeleteOrderEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*DeleteOrderPayload)
		return nil, s.DeleteOrder(ctx, p)
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
		return s.GetUserOrders(ctx)
	}
}

// NewGetProductsPostedByUserEndpoint returns an endpoint function that calls
// the method "getProductsPostedByUser" of service "store".
func NewGetProductsPostedByUserEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		return s.GetProductsPostedByUser(ctx)
	}
}

// NewUpdateOrderItemStatusEndpoint returns an endpoint function that calls the
// method "updateOrderItemStatus" of service "store".
func NewUpdateOrderItemStatusEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*UpdateOrderItemStatusPayload)
		return s.UpdateOrderItemStatus(ctx, p)
	}
}

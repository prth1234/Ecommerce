// Code generated by goa v3.19.0, DO NOT EDIT.
//
// store client
//
// Command:
// $ goa gen store/design

package store

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "store" service client.
type Client struct {
	CreateUserEndpoint              goa.Endpoint
	LoginUserEndpoint               goa.Endpoint
	GetUserEndpoint                 goa.Endpoint
	GetUserAllEndpoint              goa.Endpoint
	UpdateUserEndpoint              goa.Endpoint
	DeleteUserEndpoint              goa.Endpoint
	CreateProductEndpoint           goa.Endpoint
	GetProductEndpoint              goa.Endpoint
	ListProductsEndpoint            goa.Endpoint
	AddToCartEndpoint               goa.Endpoint
	RemoveFromCartEndpoint          goa.Endpoint
	GetCartEndpoint                 goa.Endpoint
	CreateOrderEndpoint             goa.Endpoint
	DeleteOrderEndpoint             goa.Endpoint
	GetOrderEndpoint                goa.Endpoint
	GetUserOrdersEndpoint           goa.Endpoint
	GetProductsPostedByUserEndpoint goa.Endpoint
	UpdateOrderItemStatusEndpoint   goa.Endpoint
}

// NewClient initializes a "store" service client given the endpoints.
func NewClient(createUser, loginUser, getUser, getUserAll, updateUser, deleteUser, createProduct, getProduct, listProducts, addToCart, removeFromCart, getCart, createOrder, deleteOrder, getOrder, getUserOrders, getProductsPostedByUser, updateOrderItemStatus goa.Endpoint) *Client {
	return &Client{
		CreateUserEndpoint:              createUser,
		LoginUserEndpoint:               loginUser,
		GetUserEndpoint:                 getUser,
		GetUserAllEndpoint:              getUserAll,
		UpdateUserEndpoint:              updateUser,
		DeleteUserEndpoint:              deleteUser,
		CreateProductEndpoint:           createProduct,
		GetProductEndpoint:              getProduct,
		ListProductsEndpoint:            listProducts,
		AddToCartEndpoint:               addToCart,
		RemoveFromCartEndpoint:          removeFromCart,
		GetCartEndpoint:                 getCart,
		CreateOrderEndpoint:             createOrder,
		DeleteOrderEndpoint:             deleteOrder,
		GetOrderEndpoint:                getOrder,
		GetUserOrdersEndpoint:           getUserOrders,
		GetProductsPostedByUserEndpoint: getProductsPostedByUser,
		UpdateOrderItemStatusEndpoint:   updateOrderItemStatus,
	}
}

// CreateUser calls the "createUser" endpoint of the "store" service.
func (c *Client) CreateUser(ctx context.Context, p *NewUser) (res *User, err error) {
	var ires any
	ires, err = c.CreateUserEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*User), nil
}

// LoginUser calls the "loginUser" endpoint of the "store" service.
// LoginUser may return the following errors:
//   - "unauthorized" (type Unauthorized)
//   - error: internal error
func (c *Client) LoginUser(ctx context.Context, p *LoginUserPayload) (res *LoginUserResult, err error) {
	var ires any
	ires, err = c.LoginUserEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*LoginUserResult), nil
}

// GetUser calls the "getUser" endpoint of the "store" service.
// GetUser may return the following errors:
//   - "not-found" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) GetUser(ctx context.Context, p *GetUserPayload) (res *User, err error) {
	var ires any
	ires, err = c.GetUserEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*User), nil
}

// GetUserAll calls the "getUserAll" endpoint of the "store" service.
func (c *Client) GetUserAll(ctx context.Context) (res []*User, err error) {
	var ires any
	ires, err = c.GetUserAllEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.([]*User), nil
}

// UpdateUser calls the "updateUser" endpoint of the "store" service.
func (c *Client) UpdateUser(ctx context.Context, p *UserUpdatePayload) (res *User, err error) {
	var ires any
	ires, err = c.UpdateUserEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*User), nil
}

// DeleteUser calls the "deleteUser" endpoint of the "store" service.
func (c *Client) DeleteUser(ctx context.Context) (err error) {
	_, err = c.DeleteUserEndpoint(ctx, nil)
	return
}

// CreateProduct calls the "createProduct" endpoint of the "store" service.
func (c *Client) CreateProduct(ctx context.Context, p *NewProduct) (res *Product, err error) {
	var ires any
	ires, err = c.CreateProductEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Product), nil
}

// GetProduct calls the "getProduct" endpoint of the "store" service.
// GetProduct may return the following errors:
//   - "not-found" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) GetProduct(ctx context.Context, p *GetProductPayload) (res *Product, err error) {
	var ires any
	ires, err = c.GetProductEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Product), nil
}

// ListProducts calls the "listProducts" endpoint of the "store" service.
func (c *Client) ListProducts(ctx context.Context, p *ListProductsPayload) (res []*Product, err error) {
	var ires any
	ires, err = c.ListProductsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.([]*Product), nil
}

// AddToCart calls the "addToCart" endpoint of the "store" service.
func (c *Client) AddToCart(ctx context.Context, p *CartItem) (res *Cart, err error) {
	var ires any
	ires, err = c.AddToCartEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Cart), nil
}

// RemoveFromCart calls the "removeFromCart" endpoint of the "store" service.
func (c *Client) RemoveFromCart(ctx context.Context, p *RemoveFromCartPayload) (res *Cart, err error) {
	var ires any
	ires, err = c.RemoveFromCartEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Cart), nil
}

// GetCart calls the "getCart" endpoint of the "store" service.
// GetCart may return the following errors:
//   - "not-found" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) GetCart(ctx context.Context) (res *Cart, err error) {
	var ires any
	ires, err = c.GetCartEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*Cart), nil
}

// CreateOrder calls the "createOrder" endpoint of the "store" service.
func (c *Client) CreateOrder(ctx context.Context) (res *Order, err error) {
	var ires any
	ires, err = c.CreateOrderEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*Order), nil
}

// DeleteOrder calls the "deleteOrder" endpoint of the "store" service.
func (c *Client) DeleteOrder(ctx context.Context, p *DeleteOrderPayload) (err error) {
	_, err = c.DeleteOrderEndpoint(ctx, p)
	return
}

// GetOrder calls the "getOrder" endpoint of the "store" service.
// GetOrder may return the following errors:
//   - "not-found" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) GetOrder(ctx context.Context, p *GetOrderPayload) (res *Order, err error) {
	var ires any
	ires, err = c.GetOrderEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Order), nil
}

// GetUserOrders calls the "getUserOrders" endpoint of the "store" service.
func (c *Client) GetUserOrders(ctx context.Context) (res []*Order, err error) {
	var ires any
	ires, err = c.GetUserOrdersEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.([]*Order), nil
}

// GetProductsPostedByUser calls the "getProductsPostedByUser" endpoint of the
// "store" service.
func (c *Client) GetProductsPostedByUser(ctx context.Context) (res []*Product, err error) {
	var ires any
	ires, err = c.GetProductsPostedByUserEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.([]*Product), nil
}

// UpdateOrderItemStatus calls the "updateOrderItemStatus" endpoint of the
// "store" service.
// UpdateOrderItemStatus may return the following errors:
//   - "not-found" (type NotFound)
//   - "forbidden" (type Forbidden)
//   - error: internal error
func (c *Client) UpdateOrderItemStatus(ctx context.Context, p *UpdateOrderItemStatusPayload) (res *Order, err error) {
	var ires any
	ires, err = c.UpdateOrderItemStatusEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Order), nil
}

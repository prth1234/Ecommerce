// Code generated by goa v3.19.0, DO NOT EDIT.
//
// store HTTP server encoders and decoders
//
// Command:
// $ goa gen store/design

package server

import (
	"context"
	"errors"
	"io"
	"net/http"
	store "store/gen/store"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeCreateUserResponse returns an encoder for responses returned by the
// store createUser endpoint.
func EncodeCreateUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*store.User)
		enc := encoder(ctx, w)
		body := NewCreateUserResponseBody(res)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeCreateUserRequest returns a decoder for requests sent to the store
// createUser endpoint.
func DecodeCreateUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body CreateUserRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			var gerr *goa.ServiceError
			if errors.As(err, &gerr) {
				return nil, gerr
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateUserRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewCreateUserNewUser(&body)

		return payload, nil
	}
}

// EncodeLoginUserResponse returns an encoder for responses returned by the
// store loginUser endpoint.
func EncodeLoginUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*store.LoginUserResult)
		enc := encoder(ctx, w)
		body := NewLoginUserResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeLoginUserRequest returns a decoder for requests sent to the store
// loginUser endpoint.
func DecodeLoginUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body LoginUserRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			var gerr *goa.ServiceError
			if errors.As(err, &gerr) {
				return nil, gerr
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateLoginUserRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewLoginUserPayload(&body)

		return payload, nil
	}
}

// EncodeLoginUserError returns an encoder for errors returned by the loginUser
// store endpoint.
func EncodeLoginUserError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "unauthorized":
			var res store.Unauthorized
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeGetUserResponse returns an encoder for responses returned by the store
// getUser endpoint.
func EncodeGetUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*store.User)
		enc := encoder(ctx, w)
		body := NewGetUserResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetUserRequest returns a decoder for requests sent to the store
// getUser endpoint.
func DecodeGetUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]
		payload := NewGetUserPayload(id)

		return payload, nil
	}
}

// EncodeGetUserError returns an encoder for errors returned by the getUser
// store endpoint.
func EncodeGetUserError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "not-found":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewGetUserNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeGetUserAllResponse returns an encoder for responses returned by the
// store getUserAll endpoint.
func EncodeGetUserAllResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.([]*store.User)
		enc := encoder(ctx, w)
		body := NewGetUserAllResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeUpdateUserResponse returns an encoder for responses returned by the
// store updateUser endpoint.
func EncodeUpdateUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*store.User)
		enc := encoder(ctx, w)
		body := NewUpdateUserResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeUpdateUserRequest returns a decoder for requests sent to the store
// updateUser endpoint.
func DecodeUpdateUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body UpdateUserRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			var gerr *goa.ServiceError
			if errors.As(err, &gerr) {
				return nil, gerr
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUpdateUserRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewUpdateUserUserUpdatePayload(&body)

		return payload, nil
	}
}

// EncodeDeleteUserResponse returns an encoder for responses returned by the
// store deleteUser endpoint.
func EncodeDeleteUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// EncodeCreateProductResponse returns an encoder for responses returned by the
// store createProduct endpoint.
func EncodeCreateProductResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*store.Product)
		enc := encoder(ctx, w)
		body := NewCreateProductResponseBody(res)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeCreateProductRequest returns a decoder for requests sent to the store
// createProduct endpoint.
func DecodeCreateProductRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body CreateProductRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			var gerr *goa.ServiceError
			if errors.As(err, &gerr) {
				return nil, gerr
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateProductRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewCreateProductNewProduct(&body)

		return payload, nil
	}
}

// EncodeGetProductResponse returns an encoder for responses returned by the
// store getProduct endpoint.
func EncodeGetProductResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*store.Product)
		enc := encoder(ctx, w)
		body := NewGetProductResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetProductRequest returns a decoder for requests sent to the store
// getProduct endpoint.
func DecodeGetProductRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]
		payload := NewGetProductPayload(id)

		return payload, nil
	}
}

// EncodeGetProductError returns an encoder for errors returned by the
// getProduct store endpoint.
func EncodeGetProductError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "not-found":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewGetProductNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeListProductsResponse returns an encoder for responses returned by the
// store listProducts endpoint.
func EncodeListProductsResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.([]*store.Product)
		enc := encoder(ctx, w)
		body := NewListProductsResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListProductsRequest returns a decoder for requests sent to the store
// listProducts endpoint.
func DecodeListProductsRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body ListProductsRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			var gerr *goa.ServiceError
			if errors.As(err, &gerr) {
				return nil, gerr
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		payload := NewListProductsPayload(&body)

		return payload, nil
	}
}

// EncodeAddToCartResponse returns an encoder for responses returned by the
// store addToCart endpoint.
func EncodeAddToCartResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*store.Cart)
		enc := encoder(ctx, w)
		body := NewAddToCartResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeAddToCartRequest returns a decoder for requests sent to the store
// addToCart endpoint.
func DecodeAddToCartRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body AddToCartRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			var gerr *goa.ServiceError
			if errors.As(err, &gerr) {
				return nil, gerr
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateAddToCartRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewAddToCartCartItem(&body)

		return payload, nil
	}
}

// EncodeRemoveFromCartResponse returns an encoder for responses returned by
// the store removeFromCart endpoint.
func EncodeRemoveFromCartResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*store.Cart)
		enc := encoder(ctx, w)
		body := NewRemoveFromCartResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeRemoveFromCartRequest returns a decoder for requests sent to the store
// removeFromCart endpoint.
func DecodeRemoveFromCartRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			productID string

			params = mux.Vars(r)
		)
		productID = params["productID"]
		payload := NewRemoveFromCartPayload(productID)

		return payload, nil
	}
}

// EncodeGetCartResponse returns an encoder for responses returned by the store
// getCart endpoint.
func EncodeGetCartResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*store.Cart)
		enc := encoder(ctx, w)
		body := NewGetCartResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeGetCartError returns an encoder for errors returned by the getCart
// store endpoint.
func EncodeGetCartError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "not-found":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewGetCartNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeCreateOrderResponse returns an encoder for responses returned by the
// store createOrder endpoint.
func EncodeCreateOrderResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*store.Order)
		enc := encoder(ctx, w)
		body := NewCreateOrderResponseBody(res)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// EncodeDeleteOrderResponse returns an encoder for responses returned by the
// store deleteOrder endpoint.
func EncodeDeleteOrderResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeDeleteOrderRequest returns a decoder for requests sent to the store
// deleteOrder endpoint.
func DecodeDeleteOrderRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]
		payload := NewDeleteOrderPayload(id)

		return payload, nil
	}
}

// EncodeGetOrderResponse returns an encoder for responses returned by the
// store getOrder endpoint.
func EncodeGetOrderResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*store.Order)
		enc := encoder(ctx, w)
		body := NewGetOrderResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetOrderRequest returns a decoder for requests sent to the store
// getOrder endpoint.
func DecodeGetOrderRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]
		payload := NewGetOrderPayload(id)

		return payload, nil
	}
}

// EncodeGetOrderError returns an encoder for errors returned by the getOrder
// store endpoint.
func EncodeGetOrderError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "not-found":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewGetOrderNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeGetUserOrdersResponse returns an encoder for responses returned by the
// store getUserOrders endpoint.
func EncodeGetUserOrdersResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.([]*store.Order)
		enc := encoder(ctx, w)
		body := NewGetUserOrdersResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeGetProductsPostedByUserResponse returns an encoder for responses
// returned by the store getProductsPostedByUser endpoint.
func EncodeGetProductsPostedByUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.([]*store.Product)
		enc := encoder(ctx, w)
		body := NewGetProductsPostedByUserResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeUpdateOrderItemStatusResponse returns an encoder for responses
// returned by the store updateOrderItemStatus endpoint.
func EncodeUpdateOrderItemStatusResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*store.Order)
		enc := encoder(ctx, w)
		body := NewUpdateOrderItemStatusResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeUpdateOrderItemStatusRequest returns a decoder for requests sent to
// the store updateOrderItemStatus endpoint.
func DecodeUpdateOrderItemStatusRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body UpdateOrderItemStatusRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			var gerr *goa.ServiceError
			if errors.As(err, &gerr) {
				return nil, gerr
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUpdateOrderItemStatusRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			orderID   string
			productID string

			params = mux.Vars(r)
		)
		orderID = params["orderID"]
		productID = params["productID"]
		payload := NewUpdateOrderItemStatusPayload(&body, orderID, productID)

		return payload, nil
	}
}

// EncodeUpdateOrderItemStatusError returns an encoder for errors returned by
// the updateOrderItemStatus store endpoint.
func EncodeUpdateOrderItemStatusError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "forbidden":
			var res store.Forbidden
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not-found":
			var res store.NotFound
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalStoreUserToUserResponse builds a value of type *UserResponse from a
// value of type *store.User.
func marshalStoreUserToUserResponse(v *store.User) *UserResponse {
	res := &UserResponse{
		ID:        v.ID,
		Username:  v.Username,
		Email:     v.Email,
		FirstName: v.FirstName,
		LastName:  v.LastName,
	}

	return res
}

// marshalStoreProductToProductResponse builds a value of type *ProductResponse
// from a value of type *store.Product.
func marshalStoreProductToProductResponse(v *store.Product) *ProductResponse {
	res := &ProductResponse{
		UserID:      v.UserID,
		ID:          v.ID,
		Name:        v.Name,
		Description: v.Description,
		Price:       v.Price,
		Inventory:   v.Inventory,
	}

	return res
}

// marshalStoreCartItemToCartItemResponseBody builds a value of type
// *CartItemResponseBody from a value of type *store.CartItem.
func marshalStoreCartItemToCartItemResponseBody(v *store.CartItem) *CartItemResponseBody {
	res := &CartItemResponseBody{
		ProductID: v.ProductID,
		Quantity:  v.Quantity,
	}

	return res
}

// marshalStoreOrderItemToOrderItemResponseBody builds a value of type
// *OrderItemResponseBody from a value of type *store.OrderItem.
func marshalStoreOrderItemToOrderItemResponseBody(v *store.OrderItem) *OrderItemResponseBody {
	res := &OrderItemResponseBody{
		ProductID: v.ProductID,
		SellerID:  v.SellerID,
		Quantity:  v.Quantity,
		Price:     v.Price,
		Status:    v.Status,
	}

	return res
}

// marshalStoreOrderToOrderResponse builds a value of type *OrderResponse from
// a value of type *store.Order.
func marshalStoreOrderToOrderResponse(v *store.Order) *OrderResponse {
	res := &OrderResponse{
		ID:            v.ID,
		UserID:        v.UserID,
		TotalAmount:   v.TotalAmount,
		OverallStatus: v.OverallStatus,
	}
	if v.Items != nil {
		res.Items = make([]*OrderItemResponse, len(v.Items))
		for i, val := range v.Items {
			res.Items[i] = marshalStoreOrderItemToOrderItemResponse(val)
		}
	} else {
		res.Items = []*OrderItemResponse{}
	}

	return res
}

// marshalStoreOrderItemToOrderItemResponse builds a value of type
// *OrderItemResponse from a value of type *store.OrderItem.
func marshalStoreOrderItemToOrderItemResponse(v *store.OrderItem) *OrderItemResponse {
	res := &OrderItemResponse{
		ProductID: v.ProductID,
		SellerID:  v.SellerID,
		Quantity:  v.Quantity,
		Price:     v.Price,
		Status:    v.Status,
	}

	return res
}

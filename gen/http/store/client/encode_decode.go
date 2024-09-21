// Code generated by goa v3.19.0, DO NOT EDIT.
//
// store HTTP client encoders and decoders
//
// Command:
// $ goa gen store/design

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	store "store/gen/store"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildCreateUserRequest instantiates a HTTP request object with method and
// path set to call the "store" service "createUser" endpoint
func (c *Client) BuildCreateUserRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateUserStorePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("store", "createUser", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateUserRequest returns an encoder for requests sent to the store
// createUser server.
func EncodeCreateUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*store.NewUser)
		if !ok {
			return goahttp.ErrInvalidType("store", "createUser", "*store.NewUser", v)
		}
		body := NewCreateUserRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("store", "createUser", err)
		}
		return nil
	}
}

// DecodeCreateUserResponse returns a decoder for responses returned by the
// store createUser endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeCreateUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body CreateUserResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("store", "createUser", err)
			}
			err = ValidateCreateUserResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("store", "createUser", err)
			}
			res := NewCreateUserUserCreated(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("store", "createUser", resp.StatusCode, string(body))
		}
	}
}

// BuildLoginUserRequest instantiates a HTTP request object with method and
// path set to call the "store" service "loginUser" endpoint
func (c *Client) BuildLoginUserRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: LoginUserStorePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("store", "loginUser", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeLoginUserRequest returns an encoder for requests sent to the store
// loginUser server.
func EncodeLoginUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*store.LoginUserPayload)
		if !ok {
			return goahttp.ErrInvalidType("store", "loginUser", "*store.LoginUserPayload", v)
		}
		body := NewLoginUserRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("store", "loginUser", err)
		}
		return nil
	}
}

// DecodeLoginUserResponse returns a decoder for responses returned by the
// store loginUser endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeLoginUserResponse may return the following errors:
//   - "unauthorized" (type store.Unauthorized): http.StatusUnauthorized
//   - error: internal error
func DecodeLoginUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body LoginUserResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("store", "loginUser", err)
			}
			res := NewLoginUserResultOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("store", "loginUser", err)
			}
			return nil, NewLoginUserUnauthorized(body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("store", "loginUser", resp.StatusCode, string(body))
		}
	}
}

// BuildGetUserRequest instantiates a HTTP request object with method and path
// set to call the "store" service "getUser" endpoint
func (c *Client) BuildGetUserRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*store.GetUserPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("store", "getUser", "*store.GetUserPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetUserStorePath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("store", "getUser", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetUserResponse returns a decoder for responses returned by the store
// getUser endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeGetUserResponse may return the following errors:
//   - "not-found" (type *goa.ServiceError): http.StatusNotFound
//   - error: internal error
func DecodeGetUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetUserResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("store", "getUser", err)
			}
			err = ValidateGetUserResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("store", "getUser", err)
			}
			res := NewGetUserUserOK(&body)
			return res, nil
		case http.StatusNotFound:
			var (
				body GetUserNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("store", "getUser", err)
			}
			err = ValidateGetUserNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("store", "getUser", err)
			}
			return nil, NewGetUserNotFound(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("store", "getUser", resp.StatusCode, string(body))
		}
	}
}

// BuildGetUserAllRequest instantiates a HTTP request object with method and
// path set to call the "store" service "getUserAll" endpoint
func (c *Client) BuildGetUserAllRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetUserAllStorePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("store", "getUserAll", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetUserAllResponse returns a decoder for responses returned by the
// store getUserAll endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeGetUserAllResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetUserAllResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("store", "getUserAll", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateUserResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("store", "getUserAll", err)
			}
			res := NewGetUserAllUserOK(body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("store", "getUserAll", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateUserRequest instantiates a HTTP request object with method and
// path set to call the "store" service "updateUser" endpoint
func (c *Client) BuildUpdateUserRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateUserStorePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("store", "updateUser", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateUserRequest returns an encoder for requests sent to the store
// updateUser server.
func EncodeUpdateUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*store.UserUpdatePayload)
		if !ok {
			return goahttp.ErrInvalidType("store", "updateUser", "*store.UserUpdatePayload", v)
		}
		body := NewUpdateUserRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("store", "updateUser", err)
		}
		return nil
	}
}

// DecodeUpdateUserResponse returns a decoder for responses returned by the
// store updateUser endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeUpdateUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UpdateUserResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("store", "updateUser", err)
			}
			err = ValidateUpdateUserResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("store", "updateUser", err)
			}
			res := NewUpdateUserUserOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("store", "updateUser", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteUserRequest instantiates a HTTP request object with method and
// path set to call the "store" service "deleteUser" endpoint
func (c *Client) BuildDeleteUserRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteUserStorePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("store", "deleteUser", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeDeleteUserResponse returns a decoder for responses returned by the
// store deleteUser endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeDeleteUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("store", "deleteUser", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateProductRequest instantiates a HTTP request object with method and
// path set to call the "store" service "createProduct" endpoint
func (c *Client) BuildCreateProductRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateProductStorePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("store", "createProduct", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateProductRequest returns an encoder for requests sent to the store
// createProduct server.
func EncodeCreateProductRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*store.NewProduct)
		if !ok {
			return goahttp.ErrInvalidType("store", "createProduct", "*store.NewProduct", v)
		}
		body := NewCreateProductRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("store", "createProduct", err)
		}
		return nil
	}
}

// DecodeCreateProductResponse returns a decoder for responses returned by the
// store createProduct endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeCreateProductResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body CreateProductResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("store", "createProduct", err)
			}
			err = ValidateCreateProductResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("store", "createProduct", err)
			}
			res := NewCreateProductProductCreated(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("store", "createProduct", resp.StatusCode, string(body))
		}
	}
}

// BuildGetProductRequest instantiates a HTTP request object with method and
// path set to call the "store" service "getProduct" endpoint
func (c *Client) BuildGetProductRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*store.GetProductPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("store", "getProduct", "*store.GetProductPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetProductStorePath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("store", "getProduct", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetProductResponse returns a decoder for responses returned by the
// store getProduct endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeGetProductResponse may return the following errors:
//   - "not-found" (type *goa.ServiceError): http.StatusNotFound
//   - error: internal error
func DecodeGetProductResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetProductResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("store", "getProduct", err)
			}
			err = ValidateGetProductResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("store", "getProduct", err)
			}
			res := NewGetProductProductOK(&body)
			return res, nil
		case http.StatusNotFound:
			var (
				body GetProductNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("store", "getProduct", err)
			}
			err = ValidateGetProductNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("store", "getProduct", err)
			}
			return nil, NewGetProductNotFound(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("store", "getProduct", resp.StatusCode, string(body))
		}
	}
}

// BuildListProductsRequest instantiates a HTTP request object with method and
// path set to call the "store" service "listProducts" endpoint
func (c *Client) BuildListProductsRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListProductsStorePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("store", "listProducts", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeListProductsResponse returns a decoder for responses returned by the
// store listProducts endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeListProductsResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListProductsResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("store", "listProducts", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateProductResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("store", "listProducts", err)
			}
			res := NewListProductsProductOK(body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("store", "listProducts", resp.StatusCode, string(body))
		}
	}
}

// BuildAddToCartRequest instantiates a HTTP request object with method and
// path set to call the "store" service "addToCart" endpoint
func (c *Client) BuildAddToCartRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AddToCartStorePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("store", "addToCart", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeAddToCartRequest returns an encoder for requests sent to the store
// addToCart server.
func EncodeAddToCartRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*store.CartItem)
		if !ok {
			return goahttp.ErrInvalidType("store", "addToCart", "*store.CartItem", v)
		}
		body := NewAddToCartRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("store", "addToCart", err)
		}
		return nil
	}
}

// DecodeAddToCartResponse returns a decoder for responses returned by the
// store addToCart endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeAddToCartResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body AddToCartResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("store", "addToCart", err)
			}
			err = ValidateAddToCartResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("store", "addToCart", err)
			}
			res := NewAddToCartCartOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("store", "addToCart", resp.StatusCode, string(body))
		}
	}
}

// BuildRemoveFromCartRequest instantiates a HTTP request object with method
// and path set to call the "store" service "removeFromCart" endpoint
func (c *Client) BuildRemoveFromCartRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		productID string
	)
	{
		p, ok := v.(*store.RemoveFromCartPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("store", "removeFromCart", "*store.RemoveFromCartPayload", v)
		}
		productID = p.ProductID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RemoveFromCartStorePath(productID)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("store", "removeFromCart", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeRemoveFromCartResponse returns a decoder for responses returned by the
// store removeFromCart endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeRemoveFromCartResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body RemoveFromCartResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("store", "removeFromCart", err)
			}
			err = ValidateRemoveFromCartResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("store", "removeFromCart", err)
			}
			res := NewRemoveFromCartCartOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("store", "removeFromCart", resp.StatusCode, string(body))
		}
	}
}

// BuildGetCartRequest instantiates a HTTP request object with method and path
// set to call the "store" service "getCart" endpoint
func (c *Client) BuildGetCartRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetCartStorePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("store", "getCart", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetCartResponse returns a decoder for responses returned by the store
// getCart endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeGetCartResponse may return the following errors:
//   - "not-found" (type *goa.ServiceError): http.StatusNotFound
//   - error: internal error
func DecodeGetCartResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetCartResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("store", "getCart", err)
			}
			err = ValidateGetCartResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("store", "getCart", err)
			}
			res := NewGetCartCartOK(&body)
			return res, nil
		case http.StatusNotFound:
			var (
				body GetCartNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("store", "getCart", err)
			}
			err = ValidateGetCartNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("store", "getCart", err)
			}
			return nil, NewGetCartNotFound(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("store", "getCart", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateOrderRequest instantiates a HTTP request object with method and
// path set to call the "store" service "createOrder" endpoint
func (c *Client) BuildCreateOrderRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateOrderStorePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("store", "createOrder", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeCreateOrderResponse returns a decoder for responses returned by the
// store createOrder endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeCreateOrderResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body CreateOrderResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("store", "createOrder", err)
			}
			err = ValidateCreateOrderResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("store", "createOrder", err)
			}
			res := NewCreateOrderOrderCreated(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("store", "createOrder", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteOrderRequest instantiates a HTTP request object with method and
// path set to call the "store" service "deleteOrder" endpoint
func (c *Client) BuildDeleteOrderRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*store.DeleteOrderPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("store", "deleteOrder", "*store.DeleteOrderPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteOrderStorePath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("store", "deleteOrder", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeDeleteOrderResponse returns a decoder for responses returned by the
// store deleteOrder endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeDeleteOrderResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("store", "deleteOrder", resp.StatusCode, string(body))
		}
	}
}

// BuildGetOrderRequest instantiates a HTTP request object with method and path
// set to call the "store" service "getOrder" endpoint
func (c *Client) BuildGetOrderRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*store.GetOrderPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("store", "getOrder", "*store.GetOrderPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetOrderStorePath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("store", "getOrder", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetOrderResponse returns a decoder for responses returned by the store
// getOrder endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeGetOrderResponse may return the following errors:
//   - "not-found" (type *goa.ServiceError): http.StatusNotFound
//   - error: internal error
func DecodeGetOrderResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetOrderResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("store", "getOrder", err)
			}
			err = ValidateGetOrderResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("store", "getOrder", err)
			}
			res := NewGetOrderOrderOK(&body)
			return res, nil
		case http.StatusNotFound:
			var (
				body GetOrderNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("store", "getOrder", err)
			}
			err = ValidateGetOrderNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("store", "getOrder", err)
			}
			return nil, NewGetOrderNotFound(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("store", "getOrder", resp.StatusCode, string(body))
		}
	}
}

// BuildGetUserOrdersRequest instantiates a HTTP request object with method and
// path set to call the "store" service "getUserOrders" endpoint
func (c *Client) BuildGetUserOrdersRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetUserOrdersStorePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("store", "getUserOrders", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetUserOrdersResponse returns a decoder for responses returned by the
// store getUserOrders endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeGetUserOrdersResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetUserOrdersResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("store", "getUserOrders", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateOrderResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("store", "getUserOrders", err)
			}
			res := NewGetUserOrdersOrderOK(body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("store", "getUserOrders", resp.StatusCode, string(body))
		}
	}
}

// BuildGetProductsPostedByUserRequest instantiates a HTTP request object with
// method and path set to call the "store" service "getProductsPostedByUser"
// endpoint
func (c *Client) BuildGetProductsPostedByUserRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetProductsPostedByUserStorePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("store", "getProductsPostedByUser", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetProductsPostedByUserResponse returns a decoder for responses
// returned by the store getProductsPostedByUser endpoint. restoreBody controls
// whether the response body should be restored after having been read.
func DecodeGetProductsPostedByUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetProductsPostedByUserResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("store", "getProductsPostedByUser", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateProductResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("store", "getProductsPostedByUser", err)
			}
			res := NewGetProductsPostedByUserProductOK(body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("store", "getProductsPostedByUser", resp.StatusCode, string(body))
		}
	}
}

// unmarshalUserResponseToStoreUser builds a value of type *store.User from a
// value of type *UserResponse.
func unmarshalUserResponseToStoreUser(v *UserResponse) *store.User {
	res := &store.User{
		ID:        *v.ID,
		Username:  *v.Username,
		Email:     *v.Email,
		FirstName: v.FirstName,
		LastName:  v.LastName,
	}

	return res
}

// unmarshalProductResponseToStoreProduct builds a value of type *store.Product
// from a value of type *ProductResponse.
func unmarshalProductResponseToStoreProduct(v *ProductResponse) *store.Product {
	res := &store.Product{
		UserID:      *v.UserID,
		ID:          *v.ID,
		Name:        *v.Name,
		Description: v.Description,
		Price:       *v.Price,
		Inventory:   *v.Inventory,
	}

	return res
}

// unmarshalCartItemResponseBodyToStoreCartItem builds a value of type
// *store.CartItem from a value of type *CartItemResponseBody.
func unmarshalCartItemResponseBodyToStoreCartItem(v *CartItemResponseBody) *store.CartItem {
	res := &store.CartItem{
		ProductID: *v.ProductID,
		Quantity:  *v.Quantity,
	}

	return res
}

// unmarshalOrderItemResponseBodyToStoreOrderItem builds a value of type
// *store.OrderItem from a value of type *OrderItemResponseBody.
func unmarshalOrderItemResponseBodyToStoreOrderItem(v *OrderItemResponseBody) *store.OrderItem {
	res := &store.OrderItem{
		ProductID: *v.ProductID,
		Quantity:  *v.Quantity,
		Price:     *v.Price,
	}

	return res
}

// unmarshalOrderResponseToStoreOrder builds a value of type *store.Order from
// a value of type *OrderResponse.
func unmarshalOrderResponseToStoreOrder(v *OrderResponse) *store.Order {
	res := &store.Order{
		ID:          *v.ID,
		UserID:      *v.UserID,
		TotalAmount: *v.TotalAmount,
		Status:      *v.Status,
	}
	res.Items = make([]*store.OrderItem, len(v.Items))
	for i, val := range v.Items {
		res.Items[i] = unmarshalOrderItemResponseToStoreOrderItem(val)
	}

	return res
}

// unmarshalOrderItemResponseToStoreOrderItem builds a value of type
// *store.OrderItem from a value of type *OrderItemResponse.
func unmarshalOrderItemResponseToStoreOrderItem(v *OrderItemResponse) *store.OrderItem {
	res := &store.OrderItem{
		ProductID: *v.ProductID,
		Quantity:  *v.Quantity,
		Price:     *v.Price,
	}

	return res
}

// Code generated by goa v3.19.0, DO NOT EDIT.
//
// store HTTP client CLI support package
//
// Command:
// $ goa gen store/design

package client

import (
	"encoding/json"
	"fmt"
	store "store/gen/store"

	goa "goa.design/goa/v3/pkg"
)

// BuildCreateUserPayload builds the payload for the store createUser endpoint
// from CLI flags.
func BuildCreateUserPayload(storeCreateUserBody string) (*store.NewUser, error) {
	var err error
	var body CreateUserRequestBody
	{
		err = json.Unmarshal([]byte(storeCreateUserBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"Pariatur autem labore nam.\",\n      \"firstName\": \"Placeat sapiente inventore quis omnis facilis.\",\n      \"lastName\": \"Nulla eius qui culpa.\",\n      \"password\": \"Et repudiandae.\",\n      \"username\": \"Sit velit molestias non ea.\"\n   }'")
		}
	}
	v := &store.NewUser{
		Username:  body.Username,
		Email:     body.Email,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Password:  body.Password,
	}

	return v, nil
}

// BuildLoginUserPayload builds the payload for the store loginUser endpoint
// from CLI flags.
func BuildLoginUserPayload(storeLoginUserBody string) (*store.LoginUserPayload, error) {
	var err error
	var body LoginUserRequestBody
	{
		err = json.Unmarshal([]byte(storeLoginUserBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"password\": \"Sed voluptatem aspernatur.\",\n      \"username\": \"Praesentium et aut.\"\n   }'")
		}
	}
	v := &store.LoginUserPayload{
		Username: body.Username,
		Password: body.Password,
	}

	return v, nil
}

// BuildGetUserPayload builds the payload for the store getUser endpoint from
// CLI flags.
func BuildGetUserPayload(storeGetUserID string) (*store.GetUserPayload, error) {
	var id string
	{
		id = storeGetUserID
	}
	v := &store.GetUserPayload{}
	v.ID = id

	return v, nil
}

// BuildCreateProductPayload builds the payload for the store createProduct
// endpoint from CLI flags.
func BuildCreateProductPayload(storeCreateProductBody string) (*store.NewProduct, error) {
	var err error
	var body CreateProductRequestBody
	{
		err = json.Unmarshal([]byte(storeCreateProductBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"description\": \"Aperiam et quia atque illum quis.\",\n      \"inventory\": 2719289908079484285,\n      \"name\": \"Optio fugit facilis harum omnis.\",\n      \"price\": 0.6771694639392373\n   }'")
		}
	}
	v := &store.NewProduct{
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
		Inventory:   body.Inventory,
	}

	return v, nil
}

// BuildGetProductPayload builds the payload for the store getProduct endpoint
// from CLI flags.
func BuildGetProductPayload(storeGetProductID string) (*store.GetProductPayload, error) {
	var id string
	{
		id = storeGetProductID
	}
	v := &store.GetProductPayload{}
	v.ID = id

	return v, nil
}

// BuildCreateOrderPayload builds the payload for the store createOrder
// endpoint from CLI flags.
func BuildCreateOrderPayload(storeCreateOrderBody string) (*store.NewOrder, error) {
	var err error
	var body CreateOrderRequestBody
	{
		err = json.Unmarshal([]byte(storeCreateOrderBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"items\": [\n         {\n            \"price\": 0.73241831363072,\n            \"productID\": \"Eos libero non soluta.\",\n            \"quantity\": 9175415661056716192\n         },\n         {\n            \"price\": 0.73241831363072,\n            \"productID\": \"Eos libero non soluta.\",\n            \"quantity\": 9175415661056716192\n         },\n         {\n            \"price\": 0.73241831363072,\n            \"productID\": \"Eos libero non soluta.\",\n            \"quantity\": 9175415661056716192\n         }\n      ],\n      \"userID\": \"Sunt aut distinctio.\"\n   }'")
		}
		if body.Items == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("items", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &store.NewOrder{
		UserID: body.UserID,
	}
	if body.Items != nil {
		v.Items = make([]*store.OrderItem, len(body.Items))
		for i, val := range body.Items {
			v.Items[i] = marshalOrderItemRequestBodyToStoreOrderItem(val)
		}
	} else {
		v.Items = []*store.OrderItem{}
	}

	return v, nil
}

// BuildGetOrderPayload builds the payload for the store getOrder endpoint from
// CLI flags.
func BuildGetOrderPayload(storeGetOrderID string) (*store.GetOrderPayload, error) {
	var id string
	{
		id = storeGetOrderID
	}
	v := &store.GetOrderPayload{}
	v.ID = id

	return v, nil
}

// BuildGetUserOrdersPayload builds the payload for the store getUserOrders
// endpoint from CLI flags.
func BuildGetUserOrdersPayload(storeGetUserOrdersUserID string) (*store.GetUserOrdersPayload, error) {
	var userID string
	{
		userID = storeGetUserOrdersUserID
	}
	v := &store.GetUserOrdersPayload{}
	v.UserID = userID

	return v, nil
}

// BuildAddToCartPayload builds the payload for the store addToCart endpoint
// from CLI flags.
func BuildAddToCartPayload(storeAddToCartBody string) (*store.CartItem, error) {
	var err error
	var body AddToCartRequestBody
	{
		err = json.Unmarshal([]byte(storeAddToCartBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"price\": 0.962613919416865,\n      \"productID\": \"Alias maxime itaque beatae dicta reprehenderit.\",\n      \"quantity\": 6061566538106627186,\n      \"userID\": \"Harum illum molestiae.\"\n   }'")
		}
	}
	v := &store.CartItem{
		UserID:    body.UserID,
		ProductID: body.ProductID,
		Quantity:  body.Quantity,
		Price:     body.Price,
	}

	return v, nil
}

// BuildGetCartPayload builds the payload for the store getCart endpoint from
// CLI flags.
func BuildGetCartPayload(storeGetCartBody string) (*store.GetCartPayload, error) {
	var err error
	var body GetCartRequestBody
	{
		err = json.Unmarshal([]byte(storeGetCartBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"userID\": \"In aut reprehenderit sint.\"\n   }'")
		}
	}
	v := &store.GetCartPayload{
		UserID: body.UserID,
	}

	return v, nil
}

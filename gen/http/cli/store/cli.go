// Code generated by goa v3.19.0, DO NOT EDIT.
//
// store HTTP client CLI support package
//
// Command:
// $ goa gen store/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	storec "store/gen/http/store/client"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `store (create-user|login-user|get-user|get-user-all|create-product|get-product|list-products|create-order|get-order|get-user-orders|add-to-cart|get-cart)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` store create-user --body '{
      "email": "Accusantium sit velit molestias non ea dolor.",
      "firstName": "Autem labore nam.",
      "lastName": "Placeat sapiente inventore quis omnis facilis.",
      "password": "Nulla eius qui culpa.",
      "username": "Ducimus iure."
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, any, error) {
	var (
		storeFlags = flag.NewFlagSet("store", flag.ContinueOnError)

		storeCreateUserFlags    = flag.NewFlagSet("create-user", flag.ExitOnError)
		storeCreateUserBodyFlag = storeCreateUserFlags.String("body", "REQUIRED", "")

		storeLoginUserFlags    = flag.NewFlagSet("login-user", flag.ExitOnError)
		storeLoginUserBodyFlag = storeLoginUserFlags.String("body", "REQUIRED", "")

		storeGetUserFlags  = flag.NewFlagSet("get-user", flag.ExitOnError)
		storeGetUserIDFlag = storeGetUserFlags.String("id", "REQUIRED", "")

		storeGetUserAllFlags = flag.NewFlagSet("get-user-all", flag.ExitOnError)

		storeCreateProductFlags    = flag.NewFlagSet("create-product", flag.ExitOnError)
		storeCreateProductBodyFlag = storeCreateProductFlags.String("body", "REQUIRED", "")

		storeGetProductFlags  = flag.NewFlagSet("get-product", flag.ExitOnError)
		storeGetProductIDFlag = storeGetProductFlags.String("id", "REQUIRED", "")

		storeListProductsFlags = flag.NewFlagSet("list-products", flag.ExitOnError)

		storeCreateOrderFlags    = flag.NewFlagSet("create-order", flag.ExitOnError)
		storeCreateOrderBodyFlag = storeCreateOrderFlags.String("body", "REQUIRED", "")

		storeGetOrderFlags  = flag.NewFlagSet("get-order", flag.ExitOnError)
		storeGetOrderIDFlag = storeGetOrderFlags.String("id", "REQUIRED", "")

		storeGetUserOrdersFlags      = flag.NewFlagSet("get-user-orders", flag.ExitOnError)
		storeGetUserOrdersUserIDFlag = storeGetUserOrdersFlags.String("user-id", "REQUIRED", "")

		storeAddToCartFlags    = flag.NewFlagSet("add-to-cart", flag.ExitOnError)
		storeAddToCartBodyFlag = storeAddToCartFlags.String("body", "REQUIRED", "")

		storeGetCartFlags    = flag.NewFlagSet("get-cart", flag.ExitOnError)
		storeGetCartBodyFlag = storeGetCartFlags.String("body", "REQUIRED", "")
	)
	storeFlags.Usage = storeUsage
	storeCreateUserFlags.Usage = storeCreateUserUsage
	storeLoginUserFlags.Usage = storeLoginUserUsage
	storeGetUserFlags.Usage = storeGetUserUsage
	storeGetUserAllFlags.Usage = storeGetUserAllUsage
	storeCreateProductFlags.Usage = storeCreateProductUsage
	storeGetProductFlags.Usage = storeGetProductUsage
	storeListProductsFlags.Usage = storeListProductsUsage
	storeCreateOrderFlags.Usage = storeCreateOrderUsage
	storeGetOrderFlags.Usage = storeGetOrderUsage
	storeGetUserOrdersFlags.Usage = storeGetUserOrdersUsage
	storeAddToCartFlags.Usage = storeAddToCartUsage
	storeGetCartFlags.Usage = storeGetCartUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "store":
			svcf = storeFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "store":
			switch epn {
			case "create-user":
				epf = storeCreateUserFlags

			case "login-user":
				epf = storeLoginUserFlags

			case "get-user":
				epf = storeGetUserFlags

			case "get-user-all":
				epf = storeGetUserAllFlags

			case "create-product":
				epf = storeCreateProductFlags

			case "get-product":
				epf = storeGetProductFlags

			case "list-products":
				epf = storeListProductsFlags

			case "create-order":
				epf = storeCreateOrderFlags

			case "get-order":
				epf = storeGetOrderFlags

			case "get-user-orders":
				epf = storeGetUserOrdersFlags

			case "add-to-cart":
				epf = storeAddToCartFlags

			case "get-cart":
				epf = storeGetCartFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "store":
			c := storec.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "create-user":
				endpoint = c.CreateUser()
				data, err = storec.BuildCreateUserPayload(*storeCreateUserBodyFlag)
			case "login-user":
				endpoint = c.LoginUser()
				data, err = storec.BuildLoginUserPayload(*storeLoginUserBodyFlag)
			case "get-user":
				endpoint = c.GetUser()
				data, err = storec.BuildGetUserPayload(*storeGetUserIDFlag)
			case "get-user-all":
				endpoint = c.GetUserAll()
			case "create-product":
				endpoint = c.CreateProduct()
				data, err = storec.BuildCreateProductPayload(*storeCreateProductBodyFlag)
			case "get-product":
				endpoint = c.GetProduct()
				data, err = storec.BuildGetProductPayload(*storeGetProductIDFlag)
			case "list-products":
				endpoint = c.ListProducts()
			case "create-order":
				endpoint = c.CreateOrder()
				data, err = storec.BuildCreateOrderPayload(*storeCreateOrderBodyFlag)
			case "get-order":
				endpoint = c.GetOrder()
				data, err = storec.BuildGetOrderPayload(*storeGetOrderIDFlag)
			case "get-user-orders":
				endpoint = c.GetUserOrders()
				data, err = storec.BuildGetUserOrdersPayload(*storeGetUserOrdersUserIDFlag)
			case "add-to-cart":
				endpoint = c.AddToCart()
				data, err = storec.BuildAddToCartPayload(*storeAddToCartBodyFlag)
			case "get-cart":
				endpoint = c.GetCart()
				data, err = storec.BuildGetCartPayload(*storeGetCartBodyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// storeUsage displays the usage of the store command and its subcommands.
func storeUsage() {
	fmt.Fprintf(os.Stderr, `Store service
Usage:
    %[1]s [globalflags] store COMMAND [flags]

COMMAND:
    create-user: CreateUser implements createUser.
    login-user: Login a user and return a JWT token
    get-user: GetUser implements getUser.
    get-user-all: GetUserAll implements getUserAll.
    create-product: CreateProduct implements createProduct.
    get-product: GetProduct implements getProduct.
    list-products: ListProducts implements listProducts.
    create-order: CreateOrder implements createOrder.
    get-order: GetOrder implements getOrder.
    get-user-orders: Retrieve all orders for a specific user
    add-to-cart: AddToCart implements addToCart.
    get-cart: GetCart implements getCart.

Additional help:
    %[1]s store COMMAND --help
`, os.Args[0])
}
func storeCreateUserUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] store create-user -body JSON

CreateUser implements createUser.
    -body JSON: 

Example:
    %[1]s store create-user --body '{
      "email": "Accusantium sit velit molestias non ea dolor.",
      "firstName": "Autem labore nam.",
      "lastName": "Placeat sapiente inventore quis omnis facilis.",
      "password": "Nulla eius qui culpa.",
      "username": "Ducimus iure."
   }'
`, os.Args[0])
}

func storeLoginUserUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] store login-user -body JSON

Login a user and return a JWT token
    -body JSON: 

Example:
    %[1]s store login-user --body '{
      "password": "Praesentium et aut.",
      "username": "Quis itaque sed quo saepe."
   }'
`, os.Args[0])
}

func storeGetUserUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] store get-user -id STRING

GetUser implements getUser.
    -id STRING: 

Example:
    %[1]s store get-user --id "Nostrum ea quae sint qui."
`, os.Args[0])
}

func storeGetUserAllUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] store get-user-all

GetUserAll implements getUserAll.

Example:
    %[1]s store get-user-all
`, os.Args[0])
}

func storeCreateProductUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] store create-product -body JSON

CreateProduct implements createProduct.
    -body JSON: 

Example:
    %[1]s store create-product --body '{
      "description": "Dolor itaque quia nam et optio fugit.",
      "inventory": 6708158487146252491,
      "name": "Cupiditate assumenda doloribus ea porro laborum.",
      "price": 0.6889859441205155
   }'
`, os.Args[0])
}

func storeGetProductUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] store get-product -id STRING

GetProduct implements getProduct.
    -id STRING: 

Example:
    %[1]s store get-product --id "Et nihil nesciunt odio quis et in."
`, os.Args[0])
}

func storeListProductsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] store list-products

ListProducts implements listProducts.

Example:
    %[1]s store list-products
`, os.Args[0])
}

func storeCreateOrderUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] store create-order -body JSON

CreateOrder implements createOrder.
    -body JSON: 

Example:
    %[1]s store create-order --body '{
      "items": [
         {
            "price": 0.2356674159366931,
            "productID": "Aut distinctio cupiditate quis eos libero.",
            "quantity": 4210783297621910736
         },
         {
            "price": 0.2356674159366931,
            "productID": "Aut distinctio cupiditate quis eos libero.",
            "quantity": 4210783297621910736
         }
      ]
   }'
`, os.Args[0])
}

func storeGetOrderUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] store get-order -id STRING

GetOrder implements getOrder.
    -id STRING: 

Example:
    %[1]s store get-order --id "Qui dicta porro."
`, os.Args[0])
}

func storeGetUserOrdersUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] store get-user-orders -user-id STRING

Retrieve all orders for a specific user
    -user-id STRING: 

Example:
    %[1]s store get-user-orders --user-id "Pariatur ipsum."
`, os.Args[0])
}

func storeAddToCartUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] store add-to-cart -body JSON

AddToCart implements addToCart.
    -body JSON: 

Example:
    %[1]s store add-to-cart --body '{
      "price": 0.962613919416865,
      "productID": "Alias maxime itaque beatae dicta reprehenderit.",
      "quantity": 6061566538106627186,
      "userID": "Harum illum molestiae."
   }'
`, os.Args[0])
}

func storeGetCartUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] store get-cart -body JSON

GetCart implements getCart.
    -body JSON: 

Example:
    %[1]s store get-cart --body '{
      "userID": "In aut reprehenderit sint."
   }'
`, os.Args[0])
}

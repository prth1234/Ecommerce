package storeapi

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	store "store/gen/store"
)

type storesrvc struct {
	db *sql.DB
}

func NewStore(db *sql.DB) store.Service {
	return &storesrvc{db: db}
}

func (s *storesrvc) CreateUser(ctx context.Context, p *store.NewUser) (res *store.User, err error) {
	id := uuid.New().String()
	query := `INSERT INTO users (id, username, email, first_name, last_name) VALUES ($1, $2, $3, $4, $5) RETURNING id, username, email, first_name, last_name`

	user := &store.User{}
	err = s.db.QueryRowContext(ctx, query, id, p.Username, p.Email, p.FirstName, p.LastName).Scan(&user.ID, &user.Username, &user.Email, &user.FirstName, &user.LastName)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}

	return user, nil
}

func (s *storesrvc) GetUser(ctx context.Context, p *store.GetUserPayload) (res *store.User, err error) {
	query := `SELECT id, username, email, first_name, last_name FROM users WHERE id = $1`

	user := &store.User{}
	err = s.db.QueryRowContext(ctx, query, p.ID).Scan(&user.ID, &user.Username, &user.Email, &user.FirstName, &user.LastName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.MakeNotFound(fmt.Errorf("user not found"))
		}
		return nil, fmt.Errorf("error getting user: %v", err)
	}

	return user, nil
}

func (s *storesrvc) GetUserAll(ctx context.Context) (res []*store.User, err error) {
	query := `SELECT id, username, email, first_name, last_name FROM users`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error getting users: %v", err)
	}
	defer rows.Close()

	var users []*store.User
	for rows.Next() {
		user := &store.User{}
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.FirstName, &user.LastName); err != nil {
			return nil, fmt.Errorf("error scanning user: %v", err)
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *storesrvc) CreateProduct(ctx context.Context, p *store.NewProduct) (res *store.Product, err error) {
	id := uuid.New().String()
	query := `INSERT INTO products (id, name, description, price, inventory) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, description, price, inventory`

	product := &store.Product{}
	err = s.db.QueryRowContext(ctx, query, id, p.Name, p.Description, p.Price, p.Inventory).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Inventory)
	if err != nil {
		return nil, fmt.Errorf("error creating product: %v", err)
	}

	return product, nil
}

func (s *storesrvc) GetProduct(ctx context.Context, p *store.GetProductPayload) (res *store.Product, err error) {
	query := `SELECT id, name, description, price, inventory FROM products WHERE id = $1`

	product := &store.Product{}
	err = s.db.QueryRowContext(ctx, query, p.ID).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Inventory)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.MakeNotFound(fmt.Errorf("product not found"))
		}
		return nil, fmt.Errorf("error getting product: %v", err)
	}

	return product, nil
}

func (s *storesrvc) ListProducts(ctx context.Context) (res []*store.Product, err error) {
	query := `SELECT id, name, description, price, inventory FROM products`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error listing products: %v", err)
	}
	defer rows.Close()

	var products []*store.Product
	for rows.Next() {
		product := &store.Product{}
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Inventory); err != nil {
			return nil, fmt.Errorf("error scanning product: %v", err)
		}
		products = append(products, product)
	}

	return products, nil
}

func (s *storesrvc) CreateOrder(ctx context.Context, p *store.NewOrder) (res *store.Order, err error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error starting transaction: %v", err)
	}
	defer tx.Rollback()

	id := uuid.New().String()
	query := `INSERT INTO orders (id, user_id, total_amount, status) VALUES ($1, $2, $3, $4) RETURNING id, user_id, total_amount, status`

	order := &store.Order{Items: p.Items}
	err = tx.QueryRowContext(ctx, query, id, p.UserID, 0, "pending").Scan(&order.ID, &order.UserID, &order.TotalAmount, &order.Status)
	if err != nil {
		return nil, fmt.Errorf("error creating order: %v", err)
	}

	for _, item := range p.Items {
		_, err = tx.ExecContext(ctx, `INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)`,
			order.ID, item.ProductID, item.Quantity, item.Price)
		if err != nil {
			return nil, fmt.Errorf("error adding order item: %v", err)
		}
		order.TotalAmount += item.Price * float64(item.Quantity)
	}

	_, err = tx.ExecContext(ctx, `UPDATE orders SET total_amount = $1 WHERE id = $2`, order.TotalAmount, order.ID)
	if err != nil {
		return nil, fmt.Errorf("error updating order total: %v", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("error committing transaction: %v", err)
	}

	return order, nil
}

func (s *storesrvc) GetUserOrders(ctx context.Context, p *store.GetUserOrdersPayload) (res []*store.Order, err error) {
	query := `
		SELECT o.id, o.user_id, o.total_amount, o.status, 
			   oi.product_id, oi.quantity, oi.price
		FROM orders o
		LEFT JOIN order_items oi ON o.id = oi.order_id
		WHERE o.user_id = $1
		ORDER BY o.id`

	rows, err := s.db.QueryContext(ctx, query, p.UserID)
	if err != nil {
		return nil, fmt.Errorf("error getting user orders: %v", err)
	}
	defer rows.Close()

	orders := make(map[string]*store.Order)
	for rows.Next() {
		var orderID, userID, productID sql.NullString
		var totalAmount, price sql.NullFloat64
		var status sql.NullString
		var quantity sql.NullInt32

		err = rows.Scan(&orderID, &userID, &totalAmount, &status, &productID, &quantity, &price)
		if err != nil {
			return nil, fmt.Errorf("error scanning order: %v", err)
		}

		order, exists := orders[orderID.String]
		if !exists {
			order = &store.Order{
				ID:          orderID.String,
				UserID:      userID.String,
				TotalAmount: totalAmount.Float64,
				Status:      status.String,
				Items:       []*store.OrderItem{},
			}
			orders[orderID.String] = order
		}

		if productID.Valid {
			item := &store.OrderItem{
				ProductID: productID.String,
				Quantity:  int(quantity.Int32),
				Price:     price.Float64,
			}
			order.Items = append(order.Items, item)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	// Convert map to slice
	var result []*store.Order
	for _, order := range orders {
		result = append(result, order)
	}

	if len(result) == 0 {
		return nil, store.MakeNotFound(fmt.Errorf("no orders found for user"))
	}

	return result, nil
}

func (s *storesrvc) GetOrder(ctx context.Context, p *store.GetOrderPayload) (res *store.Order, err error) {
	query := `
		SELECT o.id, o.user_id, o.total_amount, o.status, 
			   oi.product_id, oi.quantity, oi.price
		FROM orders o
		LEFT JOIN order_items oi ON o.id = oi.order_id
		WHERE o.id = $1`

	rows, err := s.db.QueryContext(ctx, query, p.ID)
	if err != nil {
		return nil, fmt.Errorf("error getting order: %v", err)
	}
	defer rows.Close()

	var order *store.Order
	for rows.Next() {
		if order == nil {
			order = &store.Order{Items: []*store.OrderItem{}}
			var productID sql.NullString
			var quantity sql.NullInt32
			var price sql.NullFloat64
			err = rows.Scan(&order.ID, &order.UserID, &order.TotalAmount, &order.Status,
				&productID, &quantity, &price)
			if err != nil {
				return nil, fmt.Errorf("error scanning order: %v", err)
			}
			if productID.Valid {
				item := &store.OrderItem{
					ProductID: productID.String,
					Quantity:  int(quantity.Int32),
					Price:     price.Float64,
				}
				order.Items = append(order.Items, item)
			}
		} else {
			var productID sql.NullString
			var quantity sql.NullInt32
			var price sql.NullFloat64
			err = rows.Scan(new(sql.NullString), new(sql.NullString), new(sql.NullFloat64), new(sql.NullString),
				&productID, &quantity, &price)
			if err != nil {
				return nil, fmt.Errorf("error scanning order item: %v", err)
			}
			if productID.Valid {
				item := &store.OrderItem{
					ProductID: productID.String,
					Quantity:  int(quantity.Int32),
					Price:     price.Float64,
				}
				order.Items = append(order.Items, item)
			}
		}
	}

	if order == nil {
		return nil, store.MakeNotFound(fmt.Errorf("order not found"))
	}

	return order, nil
}

func (s *storesrvc) AddToCart(ctx context.Context, p *store.CartItem) (res *store.Cart, err error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error starting transaction: %v", err)
	}
	defer tx.Rollback()

	// Check if the cart exists, if not create it
	var cartID string
	err = tx.QueryRowContext(ctx, `SELECT id FROM carts WHERE user_id = $1`, p.UserID).Scan(&cartID)
	if err == sql.ErrNoRows {
		cartID = uuid.New().String()
		_, err = tx.ExecContext(ctx, `INSERT INTO carts (id, user_id) VALUES ($1, $2)`, cartID, p.UserID)
		if err != nil {
			return nil, fmt.Errorf("error creating cart: %v", err)
		}
	} else if err != nil {
		return nil, fmt.Errorf("error checking cart: %v", err)
	}

	// Add or update cart item
	_, err = tx.ExecContext(ctx, `
		INSERT INTO cart_items (cart_id, product_id, quantity)
		VALUES ($1, $2, $3)
		ON CONFLICT (cart_id, product_id) DO UPDATE SET quantity = cart_items.quantity + $3`,
		cartID, p.ProductID, p.Quantity)
	if err != nil {
		return nil, fmt.Errorf("error adding item to cart: %v", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("error committing transaction: %v", err)
	}

	return s.GetCart(ctx, &store.GetCartPayload{UserID: p.UserID})
}

func (s *storesrvc) GetCart(ctx context.Context, p *store.GetCartPayload) (res *store.Cart, err error) {
	query := `
		SELECT c.id, c.user_id, ci.product_id, ci.quantity, p.price
		FROM carts c
		LEFT JOIN cart_items ci ON c.id = ci.cart_id
		LEFT JOIN products p ON ci.product_id = p.id
		WHERE c.user_id = $1`

	rows, err := s.db.QueryContext(ctx, query, p.UserID)
	if err != nil {
		return nil, fmt.Errorf("error getting cart: %v", err)
	}
	defer rows.Close()

	cart := &store.Cart{UserID: p.UserID, Items: []*store.CartItem{}, TotalAmount: 0}
	for rows.Next() {
		var productID sql.NullString
		var quantity sql.NullInt32
		var price sql.NullFloat64
		err = rows.Scan(&cart.ID, &cart.UserID, &productID, &quantity, &price)
		if err != nil {
			return nil, fmt.Errorf("error scanning cart item: %v", err)
		}

		if productID.Valid && price.Valid && quantity.Valid {
			item := &store.CartItem{
				ProductID: productID.String,
				Quantity:  int(quantity.Int32),
			}
			cart.Items = append(cart.Items, item)
			if price.Valid && quantity.Valid {
				cart.TotalAmount += price.Float64 * float64(quantity.Int32)
			}

		}
	}

	if cart.ID == "" {
		return nil, store.MakeNotFound(fmt.Errorf("cart not found"))
	}

	return cart, nil
}

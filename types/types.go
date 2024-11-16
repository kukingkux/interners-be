package types

import "time"

type ProductStore interface {
	GetProducts() ([]Product, error)
	GetProductsById(ps []int) ([]Product, error)
	UpdateProduct(Product) error
	CreateProduct(CreateProductPayload) error
}

type Product struct {
	ID          int       `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Image       string    `json:"image,omitempty"`
	Price       float64   `json:"price,omitempty"`
	Quantity    int       `json:"quantity,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

type OrderStore interface {
	CreateOrder(Order) (int, error)
	CreateOrderItem(OrderItem) error
}

type Order struct {
	ID        int       `json:"id,omitempty"`
	UserID    int       `json:"user_id,omitempty"`
	Total     float64   `json:"total,omitempty"`
	Status    string    `json:"status,omitempty"`
	Address   string    `json:"address,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type OrderItem struct {
	ID        int       `json:"id,omitempty"`
	OrderID   int       `json:"order_id,omitempty"`
	ProductID int       `json:"product_id,omitempty"`
	Quantity  int       `json:"quantity,omitempty"`
	Price     float64   `json:"price,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(User) error
}

type User struct {
	ID        int       `json:"id,omitempty"`
	FirstName string    `json:"firstName,omitempty"`
	LastName  string    `json:"lastName,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type CreateProductPayload struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName,omitempty" validate:"required"`
	LastName  string `json:"lastName,omitempty" validate:"required"`
	Email     string `json:"email,omitempty" validate:"required,email"`
	Password  string `json:"password,omitempty" validate:"required,min=8,max=120"`
}
type LoginUserPayload struct {
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,min=8,max=120"`
}

type CartItem struct {
	ProductID int `json:"product_id,omitempty"`
	Quantity  int `json:"quantity,omitempty"`
}

type CartCheckoutPayload struct {
	Items []CartItem `json:"items,omitempty" validate:"required"`
}

package types

import "time"

type ProductStore interface {
	GetProducts() ([]Product, error)
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

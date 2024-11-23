package types

import "time"

type GoAuth struct {
	Id            string `json:"id,omitempty"`
	Email         string `json:"email,omitempty"`
	VerifiedEmail bool   `json:"verified_email,omitempty"`
	Name          string `json:"name,omitempty"`
	GivenName     string `json:"given_name,omitempty"`
	FamilyName    string `json:"family_name,omitempty"`
	Picture       string `json:"picture,omitempty"`
	Locale        string `json:"locale,omitempty"`
}

type PostStore interface {
	GetPosts() ([]*Post, error)
	GetPostById(id int) (*Post, error)
	GetPostsById(ids []int) ([]Post, error)
	UpdatePost(Post) error
	CreatePost(CreatePostPayload) error
}

type Post struct {
	ID          int       `json:"id,omitempty"`
	UserID      int       `json:"user_id,omitempty"`
	CompanyID   int       `json:"company_id,omitempty"`
	CompanyName string    `json:"company_name,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Requirement string    `json:"requirement,omitempty"`
	Salary      float64   `json:"salary,omitempty"`
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
	PostID    int       `json:"post_id,omitempty"`
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
	ID             int       `json:"id,omitempty"`
	FirstName      string    `json:"firstName,omitempty"`
	LastName       string    `json:"lastName,omitempty"`
	Email          string    `json:"email,omitempty"`
	Password       string    `json:"password,omitempty"`
	PhoneNumber    string    `json:"phone_number,omitempty"`
	ZipCode        string    `json:"zip_code,omitempty"`
	City           string    `json:"city,omitempty"`
	Address        string    `json:"address,omitempty"`
	CV             string    `json:"cv,omitempty"`
	ProfilePicture string    `json:"profile_picture,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}

type CreatePostPayload struct {
	Title       string  `json:"name,omitempty" validate:"required"`
	Description string  `json:"description,omitempty" validate:"required"`
	Requirement string  `json:"image,omitempty" validate:"required"`
	Salary      float64 `json:"salary,omitempty" validate:"required"`
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
	PostID   int `json:"post_id,omitempty"`
	Quantity int `json:"quantity,omitempty"`
}

type CartCheckoutPayload struct {
	Items []CartItem `json:"items,omitempty" validate:"required"`
}

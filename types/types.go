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

type CompanyStore interface {
	GetCompanyies() ([]*Company, error)
	GetCompanyById(id int) (*Company, error)
	GetCompanUserRoleById(ids []int) ([]Company, error)
	UpdateCompany(Company) error
	CreateCompany(CreateCompanyPayload) error
}

type Company struct {
	ID          int       `json:"id,omitempty"`
	UserID      int       `json:"user_id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Contact     int       `json:"contact,omitempty"`
	Email       string    `json:"email,omitempty"`
	Address     string    `json:"address,omitempty"`
	Province    string    `json:"province,omitempty"`
	City        string    `json:"city,omitempty"`
	Logo        string    `json:"logo,omitempty"`
	Banner      string    `json:"banner,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

type UserRoleStore interface {
	GetUserRoles() ([]*UserRole, error)
	GetUserRoleById(id int) (*UserRole, error)
	GetUserRolesById(ids []int) ([]UserRole, error)
	UpdateUserRole(UserRole) error
	CreateUserRole(CreateUserRolePayload) error
}

type UserRole struct {
	ID        int       `json:"id,omitempty"`
	UserID    int       `json:"user_id,omitempty"`
	RoleID    int       `json:"role_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type RolePermission struct {
	ID           int       `json:"id,omitempty"`
	RoleID       int       `json:"role_id,omitempty"`
	PermissionID int       `json:"permission_id,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}

type Role struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type Permission struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
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

type CreateCompanyPayload struct {
	Name        string `json:"name,omitempty" validate:"required"`
	Description string `json:"description,omitempty"`
	Contact     int    `json:"contact,omitempty"`
	Email       string `json:"email,omitempty"`
	Address     string `json:"address,omitempty"`
	Province    string `json:"province,omitempty"`
	City        string `json:"city,omitempty"`
	Logo        string `json:"logo,omitempty"`
	Banner      string `json:"banner,omitempty"`
}

type CreateUserRolePayload struct {
	UserId int `json:"user_id,omitempty" validate:"required"`
	RoleId int `json:"role_id,omitempty" validate:"required"`
}

type CreateRolePermissionPayload struct {
	RoleId       int `json:"role_id,omitempty" validate:"required"`
	PermissionId int `json:"permission_id,omitempty" validate:"required"`
}

type CreateRolePayload struct {
	Name string `json:"name,omitempty" validate:"required"`
}

type CreatePermissionPayload struct {
	Name string `json:"name,omitempty" validate:"required"`
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

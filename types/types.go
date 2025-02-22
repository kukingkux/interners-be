package types

import "time"

type PostStore interface {
	GetPosts() ([]*Post, error)
	GetPostById(id int) (*Post, error)
	GetPostsById(ids []int) ([]Post, error)
	UpdatePost(Post) error
	DeletePost(Post) error
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
	GetCompanies() ([]*Company, error)
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

type RolePermissionStore interface {
	GetRolePermissions() ([]*RolePermission, error)
	GetRolePermissionById(id int) (*UserRole, error)
	UpdateRolePermission(RolePermission) error
	CreateRolePermission(CreateRolePermissionPayload) error
}

type RolePermission struct {
	ID           int       `json:"id,omitempty"`
	RoleID       int       `json:"role_id,omitempty"`
	PermissionID int       `json:"permission_id,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}

type RoleStore interface {
	GetRoles() ([]*Role, error)
	GetRoleById(id int) (*UserRole, error)
	UpdateRole(Role) error
	CreateRole(CreateRolePayload) error
}

type Role struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type PermissionStore interface {
	GetPermissions() ([]*Permission, error)
	GetPermissionById(id int) (*Permission, error)
	UpdatePermission(Permission) error
	CreatePermission(CreatePermissionPayload) error
}

type Permission struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type UserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

type UserStore interface {
	GetUsers() ([]*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	UpdateUserAtFirstLogin(User) error 
	CreateUser(CreateUserPayload) error
}

type User struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
	PhoneNumber    string    `json:"phone_number,omitempty"`
	ZipCode        string    `json:"zip_code,omitempty"`
	City           string    `json:"city,omitempty"`
	Address        string    `json:"address,omitempty"`
	CV             string    `json:"cv,omitempty"`
	ProfilePicture string    `json:"profile_picture,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}

type CreatePostPayload struct {
	Title       string  `json:"title,omitempty" validate:"required"`
	Description string  `json:"description,omitempty" validate:"required"`
	Requirement string  `json:"requirement,omitempty" validate:"required"`
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
	UserID int `json:"user_id,omitempty" validate:"required"`
	RoleID int `json:"role_id,omitempty" validate:"required"`
}

type CreateRolePermissionPayload struct {
	RoleID       int `json:"role_id,omitempty" validate:"required"`
	PermissionID int `json:"permission_id,omitempty" validate:"required"`
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

type CreateUserPayload struct {
	Email    string `json:"email,omitempty" validate:"required,email"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

package user

import (
	"database/sql"
	"fmt"

	"github.com/kukingkux/interners-be/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUsers() ([]*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	users := make([]*types.User, 0)
	for rows.Next() {
		user, err := scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.PhoneNumber,
		&user.ZipCode,
		&user.City,
		&user.Address,
		&user.CV,
		&user.ProfilePicture,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) GetUserById(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}
func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec("INSERT INTO users (firstName, lastName, email, password) VALUES (?	,?,?,?)", user.FirstName, user.LastName, user.Email) //  user.Password

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateUserAtFirstLogin(user types.User) error {
	_, err := s.db.Exec("UPDATE users SET firstName = ?, lastName = ? WHERE email = ?", user.FirstName, user.LastName, user.Email)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

package userrole

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/kukingkux/interners-be/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserRoles() ([]*types.UserRole, error) {
	rows, err := s.db.Query("SELECT * FROM userroles")
	if err != nil {
		return nil, err
	}

	userroles := make([]*types.UserRole, 0)
	for rows.Next() {
		userrole, err := scanRowsIntoUserRole(rows)
		if err != nil {
			return nil, err
		}

		userroles = append(userroles, userrole)
	}
	return userroles, nil
}

func (s *Store) GetUserRoleById(userroleID int) (*types.UserRole, error) {
	rows, err := s.db.Query("SELECT * FROM userroles where id = ?", userroleID)
	if err != nil {
		return nil, err
	}

	userrole := new(types.UserRole)
	for rows.Next() {
		userrole, err = scanRowsIntoUserRole(rows)
		if err != nil {
			return nil, err
		}
	}
	return userrole, nil
}

func (s *Store) GetUserRolesById(userroleIDs []int) ([]types.UserRole, error) {
	placeholders := strings.Repeat(",?", len(userroleIDs)-1)
	query := fmt.Sprintf("SELECT * FROM userroles WHERE id IN (?%s)", placeholders)

	args := make([]interface{}, len(userroleIDs))
	for i, v := range userroleIDs {
		args[i] = v
	}
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	userroles := []types.UserRole{}
	for rows.Next() {
		userrole, err := scanRowsIntoUserRole(rows)
		if err != nil {
			return nil, err
		}

		userroles = append(userroles, *userrole)
	}
	return userroles, nil
}

func (s *Store) CreateUserRole(userrole types.CreateUserRolePayload) error {
	_, err := s.db.Exec("INSERT INTO userroles (userid, roleid) VALUES (?, ?)", userrole.UserID, userrole.RoleID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateUserRole(userrole types.UserRole) error {
	_, err := s.db.Exec("UPDATE userroles SET userid = ?, roleid = ? WHERE id = ?", userrole.UserID, userrole.RoleID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteUserRole(userrole types.UserRole) error {
	_, err := s.db.Exec("DELETE userroles WHERE id = ?", userrole.ID)
	if err != nil {
		return err
	}

	return nil
}

func scanRowsIntoUserRole(rows *sql.Rows) (*types.UserRole, error) {
	userrole := new(types.UserRole)

	err := rows.Scan(
		&userrole.ID,
		&userrole.UserID,
		&userrole.RoleID,
		&userrole.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return userrole, nil
}

package role

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

func (s *Store) GetRoles() ([]*types.Role, error) {
	rows, err := s.db.Query("SELECT * FROM roles")
	if err != nil {
		return nil, err
	}

	roles := make([]*types.Role, 0)
	for rows.Next() {
		role, err := scanRowsIntoRole(rows)
		if err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}
	return roles, nil
}

func (s *Store) GetRoleById(roleID int) (*types.Role, error) {
	rows, err := s.db.Query("SELECT * FROM roles where id = ?", roleID)
	if err != nil {
		return nil, err
	}

	role := new(types.Role)
	for rows.Next() {
		role, err = scanRowsIntoRole(rows)
		if err != nil {
			return nil, err
		}
	}
	return role, nil
}

func (s *Store) GetRolesById(roleIDs []int) ([]types.Role, error) {
	placeholders := strings.Repeat(",?", len(roleIDs)-1)
	query := fmt.Sprintf("SELECT * FROM roles WHERE id IN (?%s)", placeholders)

	args := make([]interface{}, len(roleIDs))
	for i, v := range roleIDs {
		args[i] = v
	}
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	roles := []types.Role{}
	for rows.Next() {
		role, err := scanRowsIntoRole(rows)
		if err != nil {
			return nil, err
		}

		roles = append(roles, *role)
	}
	return roles, nil
}

func (s *Store) CreateRole(role types.CreateRolePayload) error {
	_, err := s.db.Exec("INSERT INTO roles (name) VALUES (?)", role.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateRole(role types.Role) error {
	_, err := s.db.Exec("UPDATE roles SET name = ? WHERE id = ?", role.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteRole(role types.Role) error {
	_, err := s.db.Exec("DELETE roles WHERE id = ?", role.ID)
	if err != nil {
		return err
	}

	return nil
}

func scanRowsIntoRole(rows *sql.Rows) (*types.Role, error) {
	role := new(types.Role)

	err := rows.Scan(
		&role.ID,
		&role.Name,
		&role.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return role, nil
}

package permission

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

func (s *Store) GetPermissions() ([]*types.Permission, error) {
	rows, err := s.db.Query("SELECT * FROM permissions")
	if err != nil {
		return nil, err
	}

	permissions := make([]*types.Permission, 0)
	for rows.Next() {
		permission, err := scanRowsIntoPermission(rows)
		if err != nil {
			return nil, err
		}

		permissions = append(permissions, permission)
	}
	return permissions, nil
}

func (s *Store) GetPermissionById(permissionID int) (*types.Permission, error) {
	rows, err := s.db.Query("SELECT * FROM permissions where id = ?", permissionID)
	if err != nil {
		return nil, err
	}

	permission := new(types.Permission)
	for rows.Next() {
		permission, err = scanRowsIntoPermission(rows)
		if err != nil {
			return nil, err
		}
	}
	return permission, nil
}

func (s *Store) GetPermissionsById(permissionIDs []int) ([]types.Permission, error) {
	placeholders := strings.Repeat(",?", len(permissionIDs)-1)
	query := fmt.Sprintf("SELECT * FROM permissions WHERE id IN (?%s)", placeholders)

	args := make([]interface{}, len(permissionIDs))
	for i, v := range permissionIDs {
		args[i] = v
	}
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	permissions := []types.Permission{}
	for rows.Next() {
		permission, err := scanRowsIntoPermission(rows)
		if err != nil {
			return nil, err
		}

		permissions = append(permissions, *permission)
	}
	return permissions, nil
}

func (s *Store) CreatePermission(permission types.CreatePermissionPayload) error {
	_, err := s.db.Exec("INSERT INTO permissions (name) VALUES (?)", permission.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdatePermission(permission types.Permission) error {
	_, err := s.db.Exec("UPDATE permissions SET name = ? WHERE id = ?", permission.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeletePermission(permission types.Permission) error {
	_, err := s.db.Exec("DELETE permissions WHERE id = ?", permission.ID)
	if err != nil {
		return err
	}

	return nil
}

func scanRowsIntoPermission(rows *sql.Rows) (*types.Permission, error) {
	permission := new(types.Permission)

	err := rows.Scan(
		&permission.ID,
		&permission.Name,
		&permission.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return permission, nil
}

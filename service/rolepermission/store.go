package rolepermission

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

func (s *Store) GetRolePermissions() ([]*types.RolePermission, error) {
	rows, err := s.db.Query("SELECT * FROM rolepermissions")
	if err != nil {
		return nil, err
	}

	rolepermissions := make([]*types.RolePermission, 0)
	for rows.Next() {
		rolepermission, err := scanRowsIntoRolePermission(rows)
		if err != nil {
			return nil, err
		}

		rolepermissions = append(rolepermissions, rolepermission)
	}
	return rolepermissions, nil
}

func (s *Store) GetRolePermissionById(rolepermissionID int) (*types.RolePermission, error) {
	rows, err := s.db.Query("SELECT * FROM rolepermissions where id = ?", rolepermissionID)
	if err != nil {
		return nil, err
	}

	rolepermission := new(types.RolePermission)
	for rows.Next() {
		rolepermission, err = scanRowsIntoRolePermission(rows)
		if err != nil {
			return nil, err
		}
	}
	return rolepermission, nil
}

func (s *Store) GetRolePermissionsById(rolepermissionIDs []int) ([]types.RolePermission, error) {
	placeholders := strings.Repeat(",?", len(rolepermissionIDs)-1)
	query := fmt.Sprintf("SELECT * FROM rolepermissions WHERE id IN (?%s)", placeholders)

	args := make([]interface{}, len(rolepermissionIDs))
	for i, v := range rolepermissionIDs {
		args[i] = v
	}
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	rolepermissions := []types.RolePermission{}
	for rows.Next() {
		rolepermission, err := scanRowsIntoRolePermission(rows)
		if err != nil {
			return nil, err
		}

		rolepermissions = append(rolepermissions, *rolepermission)
	}
	return rolepermissions, nil
}

func (s *Store) CreateRolePermission(rolepermission types.CreateRolePermissionPayload) error {
	_, err := s.db.Exec("INSERT INTO rolepermissions (roleid, permissionid) VALUES (?, ?)", rolepermission.RoleID, rolepermission.PermissionID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateRolePermission(rolepermission types.RolePermission) error {
	_, err := s.db.Exec("UPDATE rolepermissions SET roleid = ?, permisisonid WHERE id = ?", rolepermission.RoleID, rolepermission.PermissionID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteRolePermission(rolepermission types.RolePermission) error {
	_, err := s.db.Exec("DELETE rolepermissions WHERE id = ?", rolepermission.ID)
	if err != nil {
		return err
	}

	return nil
}

func scanRowsIntoRolePermission(rows *sql.Rows) (*types.RolePermission, error) {
	rolepermission := new(types.RolePermission)

	err := rows.Scan(
		&rolepermission.ID,
		&rolepermission.RoleID,
		&rolepermission.PermissionID,
		&rolepermission.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return rolepermission, nil
}

package company

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

func (s *Store) GetCompanies() ([]*types.Company, error) {
	rows, err := s.db.Query("SELECT * FROM companies")
	if err != nil {
		return nil, err
	}

	companies := make([]*types.Company, 0)
	for rows.Next() {
		company, err := scanRowsIntoCompany(rows)
		if err != nil {
			return nil, err
		}

		companies = append(companies, company)
	}
	return companies, nil
}

func (s *Store) GetCompanyById(companyID int) (*types.Company, error) {
	rows, err := s.db.Query("SELECT * FROM companies where id = ?", companyID)
	if err != nil {
		return nil, err
	}

	company := new(types.Company)
	for rows.Next() {
		company, err = scanRowsIntoCompany(rows)
		if err != nil {
			return nil, err
		}
	}
	return company, nil
}

func (s *Store) GetCompaniesById(companyIDs []int) ([]types.Company, error) {
	placeholders := strings.Repeat(",?", len(companyIDs)-1)
	query := fmt.Sprintf("SELECT * FROM companies WHERE id IN (?%s)", placeholders)

	args := make([]interface{}, len(companyIDs))
	for i, v := range companyIDs {
		args[i] = v
	}
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	companies := []types.Company{}
	for rows.Next() {
		company, err := scanRowsIntoCompany(rows)
		if err != nil {
			return nil, err
		}

		companies = append(companies, *company)
	}
	return companies, nil
}

func (s *Store) CreateCompany(company types.CreateCompanyPayload) error {
	_, err := s.db.Exec("INSERT INTO companies (name, description, contact, email, address, province, city, logo, banner) VALUES (?,?,?,?,?,?,?,?,?)", company.Name, company.Description, company.Contact, company.Email, company.Address, company.Province, company.City, company.Logo, company.Banner)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateCompany(company types.Company) error {
	_, err := s.db.Exec("UPDATE companies SET name = ?, description = ?, contact = ?, email = ?, address = ?, province = ?, city = ?, logo = ?, banner = ? WHERE id = ?", company.Name, company.Description, company.Contact, company.Email, company.Address, company.Province, company.City, company.Logo, company.Banner, company.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteCompany(company types.Company) error {
	_, err := s.db.Exec("DELETE companies WHERE id = ?", company.ID)
	if err != nil {
		return err
	}

	return nil
}

func scanRowsIntoCompany(rows *sql.Rows) (*types.Company, error) {
	company := new(types.Company)

	err := rows.Scan(
		&company.ID,
		&company.UserID,
		&company.Name,
		&company.Description,
		&company.Contact,
		&company.Email,
		&company.Address,
		&company.Province,
		&company.City,
		&company.Logo,
		&company.Banner,
		&company.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return company, nil
}

package repository

import (
	"database/sql"
	"myapp/domain"
)

type SQLUserRepository struct {
	db *sql.DB
}

func NewSQLUserRepository(db *sql.DB) *SQLUserRepository {
	return &SQLUserRepository{db: db}
}

// Fulfilling the Contract: CreateUser (SQL Style)
func (s *SQLUserRepository) CreateUser(user *domain.User) error {
	query := `INSERT INTO users (id, name, email) VALUES ($1, $2, $3)`
	_, err := s.db.Exec(query, user.ID, user.Name, user.Email)
	return err
}

// Fulfilling the Contract: GetUser (SQL Style)
func (s *SQLUserRepository) GetUser(id string) (*domain.User, error) {
	var user domain.User
	query := `SELECT id, name, email FROM users WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
	return &user, err
}

// Placeholder for updates
func (s *SQLUserRepository) UpdateUser(user *domain.User) error { 
	return nil 
}

// Placeholder for deletes
func (s *SQLUserRepository) DeleteUser(id string) error { 
	return nil 
}

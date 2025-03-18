package repository

import (
	"user-management/models"

	"github.com/google/uuid"
)

func InsertUser(user models.User) error {
	query := `
		INSERT INTO users (id, created_at, updated_at, is_active, email, password, role) 
		VALUES ($1, NOW(), NOW(), $2, $3, $4, $5)
	`
	_, err := db.Exec(query, uuid.New(), true, user.Email, user.Password, user.Role)

	return err
}			
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

func GetAllUser(startRow, endRow *int) ([]models.User, error) {
	query := `SELECT id, created_at, updated_at, is_active, email, role FROM users`

	var args []any
	if startRow != nil && endRow != nil {
		query += ` LIMIT $1 OFFSET $2`
		args = append(args, *endRow-*startRow+1, *startRow)
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.IsActive, &user.Email, &user.Role)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}

func GetUserByID(id uuid.UUID) (*models.User, error) {
	query := `SELECT id, created_at, updated_at, is_active, email, role FROM users WHERE id = $1`
	row := db.QueryRow(query, id)
	
	var user models.User
	err := row.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.IsActive, &user.Email, &user.Role)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
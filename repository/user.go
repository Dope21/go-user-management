package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
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

func UpdateUserByID(user models.UpdateUser) (*models.User, error) {
	query := `UPDATE users SET `
	args := []any{}
	argTrack := 1
	setClauses := []string{}

	if user.IsActive != nil {
		setClauses = append(setClauses, fmt.Sprintf("username = $%d", argTrack))
		args = append(args, *user.IsActive)
		argTrack++
	}

	if user.Email != nil {
		setClauses = append(setClauses, fmt.Sprintf("email = $%d", argTrack))
		args = append(args, *user.Email)
		argTrack++
	}

	if user.Password != nil {
		setClauses = append(setClauses, fmt.Sprintf("password = $%d", argTrack))
		args = append(args, *user.Password)
		argTrack++
	}

	if user.Role != nil {
		setClauses = append(setClauses, fmt.Sprintf("role = $%d", argTrack))
		args = append(args, *user.Role)
		argTrack++
	}

	if len(setClauses) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query += strings.Join(setClauses, ", ")
	query += fmt.Sprintf(" WHERE id = $%d RETURNING *", argTrack)
	args = append(args, user.ID)

	row := db.QueryRow(query, args...)

	var updatedUser models.User
	err := row.Scan(
		&updatedUser.ID, 
		&updatedUser.CreatedAt, 
		&updatedUser.UpdatedAt,
		&updatedUser.IsActive, 
		&updatedUser.Email, 
		&updatedUser.Password,
		&updatedUser.Role,  
	)

	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func DeleteUserByID(id uuid.UUID) error {
	query := "DELETE FROM users WHERE id = $1"	 
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no user found with id: %s", id)
	}

	return nil
}

func GetUserByEmail(email string) (*models.User, error) {
	query := "SELECT * FROM users WHERE email = $1"
	row := db.QueryRow(query, email)

	var user models.User
	err := row.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.IsActive, &user.Email, &user.Password, &user.Role)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
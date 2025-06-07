package dto

import (
	"user-management/models"

	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Email			string			`json:"email" validate:"required,email"`
	Password 	string 			`json:"password" validate:"required,min=6,max=20"`
	Role 			models.Role `json:"role" validate:"required"`
}

type UpdateUserRequest struct {
	ID				uuid.UUID
	IsActive 	*bool     		`json:"is_active"`
	Email			*string				`json:"email" validate:"email"`
	Password 	*string 			`json:"password" validate:"min=6,max=20"`
	Role 			*models.Role 	`json:"role"`
}
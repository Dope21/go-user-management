package dto

import (
	"user-management/models"

	"github.com/google/uuid"
)

type UpdateUserRequest struct {
	ID       uuid.UUID `json:"id"`
	IsActive *bool     `json:"is_active"`
	Email    *string   `json:"email"`
	Password *string   `json:"password"`
	Role     *models.Role     `json:"role"`
}
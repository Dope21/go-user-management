package models

import (
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	RoleAdmin Role = "ADMIN"
	RoleUser Role = "USER"
	RoleMod	Role	= "MODERATOR"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	IsActive	bool			`json:"is_active"`
	Email    	string 	  `json:"email"`
	Password 	string 	  `json:"password"`
	Role     	Role 	 		`json:"role"`
}
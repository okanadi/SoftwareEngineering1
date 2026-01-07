package domain

import (
	"time"

	"github.com/google/uuid"
)

type UserRole string

const (
	RoleHandwerker UserRole = "handwerker"
	RoleBauleiter  UserRole = "innendienst"
	RoleAdmin      UserRole = "admin"
)

type CreateUserDTO struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Role     UserRole `json:"role"`
}

type UserDB struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Role      UserRole  `db:"role"`
	CreatedAt time.Time `db:"created_at"`
}

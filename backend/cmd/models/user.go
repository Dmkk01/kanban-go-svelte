package models

type Role string

const (
	AdminRole Role = "ADMIN"
	UserRole  Role = "USER"
)

type User struct {
	ID             int    `json:"id"`
	Email          string `json:"email"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	InactiveStatus bool   `json:"inactive_status"`
	Role           Role   `json:"role"`
}

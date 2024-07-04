package model

import "fmt"

// UserRole user role enumerate
type UserRole int32

const (
	RoleUser  UserRole = 0
	RoleAdmin UserRole = 1
)

func DefaultSignature(username string) string {
	return fmt.Sprintf("# Hi 👋, I'm %s\n", username)
}

type User struct {
	Id        int64    `gorm:"primarykey" json:"id"`
	Nickname  string   `json:"nickname"`
	Username  string   `gorm:"unique" json:"username"`
	Password  string   `json:"-"`
	Email     string   `gorm:"unique" json:"email"`
	Avatar    string   `json:"avatar"`
	Signature string   `json:"signature"`
	Role      UserRole `json:"role"`
}

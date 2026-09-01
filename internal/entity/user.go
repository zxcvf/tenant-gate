package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

var bcryptCost = bcrypt.DefaultCost

// User -.
type User struct {
	ID           int64  `json:"id,string"` // primary key
	Username     string `json:"username" example:"johndoe"`
	Email        string `json:"email" example:"user@example.com"`
	Phone        string `json:"phone" example:"+1234567890"`
	PasswordHash string `json:"-" example:"hashed_password"`

	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
} // @name entity.User

// CheckPassword checks if the provided password matches the user's password hash.
func (u *User) CheckPassword(password string) bool {
	if bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)) != nil {
		return false
	}
	return true
}

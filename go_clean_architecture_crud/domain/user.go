package domain

// The core User entity
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// The Strict Contract: Every database MUST have these 4 methods
type UserRepository interface {
	CreateUser(user *User) error
	GetUser(id string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id string) error
}

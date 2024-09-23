package types

import "time"

type UserPostRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// User represents data about a user.
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

// GroupUsers represents a set of users.
type GroupUsers struct {
	Users []*User
}

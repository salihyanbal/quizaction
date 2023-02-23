package user

import (
	"github.com/google/uuid"
)

type Role string

const (
	Admin  Role = "admin"
	Member Role = "member"
	Guest  Role = "guest"
)

type User struct {
	ID             uuid.UUID
	Username       string
	Email          string
	HashedPassword []byte
	Roles          []Role
}

type NewUserOpts struct {
	Username       string
	Email          string
	HashedPassword []byte
}

func NewUser(opts NewUserOpts) *User {
	return &User{
		ID:             uuid.New(),
		Username:       opts.Username,
		Email:          opts.Email,
		HashedPassword: opts.HashedPassword,
		Roles:          []Role{Member},
	}
}

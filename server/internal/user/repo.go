package user

import (
	"errors"
	"log"
)

type Repo interface {
	Create(u *User) error
	GetByUsername(username string) (*User, error)
}

type InMemRepo struct {
	db map[string]*User
}

func NewInMemRepo() Repo {
	return &InMemRepo{
		db: make(map[string]*User),
	}
}

func (r *InMemRepo) Create(u *User) error {
	log.Printf("User created %#v", u)
	r.db[u.Username] = u
	return nil
}

func (r *InMemRepo) GetByUsername(username string) (*User, error) {
	user, ok := r.db[username]
	log.Printf("%#v", r.db)
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// type PgRepo struct {

// }

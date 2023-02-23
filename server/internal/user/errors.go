package user

import "errors"

// TODO: Maybe also defining HTTPError for each error??

var (
	ErrUserNotFound = errors.New("user not found")
)

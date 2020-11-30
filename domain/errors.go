package domain

import "errors"

var (
	ErrNotFound      = errors.New("Your requested resource is not found")
	ErrUnauthorized  = errors.New("You don't have permission to access this resource")
	ErrBadParamInput = errors.New("Given param is not valid")
	ErrInternal      = errors.New("Internal server error")
)

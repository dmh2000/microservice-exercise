package repository

import "errors"

// ErrNoFound is return when a requested record is not found
var ErrNotFound = errors.New("record not found")
var ErrInvalidArgument = errors.New("invalid argument")

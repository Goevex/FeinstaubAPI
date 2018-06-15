package model

import (
	"errors"
)

// Errors
var (
	ErrNoResult = errors.New("no result")
	ErrExists   = errors.New("already exists")
	ErrNotExist = errors.New("does not exist")
)

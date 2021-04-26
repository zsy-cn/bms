package trashcan

import "errors"

var (
	ErrDbNotBeNil = errors.New("db not be nil")
	ErrIdNotBeNil = errors.New("id not be nil")
)

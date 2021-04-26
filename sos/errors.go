package sos

import "errors"

var (
	ErrCanNotUseSOS = errors.New("can not use sos")
	ErrDbNotBeNil   = errors.New("db not be nil")
	ErrIdNotBeNil   = errors.New("id not be nil")
)

package environ_monitor

import "errors"

var (
	ErrDbNotBeNil = errors.New("db not be nil")
	ErrIdNotBeNil = errors.New("sn not be nil")
)

package testutil

import "errors"

var (
	ErrDB                  = errors.New("error DB")
	ErrUnexpected          = errors.New("unexpected")
	ErrJsonUnsupportedType = errors.New("json: unsupported type: chan int")
)

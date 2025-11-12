package testutil

import "errors"

var (
	ErrDB         = errors.New("error DB")
	ErrUnexpected = errors.New("unexpected")
)

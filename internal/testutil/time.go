package testutil

import "time"

var (
	CreatedAt = time.Now().UTC()
	UpdatedAt = time.Now().UTC().Add(time.Hour)
)

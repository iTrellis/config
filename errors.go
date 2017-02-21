package config

import (
	"errors"
)

var (
	ErrNotMap   = errors.New("interface not map")
	ErrValueNil = errors.New("value is nil")
)

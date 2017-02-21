// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"errors"
)

var (
	ErrNotMap          = errors.New("interface not map")
	ErrValueNil        = errors.New("value is nil")
	ErrInvalidFilePath = errors.New("invalid file path")
	ErrUnknownSuffixes = errors.New("unknown file with suffix")
)

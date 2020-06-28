// GNU GPL v3 License
// Copyright (c) 2016 github.com:go-trellis

package files

import (
	"errors"
)

// errors
var (
	ErrFileIsAlreadyOpen = errors.New("file is already open")
	ErrFailedReadFile    = errors.New("failed read file")

	ErrReadBufferLengthBelowZero = errors.New("read buffer must above zero")
)

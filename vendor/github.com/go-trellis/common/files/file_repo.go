// GNU GPL v3 License
// Copyright (c) 2016 github.com:go-trellis

package files

import "os"

// FileRepo execute file functions
type FileRepo interface {
	// judge if file is opening
	FileOpened(string) bool
	// read file
	Read(string) (b []byte, n int, err error)
	// rewrite file with context
	Write(name, context string) (int, error)
	WriteBytes(name string, b []byte) (int, error)
	// append context to the file
	WriteAppend(name, context string) (int, error)
	WriteAppendBytes(name string, b []byte) (int, error)
	// rename file
	Rename(oldpath, newpath string) error
	// set length of buffer to read file, default: 1024
	SetReadBufLength(int) error
	// get information with file name
	FileInfo(name string) (os.FileInfo, error)
}

// FileStatus defile file status
type FileStatus int

// file status
const (
	// nothing
	FileStatusClosed FileStatus = iota
	// file is opened
	FileStatusOpening
	// file is moving or rename
	FileStatusMoving
)

// ReadBufferLength default reader buffer length
const (
	ReadBufferLength = 1024
)

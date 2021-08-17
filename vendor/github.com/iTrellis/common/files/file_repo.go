/*
Copyright © 2020 Henry Huang <hhh@rutcode.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package files

import (
	"os"
)

// FileRepo execute file functions
type FileRepo interface {
	Open(file string) (*FileInfo, error)
	OpenFile(file string, flag int, perm os.FileMode) (*FileInfo, error)
	// judge if file is opening
	FileOpened(string) bool
	// get information with file name
	FileInfo(name string) (os.FileInfo, error)
	// close the file
	Close(file string) error
	// close all files which were opened
	CloseAll() error
	// read file
	Read(string) (b []byte, n int, err error)
	// rewrite file with context
	Write(name, context string, opts ...WriteOption) (int, error)
	WriteBytes(name string, b []byte, opts ...WriteOption) (int, error)
	// append context to the file
	WriteAppend(name, context string, opts ...WriteOption) (int, error)
	WriteAppendBytes(name string, b []byte, opts ...WriteOption) (int, error)
	// rename file
	Rename(oldpath, newpath string) error
	// set length of buffer to read file, default: 1024
	SetReadBufLength(int64) error
}

type Option func(*Options)
type Options struct {
	ReadBufferLength int64
	Concurrency      bool
}

func ReadBufferLength(rbuf int64) Option {
	return func(o *Options) {
		o.ReadBufferLength = rbuf
	}
}

func Concurrency() Option {
	return func(o *Options) {
		o.Concurrency = true
	}
}

type WriteOption func(*WriteOptions)
type WriteOptions struct {
	Flag *int
}

func WriteFlag(flag int) WriteOption {
	return func(o *WriteOptions) {
		if o.Flag == nil {
			o.Flag = &flag
		} else {
			*o.Flag = *o.Flag | flag
		}
	}
}

// DefaultReadBufferLength default reader buffer length
const (
	DefaultReadBufferLength = 1024
)

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
	"io"
	"os"
	"sync"
	"sync/atomic"

	"github.com/iTrellis/common/errors"
)

// FileMode
const (
	FileModeOnlyRead  os.FileMode = 0444
	FileModeReadWrite os.FileMode = 0666
)

var (
	once        sync.Once
	defaultFile FileRepo
)

type fileGem struct {
	sync.RWMutex

	executingPath map[string]*FileInfo
	options       Options
}

type FileInfo struct {
	sync.Mutex
	*os.File
}

// New return filerepo with default executor
func New(opts ...Option) FileRepo {
	once.Do(func() {
		defaultFile = new(opts...)
	})
	return defaultFile
}

func NewFileInfo(opts ...Option) FileRepo {
	return new(opts...)
}

func new(opts ...Option) FileRepo {
	options := Options{}
	for _, o := range opts {
		o(&options)
	}
	if options.ReadBufferLength == 0 {
		options.ReadBufferLength = DefaultReadBufferLength
	}
	return &fileGem{
		executingPath: make(map[string]*FileInfo),
		options:       options,
	}
}

func (p *fileGem) Close(name string) error {
	p.Lock()
	defer p.Unlock()
	return p.close(name)
}

func (p *fileGem) close(name string) error {
	fi, ok := p.executingPath[name]
	if !ok {
		return nil
	}
	err := fi.Close()
	if err != nil {
		return err
	}

	delete(p.executingPath, name)
	return nil
}

func (p *fileGem) CloseAll() error {
	p.Lock()
	defer p.Unlock()
	var errs errors.Errors
	for k, fi := range p.executingPath {
		if err := fi.Close(); err != nil {
			errs = append(errs, err)
		}
		delete(p.executingPath, k)
	}
	return errs.Errors()
}

func (p *fileGem) Open(file string) (*FileInfo, error) {
	p.Lock()
	defer p.Unlock()
	return p.tryOpen(file)
}

func (p *fileGem) OpenFile(file string, flag int, perm os.FileMode) (*FileInfo, error) {
	p.Lock()
	defer p.Unlock()
	return p.tryOpenFile(file, flag, perm)
}

func (p *fileGem) Read(name string) (b []byte, n int, err error) {
	return p.read(name, int(p.options.ReadBufferLength))
}

func (p *fileGem) read(name string, bufLen int) (b []byte, n int, err error) {
	p.Lock()
	fi, e := p.tryOpen(name)
	if e != nil {
		p.Unlock()
		err = e
		return
	}
	p.Unlock()

	if !p.options.Concurrency {
		fi.Lock()
		defer fi.Unlock()
	}
	for {
		buf := make([]byte, bufLen)
		m, e := fi.Read(buf)
		if e != nil && e != io.EOF {
			err = ErrFailedReadFile
			return
		}
		n += m
		b = append(b, buf[:m]...)
		if m < bufLen {
			break
		}
	}

	return
}

func (p *fileGem) FileOpened(name string) bool {
	return p.executingPath[name] != nil
}

func (p *fileGem) tryOpen(name string) (*FileInfo, error) {
	return p.tryOpenFile(name, os.O_RDONLY, 0)
}

func (p *fileGem) tryOpenFile(name string, flag int, perm os.FileMode) (*FileInfo, error) {
	fi, ok := p.executingPath[name]
	if ok {
		p.executingPath[name] = fi
		return fi, nil
	}

	opened, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	fi = &FileInfo{
		File: opened,
	}
	p.executingPath[name] = fi

	return fi, nil
}

func (p *fileGem) Write(name, s string, opts ...WriteOption) (int, error) {
	return p.WriteBytes(name, []byte(s), opts...)
}

func (p *fileGem) WriteBytes(name string, b []byte, opts ...WriteOption) (int, error) {
	if len(opts) == 0 {
		opts = append(opts, WriteFlag(os.O_WRONLY))
	}
	return p.write(name, b, append(opts, WriteFlag(os.O_TRUNC|os.O_CREATE))...)
}

func (p *fileGem) WriteAppend(name, s string, opts ...WriteOption) (int, error) {
	return p.WriteAppendBytes(name, []byte(s), opts...)
}

func (p *fileGem) WriteAppendBytes(name string, b []byte, opts ...WriteOption) (int, error) {
	if len(opts) == 0 {
		opts = append(opts, WriteFlag(os.O_WRONLY))
	}
	return p.write(name, b, append(opts, WriteFlag(os.O_APPEND|os.O_CREATE))...)
}

func (p *fileGem) Rename(oldpath, newpath string) error {
	p.Lock()
	if err := p.close(oldpath); err != nil {
		p.Unlock()
		return err
	}
	p.Unlock()
	return os.Rename(oldpath, newpath)
}

func (p *fileGem) SetReadBufLength(l int64) error {
	if l <= 0 {
		return ErrReadBufferLengthBelowZero
	}

	atomic.StoreInt64(&p.options.ReadBufferLength, l)

	return nil
}

func (p *fileGem) write(name string, b []byte, opts ...WriteOption) (n int, err error) {

	options := &WriteOptions{}
	for _, o := range opts {
		o(options)
	}

	flag := 0
	if options.Flag != nil {
		flag = flag | *options.Flag
	}

	p.Lock()
	fi, err := p.tryOpenFile(name, flag, FileModeReadWrite)
	if err != nil {
		p.Unlock()
		return 0, err
	}
	p.Unlock()

	if !p.options.Concurrency {
		fi.Lock()
		defer fi.Unlock()
	}
	return fi.Write(b)
}

func (p *fileGem) FileInfo(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

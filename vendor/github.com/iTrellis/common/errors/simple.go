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

package errors

import (
	"fmt"

	"github.com/google/uuid"
)

// SimpleError simple error functions
type SimpleError interface {
	ID() string
	Namespace() string
	Message() string
	FullError() string
	Error() string
}

// Error error define
type Error struct {
	id        string
	namespace string
	message   string
}

// Newf 生成简单对象
func Newf(text string, params ...interface{}) SimpleError {
	return New(fmt.Sprintf(text, params...))
}

// New 生成简单对象
func New(text string) SimpleError {
	return new("T:S", uuid.New().String(), text)
}

func new(namespace, id, message string) *Error {
	return &Error{id: id, namespace: namespace, message: message}
}

func (p *Error) Error() string {
	return p.message
}

// FullError 全部错误信息
func (p *Error) FullError() string {
	return fmt.Sprintf("%s#%s:%s", p.namespace, p.id, p.message)
}

// ID 返回ID
func (p *Error) ID() string {
	return p.id
}

// Namespace 错误的域
func (p *Error) Namespace() string {
	return p.namespace
}

// Message 信息
func (p *Error) Message() string {
	return p.message
}

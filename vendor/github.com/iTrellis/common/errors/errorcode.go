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
	"encoding/json"
	"fmt"
	"strings"

	"github.com/iTrellis/common/data-structures/stack"
	"github.com/google/uuid"
)

// ErrorCode Error functions
type ErrorCode interface {
	SimpleError
	Code() uint64
	StackTrace() string
	Context() ErrorContext
	Append(err ...error) ErrorCode
	WithContext(k string, v interface{}) ErrorCode
}

// OptionFunc set params
type OptionFunc func(*ErrorOptions)

// ErrorOptions error parmas
type ErrorOptions struct {
	namespace string
	id        string
	code      uint64
	message   string

	stackTrace stack.Stack
	ctx        map[string]interface{}
}

// NewSimpleError new simple errors by options
func (p *ErrorOptions) NewSimpleError() SimpleError {
	return &Error{namespace: p.namespace, id: p.id, message: p.message}
}

// OptionID set id into options
func OptionID(id string) OptionFunc {
	return func(p *ErrorOptions) {
		p.id = id
	}
}

// OptionCode set error code into options
func OptionCode(code uint64) OptionFunc {
	return func(p *ErrorOptions) {
		p.code = code
	}
}

// OptionNamespace set error code into options
func OptionNamespace(ns string) OptionFunc {
	return func(p *ErrorOptions) {
		p.namespace = ns
	}
}

// OptionMesssage set error code into options
func OptionMesssage(msg string) OptionFunc {
	return func(p *ErrorOptions) {
		p.message = msg
	}
}

// OptionStackTrace set error code into options
func OptionStackTrace(stackTrace stack.Stack) OptionFunc {
	return func(p *ErrorOptions) {
		p.stackTrace = stackTrace
	}
}

// OptionContext set error code into options
func OptionContext(ctx map[string]interface{}) OptionFunc {
	return func(p *ErrorOptions) {
		p.ctx = ctx
	}
}

type errorCode struct {
	err        SimpleError
	code       uint64
	stackTrace stack.Stack
	context    map[string]interface{}
	errors     []string
}

// NewErrorCode get a new error code
func NewErrorCode(ofs ...OptionFunc) ErrorCode {
	opts := &ErrorOptions{}
	for _, o := range ofs {
		o(opts)
	}
	if opts.namespace == "" {
		opts.namespace = "T:E"
	}
	if opts.id == "" {
		opts.id = uuid.New().String()
	}

	if opts.ctx == nil {
		opts.ctx = make(map[string]interface{})
	}

	return &errorCode{
		err:        opts.NewSimpleError(),
		stackTrace: opts.stackTrace,
		context:    opts.ctx,
	}
}

func (p *errorCode) Append(err ...error) ErrorCode {
	if err == nil {
		return p
	}
	return p
}

func (p *errorCode) Code() uint64 {
	return p.code
}

func (p *errorCode) Message() string {
	return p.err.Message()
}

func (p *errorCode) Context() ErrorContext {
	return p.context
}

func (p *errorCode) Error() string {
	msg := p.err.Error()
	if len(p.errors) > 0 {
		msg = msg + "; " + strings.Join(p.errors, "; ")
	}
	return msg
}

func (p *errorCode) FullError() string {
	return strings.Join(
		append([]string{},
			fmt.Sprintf("ID:%s#%s", genErrorCodeKey(p.Namespace(), p.Code()), p.ID()),
			"Error:", p.Error(),
			"Context:", p.Context().Error(),
			"StackTrace:", p.StackTrace(),
		), "\n")
}

func (p *errorCode) ID() string {
	return p.err.ID()
}

func (p *errorCode) Namespace() string {
	return p.err.Namespace()
}

func (p *errorCode) StackTrace() string {
	frames, _ := p.stackTrace.PopAll()
	return frameToString(frames)
}

func (p *errorCode) WithContext(key string, value interface{}) ErrorCode {
	p.context[key] = value
	return p
}

// ErrorContext map contexts
type ErrorContext map[string]interface{}

func (p ErrorContext) Error() string {
	if p == nil {
		return ""
	}

	if bs, e := json.Marshal(p); e == nil {
		return string(bs)
	}
	return ""
}

func genErrorCodeKey(namespace string, code uint64) string {
	return fmt.Sprintf("%s:%d", namespace, code)
}

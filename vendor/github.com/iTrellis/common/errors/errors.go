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
	"strings"
)

// Errors error slice
type Errors []error

// NewErrors 生成错误数据对象
func NewErrors(errs ...error) Errors {
	var e Errors
	return e.Append(errs...)
}

func (p Errors) Error() string {
	return strings.Join(errorsString(p...), ";")
}

// Append 增补错误对象
func (p Errors) Append(errs ...error) Errors {
	if len(errs) == 0 {
		return p
	}
	p = append(p, errs...)
	return p
}

// Errors 判断是否为空
func (p Errors) Errors() error {
	if len(p) == 0 {
		return nil
	}
	return p
}

func errorsString(errs ...error) []string {
	var ss []string
	for _, e := range errs {
		switch ev := e.(type) {
		case ErrorCode:
			ss = append(ss, fmt.Sprintf("(%s#%d:%s):%s", ev.Namespace(), ev.Code(), ev.ID(), ev.Error()))
		case SimpleError:
			ss = append(ss, fmt.Sprintf("(%s:%s):%s", ev.Namespace(), ev.ID(), ev.Error()))
		default:
			ss = append(ss, ev.Error())
		}
	}
	return ss
}

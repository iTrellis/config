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

package testutils

import (
	"fmt"
	"reflect"
	"testing"
)

// Assert fails the test if the condition is false.
func Assert(tb testing.TB, condition bool, format string, a ...interface{}) {
	tb.Helper()
	if !condition {
		tb.Fatalf("\033[31m"+format+"\033[39m\n", a...)
	}
}

// Ok fails the test if an err is not nil.
func Ok(tb testing.TB, err error) {
	tb.Helper()
	if err != nil {
		tb.Fatalf("\033[31munexpected error: %v\033[39m\n", err)
	}
}

// NotOk fails the test if an err is nil.
func NotOk(tb testing.TB, err error, a ...interface{}) {
	tb.Helper()
	if err == nil {
		if len(a) != 0 {
			format := a[0].(string)
			tb.Fatalf("\033[31m"+format+": expected error, got none\033[39m", a...)
		}
		tb.Fatalf("\033[31mexpected error, got none\033[39m")
	}
}

// Equals fails the test if exp is not equal to act.
func Equals(tb testing.TB, exp, act interface{}, msgAndArgs ...interface{}) {
	tb.Helper()
	if !reflect.DeepEqual(exp, act) {
		tb.Fatalf("\033[31m%s\n\nexp: %#v\n\ngot: %#v\033[39m\n", formatMessage(msgAndArgs), exp, act)
	}
}

// ErrorEqual compares Go errors for equality.
func ErrorEqual(tb testing.TB, left, right error, msgAndArgs ...interface{}) {
	tb.Helper()
	if left == right {
		return
	}

	if left != nil && right != nil {
		Equals(tb, left.Error(), right.Error(), msgAndArgs...)
		return
	}

	tb.Fatalf("\033[31m%s\n\nexp: %#v\n\ngot: %#v\033[39m\n", formatMessage(msgAndArgs), left, right)
}

func formatMessage(msgAndArgs []interface{}) string {
	if len(msgAndArgs) == 0 {
		return ""
	}

	if msg, ok := msgAndArgs[0].(string); ok {
		return fmt.Sprintf("\n\nmsg: "+msg, msgAndArgs[1:]...)
	}
	return ""
}

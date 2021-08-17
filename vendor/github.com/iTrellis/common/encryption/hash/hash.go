/*
Copyright © 2016 Henry Huang <hhh@rutcode.com>

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

package hash

import "crypto"

// Repo hash functions manager
type Repo interface {
	Sum(s string) string
	SumBytes(b []byte) string
	SumTimes(s string, times uint) string
	SumBytesTimes(b []byte, times uint) string
}

// NewHashRepo get hash repo by crypto type
func NewHashRepo(h crypto.Hash) Repo {
	if h == crypto.MD5 {
		return NewMD5()
	} else if h == crypto.SHA1 {
		return NewSHA1()
	} else if h == crypto.SHA224 {
		return NewSHA224()
	} else if h == crypto.SHA256 {
		return NewSHA256()
	} else if h == crypto.SHA384 {
		return NewSHA384()
	} else if h == crypto.SHA512 {
		return NewSHA512()
	} else if h == crypto.SHA512_224 {
		return NewSHA512_224()
	} else if h == crypto.SHA512_256 {
		return NewSHA512_256()
	}

	return nil
}

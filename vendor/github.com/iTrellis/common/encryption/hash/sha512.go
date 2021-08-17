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

import "crypto/sha512"

// NewSHA384 get SHA384 hash repo
func NewSHA384() Repo {
	return &defaultHash{
		Hash: sha512.New384(),
	}
}

// NewSHA512 get SHA512 hash repo
func NewSHA512() Repo {
	return &defaultHash{
		Hash: sha512.New(),
	}
}

// NewSHA512_224 get SHA512_224 hash repo
func NewSHA512_224() Repo {
	return &defaultHash{
		Hash: sha512.New512_224(),
	}
}

// NewSHA512_256 get SHA512_256 hash repo
func NewSHA512_256() Repo {
	return &defaultHash{
		Hash: sha512.New512_256(),
	}
}

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

import "hash/crc32"

// NewCRC32 get crc32 hash32repo
func NewCRC32(tab *crc32.Table) Hash32Repo {
	return &defHash32{
		Hash: crc32.New(tab),
	}
}

// NewCRCIEEE get ieee hash32repo
func NewCRCIEEE() Hash32Repo {
	return &defHash32{
		Hash: crc32.NewIEEE(),
	}
}

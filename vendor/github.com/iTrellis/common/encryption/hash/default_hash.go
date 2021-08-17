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

import (
	"encoding/hex"
	"hash"
)

type defaultHash struct {
	Hash hash.Hash
}

func (p *defaultHash) Sum(s string) string {
	return p.SumBytes([]byte(s))
}

func (p *defaultHash) SumBytes(data []byte) string {
	p.Hash.Reset()
	_, err := p.Hash.Write(data)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(p.Hash.Sum(nil))
}

func (p *defaultHash) SumTimes(s string, times uint) string {
	if times == 0 {
		return ""
	}

	for i := 0; i < int(times); i++ {
		s = p.Sum(s)
	}
	return s
}

func (p *defaultHash) SumBytesTimes(b []byte, times uint) string {
	return p.SumTimes(string(b), times)
}

type defHash32 struct {
	Hash hash.Hash32
}

func (p *defHash32) Sum(s string) string {
	return p.SumBytes([]byte(s))
}

func (p *defHash32) SumBytes(data []byte) string {
	p.Hash.Reset()
	_, err := p.Hash.Write(data)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(p.Hash.Sum(nil))
}

func (p *defHash32) SumTimes(s string, times uint) string {
	if times == 0 {
		return ""
	}

	for i := 0; i < int(times); i++ {
		s = p.Sum(s)
	}
	return s
}

func (p *defHash32) SumBytesTimes(bs []byte, times uint) string {
	return p.SumTimes(string(bs), times)
}

func (p *defHash32) Sum32(b []byte) (uint32, error) {
	_, err := p.Hash.Write(b)
	if err != nil {
		return 0, err
	}
	return p.Hash.Sum32(), nil
}

/*
Copyright © 2021 Henry Huang <hhh@rutcode.com>

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

package formats

import (
	"fmt"
	"time"
)

type Duration time.Duration

// MarshalYAML implements the yaml.Marshaler interface.
func (d Duration) MarshalYAML() (interface{}, error) {
	return d.String(), nil
}

func (d Duration) String() string {
	var (
		ds   = int64(d)
		unit = "ms"
	)
	if ds == 0 {
		return "0s"
	}

	hour := int64(time.Hour)
	factors := map[string]int64{
		"y":  hour * 24 * 365,
		"w":  hour * 24 * 7,
		"d":  hour * 24,
		"h":  hour,
		"m":  int64(time.Minute),
		"s":  int64(time.Second),
		"ms": int64(time.Millisecond),
		"us": int64(time.Microsecond),
		"ns": int64(time.Nanosecond),
	}

	switch int64(0) {
	case ds % factors["y"]:
		unit = "y"
	case ds % factors["w"]:
		unit = "w"
	case ds % factors["d"]:
		unit = "d"
	case ds % factors["h"]:
		unit = "h"
	case ds % factors["m"]:
		unit = "m"
	case ds % factors["s"]:
		unit = "s"
	case ds % factors["ms"]:
		unit = "ms"
	case ds % factors["us"]:
		unit = "us"
	case ds % factors["ns"]:
		unit = "ns"
	}
	return fmt.Sprintf("%v%v", ds/factors[unit], unit)
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (d *Duration) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	dur := ParseStringTime(s)
	*d = Duration(dur)
	return nil
}

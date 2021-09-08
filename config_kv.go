/*
Copyright © 2017 Henry Huang <hhh@rutcode.com>

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

package config

import (
	"os"
	"reflect"
	"strings"

	"github.com/iTrellis/common/formats"
)

func (p *AdapterConfig) copyDollarSymbol(key string, maps *map[string]interface{}) error {
	tokens := []string{}
	if key != "" {
		tokens = append(tokens, key)
	}
	for k, v := range *maps {
		if v == nil {
			return nil
		}
		keys := append(tokens, k)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Map:
			{
				vm, ok := v.(map[string]interface{})
				if !ok {
					continue
				}
				err := p.copyDollarSymbol(strings.Join(keys, "."), &vm)
				if err != nil {
					return err
				}
			}
		case reflect.String:
			{
				s, ok := v.(string)
				if !ok {
					continue
				}
				_, matched := formats.FindStringSubmatchMap(s, includeReg)
				if !matched {
					continue
				}

				pkey := s[2 : len(s)-1]
				if p.EnvAllowed && (p.EnvPrefix == "" || strings.HasPrefix(pkey, p.EnvPrefix)) {
					if env := os.Getenv(pkey); env != "" {
						if err := p.setKeyValue(strings.Join(keys, "."), env); err != nil {
							return err
						}
						continue
					}
				}

				vm, err := p.getKeyValue(pkey)
				if err != nil {
					return err
				}
				err = p.setKeyValue(strings.Join(keys, "."), vm)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (p *AdapterConfig) getKeyValue(key string) (interface{}, error) {

	tokens := strings.Split(key, ".")
	vm := p.configs[tokens[0]]
	for i, t := range tokens {
		if i == 0 {
			continue
		}

		switch v := vm.(type) {
		case Options:
			vm = v[t]
		case map[string]interface{}:
			vm = v[t]
		case map[interface{}]interface{}:
			vm = v[t]
		default:
			return nil, ErrNotMap
		}
	}
	return vm, nil
}

// setKeyValue set key value into *configs
func (p *AdapterConfig) setKeyValue(key string, value interface{}) (err error) {
	tokens := strings.Split(key, ".")
	for i := len(tokens) - 1; i >= 0; i-- {
		if i == 0 {
			p.configs[tokens[0]] = value
			return
		}
		v, _ := p.getKeyValue(strings.Join(tokens[:i], "."))
		switch vm := v.(type) {
		case Options:
			vm[tokens[i]] = value
			value = vm
		case map[string]interface{}:
			vm[tokens[i]] = value
			value = vm
		case map[interface{}]interface{}:
			vm[tokens[i]] = value
			value = vm
		default:
			value = map[string]interface{}{tokens[i]: value}
		}
	}
	return
}

// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"reflect"
	"strings"
)

type ymlConfig struct{}

var yConfig = &ymlConfig{}

func NewYamlConfig(name string) (Config, error) {
	return newAdapterConfig(ReaderTypeYaml, name)
}

func (p *ymlConfig) copyDollarSymbol(configs *map[string]interface{}) {

	for k, v := range *configs {
		switch reflect.TypeOf(v).Kind() {
		case reflect.Map:
			vm, ok := v.(map[interface{}]interface{})
			if !ok {
				continue
			}
			p.copyMap(configs, k, &vm)
			// p.setKeyValue(configs, k, vm)
		case reflect.String:
			s, ok := v.(string)
			if !ok {
				continue
			}
			_, matched := findStringSubmatchMap(s, includeReg)
			if !matched {
				continue
			}
			vm, e := p.getKeyValue(*configs, s[2:len(s)-1])
			if e != nil {
				continue
			}
			p.setKeyValue(configs, k, vm)
		}
	}
	return
}

func (p *ymlConfig) copyMap(configs *map[string]interface{}, key string, maps *map[interface{}]interface{}) {
	tokens := []string{}
	if key != "" {
		tokens = append(tokens, key)
	}

	for k, v := range *maps {
		keys := append(tokens, k.(string))
		switch reflect.TypeOf(v).Kind() {
		case reflect.Map:
			{
				vm, ok := v.(map[interface{}]interface{})
				if !ok {
					continue
				}
				p.copyMap(configs, strings.Join(keys, "."), &vm)
				p.setKeyValue(configs, strings.Join(keys, "."), vm)
			}
		case reflect.String:
			{
				s, ok := v.(string)
				if !ok {
					continue
				}
				_, matched := findStringSubmatchMap(s, includeReg)
				if !matched {
					continue
				}
				vm, e := p.getKeyValue(*configs, s[2:len(s)-1])
				if e != nil {
					continue
				}
				p.setKeyValue(configs, strings.Join(keys, "."), vm)
			}
		}
	}
}

func (p *ymlConfig) getKeyValue(configs map[string]interface{}, key string) (vm interface{}, err error) {

	tokens := strings.Split(key, ".")
	vm = configs[tokens[0]]
	for i, t := range tokens {
		if i != 0 {
			v, ok := vm.(map[interface{}]interface{})
			if !ok {
				return nil, ErrNotMap
			}
			vm = v[t]
		}
	}

	if vm == nil {
		err = ErrValueNil
	}

	return
}

// setKeyValue set key value into *configs
func (p *ymlConfig) setKeyValue(configs *map[string]interface{}, key string, value interface{}) (err error) {
	tokens := strings.Split(key, ".")
	for i := len(tokens) - 1; i >= 0; i-- {
		if i == 0 {
			(*configs)[tokens[0]] = value
			return
		}
		v, _ := p.getKeyValue(*configs, strings.Join(tokens[:i], "."))
		vm, ok := v.(map[interface{}]interface{})
		if !ok {
			value = map[interface{}]interface{}{tokens[i]: value}
			continue
		}
		vm[tokens[i]] = value
		value = vm
	}
	return
}

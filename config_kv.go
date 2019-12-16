// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"reflect"
	"strings"

	"github.com/go-trellis/formats"
)

func copyJSONDollarSymbol(configs *map[string]interface{}, key string, maps *map[string]interface{}) {
	tokens := []string{}
	if key != "" {
		tokens = append(tokens, key)
	}
	for k, v := range *maps {
		keys := append(tokens, k)
		switch reflect.TypeOf(v).Kind() {
		case reflect.Map:
			{
				vm, ok := v.(map[string]interface{})
				if !ok {
					continue
				}
				copyJSONDollarSymbol(configs, strings.Join(keys, "."), &vm)
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
				vm, e := getStringKeyValue(*configs, s[2:len(s)-1])
				if e != nil {
					continue
				}
				setStringKeyValue(configs, strings.Join(keys, "."), vm)
			}
		}
	}
	return
}

func copyYAMLDollarSymbol(configs *map[string]interface{}) {

	for k, v := range *configs {
		switch reflect.TypeOf(v).Kind() {
		case reflect.Map:
			{
				vm, ok := v.(map[interface{}]interface{})
				if !ok {
					continue
				}
				copyMap(configs, k, &vm)
			}
		case reflect.String:
			{
				s, ok := v.(string)
				if !ok {
					continue
				}
				if _, matched := formats.FindStringSubmatchMap(s, includeReg); !matched {
					continue
				}
				vm, e := getInterfaceKeyValue(*configs, s[2:len(s)-1])
				if e != nil {
					continue
				}
				setInterfaceKeyValue(configs, k, vm)
			}
		}
	}
	return
}

func copyMap(configs *map[string]interface{}, key string, maps *map[interface{}]interface{}) {
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
				copyMap(configs, strings.Join(keys, "."), &vm)
				setInterfaceKeyValue(configs, strings.Join(keys, "."), vm)
			}
		case reflect.String:
			{
				s, ok := v.(string)
				if !ok {
					continue
				}
				if _, matched := formats.FindStringSubmatchMap(s, includeReg); !matched {
					continue
				}
				vm, e := getInterfaceKeyValue(*configs, s[2:len(s)-1])
				if e != nil {
					continue
				}
				setInterfaceKeyValue(configs, strings.Join(keys, "."), vm)
			}
		}
	}
}

func getInterfaceKeyValue(configs map[string]interface{}, key string) (vm interface{}, err error) {

	tokens := strings.Split(key, ".")
	vm = configs[tokens[0]]
	for i, t := range tokens {
		if i == 0 {
			continue
		}
		v, ok := vm.(map[interface{}]interface{})
		if !ok {
			return nil, ErrNotMap
		}
		vm = v[t]
	}
	if vm == nil {
		err = ErrValueNil
	}
	return
}

// setInterfaceKeyValue set key value into *configs
func setInterfaceKeyValue(configs *map[string]interface{}, key string, value interface{}) (err error) {
	tokens := strings.Split(key, ".")
	for i := len(tokens) - 1; i >= 0; i-- {
		if i == 0 {
			(*configs)[tokens[0]] = value
			return
		}
		v, _ := getInterfaceKeyValue(*configs, strings.Join(tokens[:i], "."))
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

func getStringKeyValue(configs map[string]interface{}, key string) (vm interface{}, err error) {

	tokens := strings.Split(key, ".")
	vm = configs[tokens[0]]
	for i, t := range tokens {
		if i != 0 {
			v, ok := vm.(map[string]interface{})
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

// setStringKeyValue set key value into configs
func setStringKeyValue(configs *map[string]interface{}, key string, value interface{}) (err error) {
	tokens := strings.Split(key, ".")
	for i := len(tokens) - 1; i >= 0; i-- {
		if i == 0 {
			(*configs)[tokens[i]] = value
			return
		}
		v, _ := getStringKeyValue(*configs, strings.Join(tokens[:i], "."))
		vm, ok := v.(map[string]interface{})
		if !ok {
			value = map[string]interface{}{tokens[i]: value}
			continue
		}
		vm[tokens[i]] = value
		value = vm
	}
	return
}

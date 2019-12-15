// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"strings"
)

type jsonConfig struct{}

var jConfig = &jsonConfig{}

// NewJSONConfig get json config reader
func NewJSONConfig(name string) (Config, error) {
	return newAdapterConfig(ReaderTypeJSON, name)
}

func (*jsonConfig) getKeyValue(configs map[string]interface{}, key string) (vm interface{}, err error) {

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

// setKeyValue set key value into configs
func (p *jsonConfig) setKeyValue(configs *map[string]interface{}, key string, value interface{}) (err error) {
	tokens := strings.Split(key, ".")
	for i := len(tokens) - 1; i >= 0; i-- {
		if i == 0 {
			(*configs)[tokens[i]] = value
			return
		}
		v, _ := p.getKeyValue(*configs, strings.Join(tokens[:i], "."))
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

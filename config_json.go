// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"reflect"
	"strings"
	"sync"

	"github.com/go-trellis/formats/inner-types"
)

type jsonConfig struct {
	name    string
	locker  sync.RWMutex
	reader  Reader
	configs map[string]interface{}
}

func NewJsonConfig(name string) (Config, error) {
	if name == "" {
		return nil, ErrInvalidFilePath
	}
	c := &jsonConfig{
		reader:  NewJsonReader(),
		configs: make(map[string]interface{}),
		name:    name,
	}

	e := c.reader.Read(c.name, &c.configs)

	return c, e
}

// GetInterface return a interface object in p.configs by key
func (p *jsonConfig) GetInterface(key string, defValue ...interface{}) (res interface{}) {

	if key == "" {
		return nil
	}

	var err error
	var v interface{}

	defer func() {
		if err != nil {
			if len(defValue) == 0 {
				return
			}
			res = defValue[0]
		} else {
			res = v
		}
	}()

	p.locker.RLock()
	defer p.locker.RUnlock()
	v, err = p.getKeyValue(key)
	return
}

// GetString return a string object in p.configs by key
func (p *jsonConfig) GetString(key string, defValue ...string) (res string) {

	var ok bool
	defer func() {
		if !ok {
			if len(defValue) == 0 {
				return
			}
			res = defValue[0]
		}
	}()
	v := p.GetInterface(key, defValue)

	res, ok = v.(string)
	return
}

// GetBoolean return a bool object in p.configs by key
func (p *jsonConfig) GetBoolean(key string, defValue ...bool) (b bool) {

	var ok bool
	defer func() {
		if !ok {
			if len(defValue) == 0 {
				return
			}
			b = defValue[0]
		}
	}()
	v := p.GetInterface(key, defValue)

	switch reflect.TypeOf(v).Kind() {
	case reflect.Bool:
		b = v.(bool)
	case reflect.String:
		s := v.(string)
		return s == "ON" || s == "on"
	}

	return
}

// GetInt return a int object in p.configs by key
func (p *jsonConfig) GetInt(key string, defValue ...int) (res int) {

	var err error
	defer func() {
		if err != nil {
			if len(defValue) == 0 {
				return
			}
			res = defValue[0]
		}
	}()

	v, e := itypes.ToInt64(p.GetInterface(key, defValue))
	if e != nil {
		err = e
		return
	}
	return int(v)
}

// GetFloat return a float object in p.configs by key
func (p *jsonConfig) GetFloat(key string, defValue ...float64) (res float64) {

	var err error
	defer func() {
		if err != nil {
			if len(defValue) == 0 {
				return
			}
			res = defValue[0]
		}
	}()

	v, e := itypes.ToFloat64(p.GetInterface(key, defValue))
	if e != nil {
		err = e
		return
	}
	return v
}

// GetList return a list object in p.configs by key
func (p *jsonConfig) GetList(key string, defValue ...interface{}) (res interface{}) {

	v := p.GetInterface(key, defValue)
	if reflect.TypeOf(v).Kind() != reflect.Slice {
		return nil
	}
	return v
}

// GetConfig return object config in p.configs by key
func (p *jsonConfig) GetConfig(key string) Config {
	p.locker.RLock()
	defer p.locker.RUnlock()
	vm, err := p.getKeyValue(key)
	if err != nil {
		return nil
	}

	c := &jsonConfig{
		reader:  p.reader,
		configs: map[string]interface{}{key: vm},
	}

	return c
}

func (p *jsonConfig) getKeyValue(key string) (vm interface{}, err error) {

	defer func() {
		if vm == nil {
			err = ErrValueNil
			return
		}
	}()

	tokens := strings.Split(key, ".")
	lenTokens := len(tokens) - 1

	vm = p.configs[tokens[0]]
	for i, t := range tokens {
		if i != 0 {
			v, ok := vm.(map[string]interface{})
			if !ok {
				return nil, ErrNotMap
			}
			vm = v[t]
		}

		if lenTokens == i {
			return
		}
	}

	return
}

// SetKeyValue set key value into p.configs
func (p *jsonConfig) SetKeyValue(key string, value interface{}) (err error) {
	if key == "" {
		return
	}
	p.locker.Lock()
	defer p.locker.Unlock()

	tokens := strings.Split(key, ".")

	for i := len(tokens) - 1; i >= 0; i-- {
		if i == 0 {
			p.configs[tokens[i]] = value
		} else {
			v, _ := p.getKeyValue(strings.Join(tokens[:i], "."))
			if vm, ok := v.(map[string]interface{}); !ok {
				value = map[string]interface{}{tokens[i]: value}
			} else {
				vm[tokens[i]] = value
				value = vm
			}
		}
	}

	return
}

// Dump return p.configs' bytes
func (p *jsonConfig) Dump() (bs []byte, err error) {
	return p.reader.Dump(p.configs)
}

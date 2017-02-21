package config

import (
	"reflect"
	"strings"
)

type Config struct {
	reader  Reader
	configs map[string]interface{}
}

func NewConfig(name string) (*Config, error) {
	c := &Config{
		reader:  NewSuffixReader(),
		configs: make(map[string]interface{})}

	e := c.reader.Read(name, &c.configs)

	return c, e
}

func (p *Config) GetString(key string, defValue ...string) (str string) {
	v, e := p.getKey(key)
	if e != nil {
		if len(defValue) == 0 {
			return
		} else {
			return defValue[0]
		}
	}

	var ok bool
	str, ok = v.(string)
	if !ok {
		if len(defValue) == 0 {
			return
		} else {
			return defValue[0]
		}
	}
	return
}

func (p *Config) GetBoolean(key string, defValue ...bool) (b bool) {

	var (
		ok  bool
		err error
	)

	defer func() {
		if err != nil || !ok {
			if len(defValue) == 0 {
				return
			}
			b = defValue[0]
		}
	}()

	var v interface{}
	v, err = p.getKey(key)
	if err != nil {
		return
	}

	if v == nil {
		return
	}

	switch reflect.TypeOf(v).Kind() {
	case reflect.Bool:
		b = v.(bool)
	case reflect.String:
		s := v.(string)
		return s == "ON" || s == "on"
	}

	return
}

func (p *Config) GetInt(key string, defValue ...int) (res int) {

	var (
		ok  bool
		err error
	)

	defer func() {
		if err != nil || !ok {
			if len(defValue) == 0 {
				return
			}
			res = defValue[0]
		}
	}()

	var v interface{}
	v, err = p.getKey(key)
	if err != nil {
		return
	}

	res, ok = v.(int)
	return
}

func (p *Config) GetFloat(key string, defValue ...float64) (res float64) {

	var (
		ok  bool
		err error
	)

	defer func() {
		if err != nil || !ok {
			if len(defValue) == 0 {
				return
			}
			res = defValue[0]
		}
	}()

	var v interface{}
	v, err = p.getKey(key)
	if err != nil {
		return
	}

	res, ok = v.(float64)
	return
}

func (p *Config) GetConfig(key string) *Config {
	vm, err := p.getKey(key)
	if err != nil {
		return nil
	}

	c := &Config{
		reader:  p.reader,
		configs: map[string]interface{}{key: vm},
	}

	return c
}

func (p *Config) getKey(key string) (vm interface{}, err error) {
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
			v, ok := vm.(map[interface{}]interface{})
			if !ok {
				err = ErrNotMap
				return
			}
			vm = v[t]
		}

		if lenTokens == i {
			return
		}
	}

	return
}

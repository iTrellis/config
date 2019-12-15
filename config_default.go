// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"math/big"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/go-trellis/formats"
)

const (
	includeReg = `\$\{([0-9|a-z|A-Z]|\.)+\}`
)

// default config adapter
type adapterConfig struct {
	name       string
	readerType ReaderType
	reader     Reader
	locker     sync.RWMutex
	configs    map[string]interface{}
}

// newAdapterConfig return default config adapter
// name is file's path
func newAdapterConfig(rt ReaderType, name string) (Config, error) {
	if name == "" {
		return nil, ErrInvalidFilePath
	}
	a := &adapterConfig{
		readerType: rt,
		name:       name,
		configs:    make(map[string]interface{}),
	}

	switch rt {
	case ReaderTypeJSON:
		a.reader = NewJSONReader()
	case ReaderTypeYAML:
		a.reader = NewYAMLReader()
	default:
		return nil, ErrNotSupportedReaderType
	}
	if e := a.reader.Read(a.name, &a.configs); e != nil {
		return nil, e
	}

	return a.copy(), nil
}

func (p *adapterConfig) GetKeys() []string {
	p.locker.RLock()
	defer p.locker.RUnlock()

	var keys []string
	for key := range p.configs {
		keys = append(keys, key)
	}
	return keys
}

func (p *adapterConfig) copy() Config {
	p.locker.RLock()
	defer p.locker.RUnlock()

	values := DeepCopy(p.configs)

	valuesMap := values.(map[string]interface{})
	return &adapterConfig{
		name:       p.name,
		readerType: p.readerType,
		reader:     p.reader,
		configs:    valuesMap,
	}
}

// GetTimeDuration return time in p.configs by key
func (p *adapterConfig) GetTimeDuration(key string, defValue ...time.Duration) time.Duration {
	return formats.ParseStringTime(strings.ToLower(p.GetString(key)))
}

// GetByteSize return time in p.configs by key
func (p *adapterConfig) GetByteSize(key string) *big.Int {
	return formats.ParseStringByteSize(strings.ToLower(p.GetString(key)))
}

// GetInterface return a interface object in p.configs by key
func (p *adapterConfig) GetInterface(key string, defValue ...interface{}) (res interface{}) {

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

	if key == "" {
		return ErrInvalidKey
	}

	v, err = p.getKeyValue(key)
	return
}

// GetString return a string object in p.configs by key
func (p *adapterConfig) GetString(key string, defValue ...string) (res string) {

	var ok bool
	defer func() {
		if ok || len(defValue) == 0 {
			return
		}
		res = defValue[0]
	}()
	v := p.GetInterface(key, defValue)

	res, ok = v.(string)
	return
}

// GetBoolean return a bool object in p.configs by key
func (p *adapterConfig) GetBoolean(key string, defValue ...bool) (b bool) {

	var ok bool
	defer func() {
		if ok || len(defValue) == 0 {
			return
		}
		b = defValue[0]
	}()
	v := p.GetInterface(key, defValue)

	switch reflect.TypeOf(v).String() {
	case "bool":
		ok, b = true, v.(bool)
	case "string":
		ok, b = true, (v.(string) == "ON" || v.(string) == "on")
	}

	return
}

// GetInt return a int object in p.configs by key
func (p *adapterConfig) GetInt(key string, defValue ...int) (res int) {

	var err error
	defer func() {
		if err != nil {
			if len(defValue) == 0 {
				return
			}
			res = defValue[0]
		}
	}()

	v, e := formats.ToInt64(p.GetInterface(key, defValue))
	if e != nil {
		err = e
		return
	}
	return int(v)
}

// GetFloat return a float object in p.configs by key
func (p *adapterConfig) GetFloat(key string, defValue ...float64) (res float64) {

	var err error
	defer func() {
		if err != nil {
			if len(defValue) == 0 {
				return
			}
			res = defValue[0]
		}
	}()

	v, e := formats.ToFloat64(p.GetInterface(key, defValue))
	if e != nil {
		err = e
		return
	}
	return v
}

// GetList return a list of interface{} in p.configs by key
func (p *adapterConfig) GetList(key string) (res []interface{}) {

	vS := reflect.Indirect(reflect.ValueOf(p.GetInterface(key)))
	if vS.Kind() != reflect.Slice {
		return nil
	}

	var vs []interface{}
	for i := 0; i < vS.Len(); i++ {
		vs = append(vs, vS.Index(i).Interface())
	}
	return vs
}

// GetStringList return a list of strings in p.configs by key
func (p *adapterConfig) GetStringList(key string) []string {

	var items []string
	for _, v := range p.GetList(key) {
		item, ok := v.(string)
		if !ok {
			return nil
		}

		items = append(items, item)
	}
	return items
}

// GetBooleanList return a list of booleans in p.configs by key
func (p *adapterConfig) GetBooleanList(key string) []bool {

	var items []bool
	for _, v := range p.GetList(key) {
		item, ok := v.(bool)
		if !ok {
			return nil
		}

		items = append(items, item)
	}
	return items
}

// GetIntList return a list of ints in p.configs by key
func (p *adapterConfig) GetIntList(key string) []int {

	var items []int
	for _, v := range p.GetList(key) {
		i, e := formats.ToInt(v)
		if e != nil {
			return nil
		}
		items = append(items, i)
	}
	return items
}

// GetFloatList return a list of floats in p.configs by key
func (p *adapterConfig) GetFloatList(key string) []float64 {

	var items []float64
	for _, v := range p.GetList(key) {
		f, e := formats.ToFloat64(v)
		if e != nil {
			return nil
		}
		items = append(items, f)
	}
	return items
}

// get map value
func (p *adapterConfig) GetMap(key string) Options {

	vm, err := p.getKeyValue(key)
	if err != nil {
		return nil
	}

	mapVM, ok := vm.(map[string]interface{})
	if ok {
		return mapVM
	}

	mapVMs, ok := vm.(map[interface{}]interface{})
	if !ok {
		return nil
	}

	result := make(map[string]interface{})
	for k, v := range mapVMs {
		sk, _ := k.(string)
		result[sk] = v
	}
	return result
}

// GetConfig return object config in p.configs by key
func (p *adapterConfig) GetConfig(key string) Config {

	vm, err := p.getKeyValue(key)
	if err != nil {
		return nil
	}

	c := &adapterConfig{
		reader:  p.reader,
		configs: map[string]interface{}{key: vm},
	}

	return c
}

// get key's values if values can be Config, or panic
func (p *adapterConfig) GetValuesConfig(key string) Config {
	opt := p.GetMap(key)
	if opt == nil {
		return nil
	}

	return MapGetter().GenMapConfig(opt)
}

func (p *adapterConfig) getKeyValue(key string) (vm interface{}, err error) {
	if key == "" {
		return nil, ErrInvalidKey
	}
	p.locker.RLock()
	defer p.locker.RUnlock()

	switch p.readerType {
	case ReaderTypeJSON:
		return jConfig.getKeyValue(p.configs, key)
	case ReaderTypeYAML:
		return yConfig.getKeyValue(p.configs, key)
	default:
		return nil, ErrNotSupportedReaderType
	}
}

// SetKeyValue set key value into p.configs
func (p *adapterConfig) SetKeyValue(key string, value interface{}) (err error) {
	if key == "" {
		return ErrInvalidKey
	}
	p.locker.Lock()
	defer p.locker.Unlock()

	switch p.readerType {
	case ReaderTypeJSON:
		return jConfig.setKeyValue(&p.configs, key, value)
	case ReaderTypeYAML:
		return yConfig.setKeyValue(&p.configs, key, value)
	default:
		return ErrNotSupportedReaderType
	}
}

// Dump return p.configs' bytes
func (p *adapterConfig) Dump() (bs []byte, err error) {
	p.locker.Lock()
	defer p.locker.Unlock()

	return p.reader.Dump(p.configs)
}

// Dump return p.configs' bytes
func (p *adapterConfig) Copy() Config {
	return p.copy()
}

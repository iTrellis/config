package config

import (
	"reflect"
	"sync"
	"time"

	"github.com/go-trellis/formats/inner-types"
)

const (
	timeReg = `^(?P<value>([0-9]+(\.[0-9]+)?))\s*(?P<unit>(nanoseconds|nanosecond|nanos|nano|ns|microseconds|microsecond|micros|micro|us|milliseconds|millisecond|millis|milli|ms|seconds|second|s|minutes|minute|m|hours|hour|h|days|day|d))$`
)

type adapterConfig struct {
	name       string
	readerType ReaderType
	reader     Reader
	locker     sync.RWMutex
	configs    map[string]interface{}
}

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
	case ReaderTypeJson:
		a.reader = NewJsonReader()
	case ReaderTypeYaml:
		a.reader = NewYamlReader()
	default:
		return nil, ErrNotSupportedReaderType
	}
	e := a.reader.Read(a.name, &a.configs)

	return a, e
}

// GetTimeDuration return time in p.configs by key
func (p *adapterConfig) GetTimeDuration(key string, defValue ...time.Duration) time.Duration {
	groups, matched := findStringSubmatchMap(p.GetString(key), timeReg)

	if matched {
		i, _ := itypes.ToInt64(groups["value"])

		switch groups["unit"] {
		case "nanoseconds", "nanosecond", "nanos", "nano", "ns":
			return time.Nanosecond * time.Duration(i)
		case "microseconds", "microsecond", "micros", "micro", "us":
			return time.Microsecond * time.Duration(i)
		case "milliseconds", "millisecond", "millis", "milli", "ms":
			return time.Millisecond * time.Duration(i)
		case "seconds", "second", "s":
			return time.Second * time.Duration(i)
		case "minutes", "minute", "m":
			return time.Minute * time.Duration(i)
		case "hours", "hour", "h":
			return time.Hour * time.Duration(i)
		case "days", "day", "d":
			return time.Hour * 24 * time.Duration(i)
		}
	}

	if len(defValue) == 0 {
		return 0
	}

	return defValue[0]
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
func (p *adapterConfig) GetBoolean(key string, defValue ...bool) (b bool) {

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

	v, e := itypes.ToInt64(p.GetInterface(key, defValue))
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

	v, e := itypes.ToFloat64(p.GetInterface(key, defValue))
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
		item, ok := v.(int)
		if !ok {
			return nil
		}

		items = append(items, item)
	}
	return items
}

// GetFloatList return a list of floats in p.configs by key
func (p *adapterConfig) GetFloatList(key string) []float64 {

	var items []float64
	for _, v := range p.GetList(key) {
		item, ok := v.(float64)
		if !ok {
			return nil
		}

		items = append(items, item)
	}
	return items
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

func (p *adapterConfig) getKeyValue(key string) (vm interface{}, err error) {
	if key == "" {
		return nil, ErrInvalidKey
	}
	p.locker.RLock()
	defer p.locker.RUnlock()

	switch p.readerType {
	case ReaderTypeJson:
		return jConfig.getKeyValue(p.configs, key)
	case ReaderTypeYaml:
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
	case ReaderTypeJson:
		return jConfig.setKeyValue(&p.configs, key, value)
	case ReaderTypeYaml:
		return yConfig.setKeyValue(&p.configs, key, value)
	default:
		return ErrNotSupportedReaderType
	}
}

// Dump return p.configs' bytes
func (p *adapterConfig) Dump() (bs []byte, err error) {
	return p.reader.Dump(p.configs)
}

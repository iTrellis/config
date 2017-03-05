package config

import (
	"math/big"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/go-trellis/formats/inner-types"
)

const (
	timeReg = `^(?P<value>([0-9]+(\.[0-9]+)?))\s*(?P<unit>(nanoseconds|nanosecond|nanos|nano|ns|microseconds|microsecond|micros|micro|us|milliseconds|millisecond|millis|milli|ms|seconds|second|s|minutes|minute|m|hours|hour|h|days|day|d))$`

	bitReg = `^(?P<value>([0-9]+(\.[0-9]+)?))\s*(?P<unit>(b|byte|bytes|kb|kilobyte|kilobytes|mb|megabyte|megabytes|gb|gigabyte|gigabytes|tb|terabyte|terabytes|pb|petabyte|petabytes|eb|exabyte|exabytes|zb|zettabyte|zettabytes|yb|yottabyte|yottabytes|k|ki|kib|kibibyte|kibibytes|m|mi|mib|mebibyte|mebibytes|g|gi|gib|gibibyte|gibibytes|t|ti|tib|tebibyte|tebibytes|p|pi|pib|pebibyte|pebibytes|e|ei|eib|exbibyte|exbibytes|z|zi|zib|zebibyte|zebibytes|y|yi|yib|yobibyte|yobibytes))$`

	includeReg = `\$\{([0-9|a-z|A-Z]|\.)+\}`
)

// ByteSizes
var (
	_Num1000 = big.NewInt(1000)
	_Num1024 = big.NewInt(1024)

	_Byte   = big.NewInt(1)
	_KiByte = (&big.Int{}).Mul(_Byte, _Num1024)
	_MiByte = (&big.Int{}).Mul(_KiByte, _Num1024)
	_GiByte = (&big.Int{}).Mul(_MiByte, _Num1024)
	_TiByte = (&big.Int{}).Mul(_GiByte, _Num1024)
	_PiByte = (&big.Int{}).Mul(_TiByte, _Num1024)
	_EiByte = (&big.Int{}).Mul(_PiByte, _Num1024)
	_ZiByte = (&big.Int{}).Mul(_EiByte, _Num1024)
	_YiByte = (&big.Int{}).Mul(_ZiByte, _Num1024)

	_KByte = (&big.Int{}).Mul(_Byte, _Num1000)
	_MByte = (&big.Int{}).Mul(_KByte, _Num1000)
	_GByte = (&big.Int{}).Mul(_MByte, _Num1000)
	_TByte = (&big.Int{}).Mul(_GByte, _Num1000)
	_PByte = (&big.Int{}).Mul(_TByte, _Num1000)
	_EByte = (&big.Int{}).Mul(_PByte, _Num1000)
	_ZByte = (&big.Int{}).Mul(_EByte, _Num1000)
	_YByte = (&big.Int{}).Mul(_ZByte, _Num1000)
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
	case ReaderTypeJson:
		a.reader = NewJsonReader()
	case ReaderTypeYaml:
		a.reader = NewYamlReader()
	default:
		return nil, ErrNotSupportedReaderType
	}
	if e := a.reader.Read(a.name, &a.configs); e != nil {
		return nil, e
	}

	a.copyDollarSymbol()

	return a, nil
}

func (p *adapterConfig) copyDollarSymbol() {
	p.locker.RLock()
	defer p.locker.RUnlock()

	switch p.readerType {
	case ReaderTypeJson:
		jConfig.copyDollarSymbol(&p.configs, "", &p.configs)
	case ReaderTypeYaml:
		yConfig.copyDollarSymbol(&p.configs)
	}

	return
}

// GetTimeDuration return time in p.configs by key
func (p *adapterConfig) GetTimeDuration(key string, defValue ...time.Duration) time.Duration {
	groups, matched := findStringSubmatchMap(strings.ToLower(p.GetString(key)), timeReg)

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

// GetByteSize return time in p.configs by key
func (p *adapterConfig) GetByteSize(key string) *big.Int {
	groups, matched := findStringSubmatchMap(strings.ToLower(p.GetString(key)), bitReg)

	if matched {
		i, _ := itypes.ToInt64(groups["value"])

		switch groups["unit"] {
		case "b", "byte", "bytes":
			return (&big.Int{}).Mul(big.NewInt(i), _Byte)
		case "kb", "kilobyte", "kilobytes":
			return (&big.Int{}).Mul(big.NewInt(i), _KByte)
		case "mb", "megabyte", "megabytes":
			return (&big.Int{}).Mul(big.NewInt(i), _MByte)
		case "gb", "gigabyte", "gigabytes":
			return (&big.Int{}).Mul(big.NewInt(i), _GByte)
		case "tb", "terabyte", "terabytes":
			return (&big.Int{}).Mul(big.NewInt(i), _TByte)
		case "pb", "petabyte", "petabytes":
			return (&big.Int{}).Mul(big.NewInt(i), _PByte)
		case "eb", "exabyte", "exabytes":
			return (&big.Int{}).Mul(big.NewInt(i), _EByte)
		case "zb", "zettabyte", "zettabytes":
			return (&big.Int{}).Mul(big.NewInt(i), _ZByte)
		case "yb", "yottabyte", "yottabytes":
			return (&big.Int{}).Mul(big.NewInt(i), _YByte)
		case "k", "ki", "kib", "kibibyte", "kibibytes":
			return (&big.Int{}).Mul(big.NewInt(i), _Byte)
		case "m", "mi", "mib", "mebibyte", "mebibytes":
			return (&big.Int{}).Mul(big.NewInt(i), _MiByte)
		case "g", "gi", "gib", "gibibyte", "gibibytes":
			return (&big.Int{}).Mul(big.NewInt(i), _GiByte)
		case "t", "ti", "tib", "tebibyte", "tebibytes":
			return (&big.Int{}).Mul(big.NewInt(i), _TiByte)
		case "p", "pi", "pib", "pebibyte", "pebibytes":
			return (&big.Int{}).Mul(big.NewInt(i), _PiByte)
		case "e", "ei", "eib", "exbibyte", "exbibytes":
			return (&big.Int{}).Mul(big.NewInt(i), _EiByte)
		case "z", "zi", "zib", "zebibyte", "zebibytes":
			return (&big.Int{}).Mul(big.NewInt(i), _ZiByte)
		case "y", "yi", "yib", "yobibyte", "yobibytes":
			return (&big.Int{}).Mul(big.NewInt(i), _YiByte)
		}
	}

	return nil
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
		i, e := itypes.ToInt64(v)
		if e != nil {
			return nil
		}
		items = append(items, int(i))
	}
	return items
}

// GetFloatList return a list of floats in p.configs by key
func (p *adapterConfig) GetFloatList(key string) []float64 {

	var items []float64
	for _, v := range p.GetList(key) {
		f, e := itypes.ToFloat64(v)
		if e != nil {
			return nil
		}
		items = append(items, f)
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
	p.locker.Lock()
	defer p.locker.Unlock()

	return p.reader.Dump(p.configs)
}

// internal functions
// findStringSubmatchMap
func findStringSubmatchMap(s, exp string) (map[string]string, bool) {
	reg := regexp.MustCompile(exp)
	captures := make(map[string]string)

	match := reg.FindStringSubmatch(s)
	if match == nil {
		return captures, false
	}

	for i, name := range reg.SubexpNames() {
		if i == 0 || name == "" {
			continue
		}
		captures[name] = match[i]
	}
	return captures, true
}

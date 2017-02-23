// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"regexp"
	"time"
)

type Config interface {
	GetInterface(key string, defValue ...interface{}) (res interface{})
	GetString(key string, defValue ...string) (res string)
	GetBoolean(key string, defValue ...bool) (b bool)
	GetInt(key string, defValue ...int) (res int)
	GetFloat(key string, defValue ...float64) (res float64)
	GetList(key string) (res []interface{})
	GetStringList(key string) []string
	GetBooleanList(key string) []bool
	GetIntList(key string) []int
	GetFloatList(key string) []float64
	GetTimeDuration(key string, defValue ...time.Duration) time.Duration
	GetConfig(key string) Config
	SetKeyValue(key string, value interface{}) (err error)
	Dump() (bs []byte, err error)
}

func NewConfig(name string) (Config, error) {
	switch fileToReaderType(name) {
	case ReaderTypeJson:
		return NewJsonConfig(name)
	case ReaderTypeYaml:
		return NewYamlConfig(name)
	}

	return nil, ErrUnknownSuffixes
}

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

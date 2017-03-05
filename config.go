// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"math/big"
	"time"
)

//  Config
type Config interface {
	// get a object
	GetInterface(key string, defValue ...interface{}) (res interface{})
	// get a string
	GetString(key string, defValue ...string) (res string)
	// get a bool
	GetBoolean(key string, defValue ...bool) (b bool)
	// get a int
	GetInt(key string, defValue ...int) (res int)
	// get a float
	GetFloat(key string, defValue ...float64) (res float64)
	// get list of objects
	GetList(key string) (res []interface{})
	// get list of strings
	GetStringList(key string) []string
	// get list of bools
	GetBooleanList(key string) []bool
	// get list of ints
	GetIntList(key string) []int
	// get list of float64s
	GetFloatList(key string) []float64
	// get time duration by (int)(uint), exp: 1s, 1day
	GetTimeDuration(key string, defValue ...time.Duration) time.Duration
	// get byte size by (int)(uint), exp: 1k, 1m
	GetByteSize(key string) *big.Int
	// get key config
	GetConfig(key string) Config
	// set key's value into config
	SetKeyValue(key string, value interface{}) (err error)
	// get all config
	Dump() (bs []byte, err error)
}

// NewConfig return Config by file's path, judge path's suffix, supported .json, .yml, .yaml
func NewConfig(name string) (Config, error) {
	switch fileToReaderType(name) {
	case ReaderTypeJson:
		return NewJsonConfig(name)
	case ReaderTypeYaml:
		return NewYamlConfig(name)
	}

	return nil, ErrUnknownSuffixes
}

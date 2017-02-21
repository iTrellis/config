// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package config

type Config interface {
	GetInterface(key string, defValue ...interface{}) (res interface{})
	GetString(key string, defValue ...string) (res string)
	GetBoolean(key string, defValue ...bool) (b bool)
	GetInt(key string, defValue ...int) (res int)
	GetFloat(key string, defValue ...float64) (res float64)
	GetList(key string, defValue ...interface{}) (res interface{})
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

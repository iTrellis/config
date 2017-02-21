// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package config

// ReaderType
type ReaderType int

const (
	ReaderTypeSuffix ReaderType = iota
	ReaderTypeJson
	ReaderTypeXml
	ReaderTypeYaml
)

type Reader interface {
	// read file into model
	Read(path string, model interface{}) error
	// dump configs' cache
	Dump(model interface{}) ([]byte, error)
}

// New return a reader by ReaderType
func New(rt ReaderType) Reader {
	switch rt {
	case ReaderTypeJson:
		return NewJsonReader()
	case ReaderTypeXml:
		return NewXmlReader()
	case ReaderTypeYaml:
		return NewYamlReader()
	default:
		return NewSuffixReader()
	}
}

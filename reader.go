// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package config

type ReaderType int

const (
	ReaderTypeJson ReaderType = iota + 1
	ReaderTypeXml
	ReaderTypeYaml
)

type Reader interface {
	Read(name string, model interface{}) error
}

func New(rt ReaderType) Reader {
	switch rt {
	case ReaderTypeJson:
		return NewJsonReader()
	case ReaderTypeXml:
		return NewXmlReader()
	case ReaderTypeYaml:
		return NewYamlReader()
	default:
		return nil
	}
}

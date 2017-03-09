// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package config

// ReaderType define reader type
type ReaderType int

const (
	// ReaderTypeSuffix judge by file suffix
	ReaderTypeSuffix ReaderType = iota
	// ReaderTypeJSON json reader type
	ReaderTypeJSON
	// ReaderTypeXML xml reader type
	ReaderTypeXML
	// ReaderTypeYAML yaml reader type
	ReaderTypeYAML
)

// Reader reader repo
type Reader interface {
	// read file into model
	Read(path string, model interface{}) error
	// dump configs' cache
	Dump(model interface{}) ([]byte, error)
}

// New return a reader by ReaderType
func New(rt ReaderType) Reader {
	switch rt {
	case ReaderTypeJSON:
		return NewJSONReader()
	case ReaderTypeXML:
		return NewXMLReader()
	case ReaderTypeYAML:
		return NewYAMLReader()
	default:
		return NewSuffixReader()
	}
}

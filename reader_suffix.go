// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"strings"
	"sync"
)

type defSuffixReader struct {
	mu sync.Mutex

	name string
}

// NewJsonReader return a suffix reader
// supportted: .json, .xml, .yaml, .yml
func NewSuffixReader() Reader {
	return &defSuffixReader{}
}

func (p *defSuffixReader) Read(name string, model interface{}) error {
	if name == "" {
		return ErrInvalidFilePath
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	p.name = name

	switch {
	case strings.HasSuffix(p.name, ".json"):
		return jsonReader.Read(p.name, model)
	case strings.HasSuffix(p.name, ".xml"):
		return xmlReader.Read(p.name, model)
	case strings.HasSuffix(p.name, ".yml"), strings.HasSuffix(p.name, ".yaml"):
		return yamlReader.Read(p.name, model)
	default:
		return ErrUnknownSuffixes
	}
}

func (p *defSuffixReader) Dump(v interface{}) ([]byte, error) {
	if p.name == "" {
		return nil, ErrInvalidFilePath
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	switch fileToReaderType(p.name) {
	case ReaderTypeJson:
		return jsonReader.Dump(v)
	case ReaderTypeXml:
		return xmlReader.Dump(v)
	case ReaderTypeYaml:
		return yamlReader.Dump(v)
	}

	return nil, ErrUnknownSuffixes
}

func fileToReaderType(name string) ReaderType {
	switch {
	case strings.HasSuffix(name, ".json"):
		return ReaderTypeJson
	case strings.HasSuffix(name, ".xml"):
		return ReaderTypeXml
	case strings.HasSuffix(name, ".yml"), strings.HasSuffix(name, ".yaml"):
		return ReaderTypeYaml
	}

	return ReaderTypeSuffix
}

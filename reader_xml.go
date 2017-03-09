// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"encoding/xml"
	"sync"
)

type defXMLReader struct {
	mu sync.Mutex
}

var xmlReader = &defXMLReader{}

// NewXMLReader return xml config reader
func NewXMLReader() Reader {
	return xmlReader
}

func (p *defXMLReader) Read(name string, model interface{}) error {
	if name == "" {
		return ErrInvalidFilePath
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	data, err := readFile(name)
	if err != nil {
		return err
	}
	return xml.Unmarshal(data, model)
}

func (*defXMLReader) Dump(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

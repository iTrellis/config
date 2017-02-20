// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"encoding/xml"
	"sync"
)

type defXmlReader struct {
	mu sync.Mutex
}

var xmlReader = &defXmlReader{}

func NewXmlReader() Reader {
	return xmlReader
}

func (p *defXmlReader) Read(name string, model interface{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	data, err := readFile(name)
	if err != nil {
		return err
	}
	return xml.Unmarshal(data, model)
}

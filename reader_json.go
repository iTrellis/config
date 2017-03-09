// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"bytes"
	"encoding/json"
	"sync"
)

type defJSONReader struct {
	mu sync.Mutex

	name string
}

var jsonReader = &defJSONReader{}

// NewJSONReader return a json reader
func NewJSONReader() Reader {
	return jsonReader
}

func (p *defJSONReader) Read(name string, model interface{}) error {
	if name == "" {
		return ErrInvalidFilePath
	}
	p.mu.Lock()
	defer p.mu.Unlock()

	data, err := readFile(name)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(bytes.NewBuffer(data))
	decoder.UseNumber()

	return decoder.Decode(model)
}

func (*defJSONReader) Dump(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

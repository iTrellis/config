// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"encoding/json"
	"sync"
)

type defJsonReader struct {
	mu sync.Mutex

	name string
}

var jsonReader = &defJsonReader{}

// NewJsonReader return a json reader
func NewJsonReader() Reader {
	return jsonReader
}

func (p *defJsonReader) Read(name string, model interface{}) error {
	if name == "" {
		return ErrInvalidFilePath
	}
	p.mu.Lock()
	defer p.mu.Unlock()

	data, err := readFile(name)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, model)
}

func (*defJsonReader) Dump(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

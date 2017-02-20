// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"encoding/json"
	"sync"
)

type defJsonReader struct {
	mu sync.Mutex
}

var jsonReader = &defJsonReader{}

func NewJsonReader() Reader {
	return jsonReader
}

func (p *defJsonReader) Read(name string, model interface{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	data, err := readFile(name)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, model)
}

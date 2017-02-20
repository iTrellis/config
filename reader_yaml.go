// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"sync"

	"gopkg.in/yaml.v2"
)

type defYamlReader struct {
	mu sync.Mutex
}

var yamlReader = &defYamlReader{}

func NewYamlReader() Reader {
	return yamlReader
}

func (p *defYamlReader) Read(name string, model interface{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	data, err := readFile(name)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, model)
}

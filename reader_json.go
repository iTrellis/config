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

	data, err := p.readFile(name)
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

func (p *defJSONReader) readFile(name string) ([]byte, error) {
	data, err := readFile(name)
	if err != nil {
		return nil, err
	}
	var escaped bool // string value flag
	var comments int // 1 line; 2 multi line
	var returns []byte

	length := len(data)
	for i, w := 0, 0; i < length; i += w {
		w = 1

		switch comments {
		case 1:
			if data[i] == '\n' {
				comments = 0
				escaped = false
			}
			continue
		case 2:
			if data[i] != '*' || length == i+1 {
				continue
			}
			if data[i+1] != '/' {
				continue
			}
			w = 2
			comments = 0
			escaped = false
			continue
		}
		switch data[i] {
		case '"':
			{
				if escaped {
					escaped = false
				} else {
					escaped = true
				}
				returns = append(returns, data[i])
			}
		case '/':
			{
				if escaped || length == i+1 {
					returns = append(returns, data[i])
					break
				}
				switch data[i+1] {
				case '/':
					w = 2
					comments = 1
				case '*':
					w = 2
					comments = 2
				default:
					returns = append(returns, data[i])
				}
			}
		default:
			if escaped || !isWhitespace(data[i]) {
				returns = append(returns, data[i])
			}

		}
	}
	return returns, nil
}

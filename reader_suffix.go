// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"errors"
	"strings"
)

var ErrorUnknownSuffixes = errors.New("unknown file with suffix")

type defSuffixReader struct{}

func NewSuffixReader() Reader {
	return (*defSuffixReader)(nil)
}

func (*defSuffixReader) Read(name string, model interface{}) error {
	switch {
	case strings.HasSuffix(name, ".json"):
		return jsonReader.Read(name, model)
	case strings.HasSuffix(name, ".xml"):
		return xmlReader.Read(name, model)
	case strings.HasSuffix(name, ".yml"), strings.HasSuffix(name, ".yaml"):
		return yamlReader.Read(name, model)
	default:
		return ErrorUnknownSuffixes
	}
}

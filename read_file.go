// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package config

import (
	"io/ioutil"
)

func readFile(name string) ([]byte, error) {
	return ioutil.ReadFile(name)
}

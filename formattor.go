package config

import "fmt"

// GetMapKeyValue get value from map[interface{} | string] interface{}
func GetMapKeyValue(ms interface{}, key string) (interface{}, error) {
	m, ok := ms.(map[interface{}]interface{})
	if ok {
		return m[key], nil
	}

	s, ok := ms.(map[string]interface{})
	if ok {
		return s[key], nil
	}

	return nil, fmt.Errorf("config invalid: %v", ms)
}

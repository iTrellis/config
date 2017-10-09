// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package config

// Options initial params
type Options map[string]interface{}

// String return string(value) with key in options
func (p Options) String(key string) (string, error) {
	if 0 == len(p) {
		return "", nil
	}
	return defaultGetter.GetMapKeyValueString(p, key)
}

// Bool return bool(value) with key in options
func (p Options) Bool(key string) (bool, error) {
	if 0 == len(p) {
		return false, nil
	}
	return defaultGetter.GetMapKeyValueBool(p, key)
}

// Int return int(value) with key in options
func (p Options) Int(key string) (int, error) {
	if 0 == len(p) {
		return 0, nil
	}
	return defaultGetter.GetMapKeyValueInt(p, key)
}

// Int64 return int64(value) with key in options
func (p Options) Int64(key string) (int64, error) {
	if 0 == len(p) {
		return 0, nil
	}
	return defaultGetter.GetMapKeyValueInt64(p, key)
}

# Config Reader

Go package for reading cofig file by JSON, XML, YAML.

## Installation

```bash
go get github.com/go-trellis/formats/inner-types
go get gopkg.in/yaml.v2
go get github.com/go-trellis/config
```

```bash
go get -u github.com/go-trellis/config
```

### imports

import [gopkg.in/yaml.v2](https://github.com/go-yaml/yaml)

## Usage

### Config

* dot separator to get values, and if return nil, you should set default
* You can do like this: c.GetString("a.b.c") Or c.GetString("a.b.c", "default")
* Supported: .json, .yaml, .yml

```go
c, e := NewConfig(name)
c.GetString("a.b.c")
```

### Feature

```go
type Config interface {
	GetInterface(key string, defValue ...interface{}) (res interface{})
	GetString(key string, defValue ...string) (res string)
	GetBoolean(key string, defValue ...bool) (b bool)
	GetInt(key string, defValue ...int) (res int)
	GetFloat(key string, defValue ...float64) (res float64)
	GetList(key string) (res []interface{})
	GetStringList(key string) []string
	GetBooleanList(key string) []bool
	GetIntList(key string) []int
	GetFloatList(key string) []float64
	GetTimeDuration(key string, defValue ...time.Duration) time.Duration
	GetConfig(key string) Config
	SetKeyValue(key string, value interface{}) (err error)
	Dump() (bs []byte, err error)
}
```

### More Example

[See More Example](example/suffix.go)


### Reader Repo

```go
type Reader interface {
	// read file into model
	Read(path string, model interface{}) error
	// dump configs' cache
	Dump(model interface{}) ([]byte, error)
}
```

```go
r := New(ReaderType)
if err := r.Read(filename, model); err != nil {
	return
}
```

### Readers

```go
jReader := NewJsonReader()
xReader := NewXmlReader()
yReader := NewYamlReader()
```

* if you want to judge reader by file's suffix

```go
sReader := NewSuffixReader()
```

* .json = NewJsonReader()
* .xml = NewXmlReader()
* .yaml | .yml = NewYamlReader()

### TODO

supported Config .xml: now go encoding/xml not supported map[string]interface{}

# Config Reader

Go package for reading cofig file by JSON, XML, YAML.

## Installation

```bash
go get github.com/go-trellis/config
go get gopkg.in/yaml.v2
```

Or 

```bash
go get -u github.com/go-trellis/config
```

### imports

import [gopkg.in/yaml.v2](https://github.com/go-yaml/yaml)

## Usage

### Config

```go
c, e := NewConfig(name)
```

* GetString
* GetInt
* GetFloat
* GetBoolean
* GetInterface
* GetList
* SetKeyValue
* Dump

[**See More Example**](example/suffix.go)


### Repo

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



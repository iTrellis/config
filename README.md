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

import gopkg.in/yaml.v2

## Usage

### Repo

```go
type Reader interface {
	Read(name string, model interface{}) error
}
```

```go
r := New(ReaderType)
if err := r.Read(filename, model); err != nil {
	return
}

rJson := NewJsonReader()
xJson := NewXmlReader()
yJson := NewYamlReader()
```

* judge reader by file's suffix

```go
suffixJson := NewSuffixReader()
```

* .json = NewJsonReader()
* .xml = NewXmlReader()
* .yaml | .yml = NewYamlReader()

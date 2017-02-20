# Config Reader

Go package for reading cofig file by JSON, XML, YAML.

## TODO

read file config into map[string]interface{}

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
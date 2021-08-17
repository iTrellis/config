package json

import (
	"io"

	jsoniter "github.com/json-iterator/go"
)

func Marshal(v interface{}) ([]byte, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(v)
}

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return jsoniter.MarshalIndent(v, prefix, indent)
}

func MarshalToString(v interface{}) (string, error) {
	return jsoniter.MarshalToString(v)
}

func Unmarshal(bs []byte, v interface{}) error {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(bs, v)
}

func UnmarshalFromString(bs string, v interface{}) error {
	return jsoniter.UnmarshalFromString(bs, v)
}

func NewDecoder(reader io.Reader) *jsoniter.Decoder {
	return jsoniter.NewDecoder(reader)
}

func NewEncoder(writer io.Writer) *jsoniter.Encoder {
	return jsoniter.NewEncoder(writer)
}

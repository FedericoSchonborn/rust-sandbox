package json

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/fdschonborn/go-sandbox/zero"
)

func Parse[T any](data []byte) (value T, err error) {
	err = json.Unmarshal(data, &value)
	return value, err
}

func ParseReader[T any](r io.Reader) (value T, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return zero.Zero[T](), err
	}

	err = json.Unmarshal(data, &value)
	return value, err
}

func ParseObject[T any](data []byte) (value map[string]T, err error) {
	err = json.Unmarshal(data, &value)
	return value, err
}

func ParseObjectReader[T any](r io.Reader) (value map[string]T, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &value)
	return value, err
}

func ParseArray[T any](data []byte) (value []T, err error) {
	err = json.Unmarshal(data, &value)
	return value, err
}

func ParseArrayReader[T any](r io.Reader) (value []T, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &value)
	return value, err
}

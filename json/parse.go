package json

import (
	"encoding/json"
	"io"
)

func Unmarshal[T any](data []byte) (value T, err error) {
	err = json.Unmarshal(data, &value)
	return value, err
}

func Parse[T any](r io.Reader) (value T, err error) {
	data, err := io.ReadAll(r)
	if err != nil {
		var zero T
		return zero, err
	}

	err = json.Unmarshal(data, &value)
	return value, err
}

func UnmarshalObject[T any](data []byte) (value map[string]T, err error) {
	err = json.Unmarshal(data, &value)
	return value, err
}

func ParseObject[T any](r io.Reader) (value map[string]T, err error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &value)
	return value, err
}

func UnmarshalArray[T any](data []byte) (value []T, err error) {
	err = json.Unmarshal(data, &value)
	return value, err
}

func ParseArray[T any](r io.Reader) (value []T, err error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &value)
	return value, err
}

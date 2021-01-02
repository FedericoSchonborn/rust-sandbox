package json

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func Parse(data []byte) (value interface{}, err error) {
	err = json.Unmarshal(data, &value)
	return value, err
}

func ParseReader(r io.Reader) (value interface{}, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &value)
	return value, err
}

func ParseObject(data []byte) (value map[string]interface{}, err error) {
	err = json.Unmarshal(data, &value)
	return value, err
}

func ParseObjectReader(r io.Reader) (value map[string]interface{}, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &value)
	return value, err
}

func ParseArray(data []byte) (value []map[string]interface{}, err error) {
	err = json.Unmarshal(data, &value)
	return value, err
}

func ParseArrayReader(r io.Reader) (value []map[string]interface{}, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &value)
	return value, err
}

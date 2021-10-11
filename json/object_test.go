package json_test

import (
	"testing"

	"github.com/fdschonborn/go-sandbox/json"
)

func TestObject_Get(t *testing.T) {
	result, err := json.Object[string, bool]{"result": true}.Get("result")
	if err != nil {
		t.Fatal(err)
	}

	if !result {
		t.Fatal("Expected result to be true")
	}
}

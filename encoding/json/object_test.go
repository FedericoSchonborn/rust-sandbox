package json_test

import (
	"testing"

	"github.com/fdschonborn/x/encoding/json"
)

func TestObject_Get(t *testing.T) {
	o := json.Object{"result": true}
	t.Log(o)

	var result bool
	if err := o.Get("result", &result); err != nil {
		t.Fatal(err)
	}

	if !result {
		t.Fatal("Expected result to be true")
	}
	t.Log(result)
}

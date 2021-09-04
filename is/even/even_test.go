package even_test

import (
	"fmt"
	"testing"

	"github.com/fdschonborn/go-sandbox/is/even"
)

func TestIsEven(t *testing.T) {
	tt := []struct {
		Value int
		Even  bool
	}{
		{2, true},
		{5, false},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v", tc.Value), func(t *testing.T) {
			result := even.IsEven(tc.Value)
			if result != tc.Even {
				t.Errorf("Expected %t, got %t", tc.Even, result)
			}
		})
	}
}

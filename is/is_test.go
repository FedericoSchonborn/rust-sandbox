package is_test

import (
	"fmt"
	"testing"

	"github.com/fdschonborn/x/is"
)

func TestIsEvenOdd(t *testing.T) {
	tt := []struct {
		Value interface{}
		Even  bool
	}{
		{2, true},
		{5, false},
		{2.0, true},
		{5.0, false},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v", tc.Value), func(t *testing.T) {
			t.Parallel()

			result := is.Even(tc.Value)
			if result != tc.Even {
				t.Errorf("Expected %t, got %t", tc.Even, result)
			}
		})
	}
}

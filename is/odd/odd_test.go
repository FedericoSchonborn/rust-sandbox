package odd_test

import (
	"fmt"
	"testing"

	"github.com/fdschonborn/go-sandbox/is/odd"
)

func TestIsOdd(t *testing.T) {
	tt := []struct {
		Value int
		Even  bool
	}{
		{2, false},
		{5, true},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v", tc.Value), func(t *testing.T) {
			result := odd.IsOdd(tc.Value)
			if result != tc.Even {
				t.Errorf("Expected %t, got %t", tc.Even, result)
			}
		})
	}
}

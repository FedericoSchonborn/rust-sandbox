package or_test

import (
	"testing"

	"github.com/fdschonborn/go-sandbox/or"
)

func TestOr(t *testing.T) {
	tt := []struct {
		Name           string
		V, D, Expected interface{}
	}{
		{"Int", 1, 2, 1},
		{"IntZero", 0, 2, 2},
		{"String", "banana", "pear", "banana"},
		{"StringEmpty", "", "pear", "pear"},
		{"Float", 1.5, 2.2, 1.5},
		{"FloatZero", 0.0, 2.2, 2.2},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			result := or.Or(tc.V, tc.D)
			if result != tc.Expected {
				t.Errorf("Expected value to be %v, got %v instead", tc.Expected, result)
			}
		})
	}
}

package or_test

import (
	"testing"

	"github.com/fdschonborn/go-sandbox/or"
)

type TestTable[T any] []TestCase[T]

type TestCase[T any] struct {
	Name     string
	Value    T
	Default  T
	Expected T
}

func TestOrInt(t *testing.T) {
	tt := TestTable[int]{
		{"Int", 1, 2, 1},
		{"IntZero", 0, 2, 2},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			result := or.Or(tc.Value, tc.Default)
			if result != tc.Expected {
				t.Errorf("Expected value to be %v, got %v instead", tc.Expected, result)
			}
		})
	}
}

func TestOrString(t *testing.T) {
	tt := TestTable[string]{
		{"String", "banana", "pear", "banana"},
		{"StringEmpty", "", "pear", "pear"},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			result := or.Or(tc.Value, tc.Default)
			if result != tc.Expected {
				t.Errorf("Expected value to be %v, got %v instead", tc.Expected, result)
			}
		})
	}
}

func TestOrFloat(t *testing.T) {
	tt := TestTable[float64]{
		{"Float", 1.5, 2.2, 1.5},
		{"FloatZero", 0.0, 2.2, 2.2},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			result := or.Or(tc.Value, tc.Default)
			if result != tc.Expected {
				t.Errorf("Expected value to be %v, got %v instead", tc.Expected, result)
			}
		})
	}
}

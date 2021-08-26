package ops_test

import (
	"testing"

	"github.com/fdschonborn/go-sandbox/ops"
)

func TestRangeContains(t *testing.T) {
	tt := []struct {
		Name       string
		Start, End int
		Item       int
		Expect     bool
	}{
		{"0..15:15", 0, 15, 15, false},
		{"0..15:14", 0, 15, 14, true},
		{"0..15:0", 0, 15, 0, true},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			r := ops.NewRange[int](tc.Start, tc.End)

			result := r.Contains(tc.Item)
			if result != tc.Expect {
				t.Errorf("Expected %v, got %v instead", tc.Expect, result)
			}
		})
	}
}

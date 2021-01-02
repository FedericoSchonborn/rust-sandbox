package ascii_test

import (
	"testing"

	"github.com/fdschonborn/x/ascii"
)

func TestStringAppend(t *testing.T) {
	tt := []struct {
		Start  ascii.String
		End    ascii.String
		Result ascii.String
	}{
		{ascii.Empty(), ascii.Empty(), ascii.Empty()},
		{ascii.Empty(), ascii.FromString("banana"), ascii.FromString("banana")},
	}

	for _, tc := range tt {
		t.Run(tc.Result.String(), func(t *testing.T) {
			result := tc.Start.Append(tc.End)
			if !result.Equals(tc.Result) {
				t.Errorf("Expected %q got %q", tc.Result, result)
			}
		})
	}
}

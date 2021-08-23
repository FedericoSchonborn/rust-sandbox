package strings_test

import (
	"testing"

	"github.com/fdschonborn/go-sandbox/ascii/strings"
)

func TestStringAppend(t *testing.T) {
	tt := []struct {
		Start  strings.String
		End    strings.String
		Result strings.String
	}{
		{strings.New(), strings.New(), strings.New()},
		{strings.New(), strings.From("banana"), strings.From("banana")},
	}

	for _, tc := range tt {
		t.Run(tc.Result.String(), func(t *testing.T) {
			t.Parallel()

			result := tc.Start.Append(tc.End)
			if !result.Equals(tc.Result) {
				t.Errorf("Expected %q got %q", tc.Result, result)
			}
		})
	}
}

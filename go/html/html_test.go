package html_test

import (
	"testing"

	"github.com/fdschonborn/sandbox/go/html/a"
	"github.com/fdschonborn/sandbox/go/html/p"
	"github.com/fdschonborn/sandbox/go/html/span"
)

func Test(t *testing.T) {
	p.New(
		"Here is a ",
		a.New("https://go.dev", "link").
			Target(a.TargetBlank),
		", and a ",
		span.New("span"),
		".",
	)
}

package html_test

import (
	"testing"

	"github.com/fdschonborn/sandbox/html/a"
	"github.com/fdschonborn/sandbox/html/p"
	"github.com/fdschonborn/sandbox/html/span"
)

func Test(t *testing.T) {
	p.New(
		"Here is a ",
		a.New("https://go.dev", "link").
			Target(a.Blank),
		", and a ",
		span.New("span"),
		".",
	)
}

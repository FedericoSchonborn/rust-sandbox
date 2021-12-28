package html_test

import (
	"testing"

	"github.com/fdschonborn/go-sandbox/html/a"
	"github.com/fdschonborn/go-sandbox/html/p"
	"github.com/fdschonborn/go-sandbox/html/span"
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

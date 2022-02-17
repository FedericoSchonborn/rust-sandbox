package html_test

import (
	"testing"

	"github.com/fdschonborn/sandbox/go/html/a"
	"github.com/fdschonborn/sandbox/go/html/p"
	"github.com/fdschonborn/sandbox/go/html/span"
	"github.com/fdschonborn/sandbox/go/html/text"
)

func Test(t *testing.T) {
	p.New(
		text.New("Here is a "),
		a.New("https://go.dev", text.New("link")).
			Target(a.Blank),
		text.New(", and a "),
		span.New(text.New("span")),
		text.New("."),
	)
}

package html_test

import (
	"testing"

	"github.com/FedericoSchonborn/sandbox/go/html/a"
	"github.com/FedericoSchonborn/sandbox/go/html/p"
	"github.com/FedericoSchonborn/sandbox/go/html/span"
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

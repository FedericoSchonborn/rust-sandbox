package html_test

import (
	"testing"

	"github.com/fdschonborn/go-sandbox/html"
)

func Test(t *testing.T) {
	html.P(
		"Here is a ",
		html.A("https://go.dev", "link").
			Target(html.ATargetBlank),
		", and a ",
		html.Span("span"),
	)
}

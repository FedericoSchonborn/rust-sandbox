package html_test

import (
	"bytes"
	"testing"

	. "github.com/fdschonborn/x/html"
)

func TestDocument(t *testing.T) {
	doc := New(
		Head(),
		Body(
			Div(
				H1("Hello, world!"),
			),
			Div(
				P("This is a test"),
			),
		),
	)
	t.Logf("%#v", doc)

	buffer := bytes.NewBuffer(nil)
	if err := doc.Render(buffer); err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", buffer)
}

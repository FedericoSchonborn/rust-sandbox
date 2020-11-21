package html_test

import (
	"bytes"
	"testing"

	. "github.com/fdschonborn/x/go/html"
)

func TestDocument(t *testing.T) {
	doc := New(
		Head(),
		Body(
			Div(
				H(1, "Hello, world!"),
			).Class("title"),
		),
	)
	t.Logf("%#v", doc)

	buffer := bytes.NewBuffer(nil)
	if err := doc.Render(buffer); err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", buffer)
}

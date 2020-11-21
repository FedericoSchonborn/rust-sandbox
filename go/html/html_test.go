package html_test

import (
	"testing"

	"github.com/fdschonborn/x/go/html"
)

func TestDocument(t *testing.T) {
	elem := html.Document{
		Body: html.Body{
			Content: []html.Element{
				&html.Header{
					Text: "Hello, world!",
					Rank: 1,
				},
			},
		},
	}
	t.Logf("%#v", elem)
}

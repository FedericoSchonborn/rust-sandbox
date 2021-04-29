package html2_test

import (
	"github.com/fdschonborn/x/html2/body"
	"github.com/fdschonborn/x/html2/document"
	"github.com/fdschonborn/x/html2/p"
)

func Example() {
	document.New(
		document.Body(
			body.New(
				body.Children(
					p.New(
						p.Text("Hello, world!"),
					),
				),
			),
		),
	)
}

package test1_test

import (
	"github.com/fdschonborn/sandbox/html/test1/body"
	"github.com/fdschonborn/sandbox/html/test1/document"
	"github.com/fdschonborn/sandbox/html/test1/p"
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

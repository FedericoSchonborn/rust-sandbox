package test2_test

import (
	"github.com/fdschonborn/sandbox/html/test2/body"
	"github.com/fdschonborn/sandbox/html/test2/document"
	"github.com/fdschonborn/sandbox/html/test2/p"
)

func Example() {
	document.New().
		Body(
			body.New().
				Children(
					p.New().
						Text("Hello, world!"),
				),
		)
}

package html3_test

import (
	"github.com/fdschonborn/x/html3/body"
	"github.com/fdschonborn/x/html3/document"
	"github.com/fdschonborn/x/html3/p"
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

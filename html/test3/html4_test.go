package test3_test

import (
	. "github.com/fdschonborn/sandbox/html/test3"
)

func Example() {
	Body(
		Children(
			P(
				Text("banana"),
			),
		),
	)
}

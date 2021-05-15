package test3_test

import (
	. "github.com/fdschonborn/x/html/test3"
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

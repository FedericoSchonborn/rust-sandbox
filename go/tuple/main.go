//go:build generator

package main

// THIS HURTS TO WATCH, YES, I KNOW.

// TODO:
// - Convert to "text/template"

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	file, err := os.Create("go/tuple/tuple.go")
	if err != nil {
		panic(err)
	}

	if _, err := io.WriteString(file, "// Package generated by tuple/main.go, DO NOT EDIT!\n\npackage tuple\n"); err != nil {
		panic(err)
	}

	for n := 0; n < 128; n++ {
		typeName := "Tuple" + strconv.Itoa(n+1)
		funcName := "New" + strconv.Itoa(n+1)

		var fields string
		var generics string
		var params string
		var args string
		var assign string
		var accesses string
		var returns string
		for index := 0; index < n+1; index++ {
			upTName := "T" + strconv.Itoa(index)
			upVName := "V" + strconv.Itoa(index)
			downVName := "v" + strconv.Itoa(index)

			fields += upVName + " " + upTName
			generics += upTName
			params += upTName
			args += downVName + " " + upTName
			assign += downVName
			accesses += "t." + upVName
			returns += upTName

			if index != n {
				fields += ";"
				generics += ","
				params += ","
				args += ", "
				assign += ","
				accesses += ","
				returns += ","
			}

			if index == n {
				generics += " any"
			}
		}

		if _, err := fmt.Fprintf(file, `
// %[1]s is a tuple containing %[10]d value(s).
type %[1]s[%[2]s] struct {
	%[3]s
}

// %[9]s creates a new tuple of %[10]d value(s).
func %[9]s[%[2]s](%[5]s) *%[1]s[%[4]s] {
	return &%[1]s[%[4]s]{%[6]s`+", "+`}
}
`, typeName, generics, fields, params, args, assign, returns, accesses, funcName, n+1); err != nil {
			panic(err)
		}
	}

	cmd := exec.Command("go", "fmt", "./go/tuple")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

package fmt

import (
	"fmt"
	"io"
)

func Printfln(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format+"\n", a...)
}

func Fprintfln(w io.Writer, format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, format+"\n", a...)
}

func Sprintfln(format string, a ...interface{}) string {
	return fmt.Sprintf(format+"\n", a...)
}

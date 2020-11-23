package fmt

import (
	"fmt"
	"os"
)

func Eprint(a ...interface{}) (n int, err error) {
	return fmt.Fprint(os.Stderr, a...)
}

func Eprintf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(os.Stderr, format, a...)
}

func Eprintln(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(os.Stderr, a...)
}

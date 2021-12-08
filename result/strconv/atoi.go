package strconv

import (
	"strconv"

	. "github.com/fdschonborn/go-sandbox/result"
)

func Atoi(s string) Result[int, error] {
	v, err := strconv.Atoi(s)
	if err != nil {
		return Err[int, error](err)
	}

	return Ok[int, error](v)
}

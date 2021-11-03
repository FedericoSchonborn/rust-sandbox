package strconv

import (
	"strconv"

	"github.com/fdschonborn/go-sandbox/result"
)

func Atoi(s string) result.Result[int, error] {
	v, err := strconv.Atoi(s)
	if err != nil {
		return result.Err[int](err)
	}

	return result.Ok[int, error](v)
}

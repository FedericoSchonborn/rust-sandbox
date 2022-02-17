package strconv

import (
	"strconv"

	"github.com/fdschonborn/sandbox/go/result"
)

func Atoi(s string) result.Result[int] {
	v, err := strconv.Atoi(s)
	if err != nil {
		return result.Err[int](err)
	}

	return result.Ok(v)
}

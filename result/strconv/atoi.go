package strconv

import (
	"strconv"

	. "github.com/fdschonborn/go-sandbox/result"
)

func Atoi(s string) Result[int] {
	v, err := strconv.Atoi(s)
	if err != nil {
		return Err[int](err)
	}

	return Ok(v)
}

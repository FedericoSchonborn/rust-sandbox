package ascii

import (
	"github.com/fdschonborn/go-sandbox/ascii"
	"github.com/fdschonborn/go-sandbox/option"
)

func ToAscii(r rune) option.Option[byte] {
	if !ascii.IsAscii(r) {
		return option.None[byte]()
	}

	return option.Some(byte(r))
}

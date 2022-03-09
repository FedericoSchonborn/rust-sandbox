package ascii

import (
	"github.com/FedericoSchonborn/sandbox/go/ascii"
	"github.com/FedericoSchonborn/sandbox/go/option"
)

func ToAscii(r rune) option.Option[byte] {
	if !ascii.IsAscii(r) {
		return option.None[byte]()
	}

	return option.Some(byte(r))
}

package iter

import "github.com/fdschonborn/go-sandbox/option"

type runes struct {
	inner []rune
	len   int
	index int
}

func Runes(s string) Iterator[rune] {
	inner := []rune(s)

	return &runes{
		inner: inner,
		len:   len(inner),
		index: 0,
	}
}

func (b *runes) Next() option.Option[rune] {
	if b.index >= b.len {
		return option.None[rune]()
	}

	result := b.inner[b.index]
	b.index++
	return option.Some(result)
}

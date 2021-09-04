package iter

import "github.com/fdschonborn/go-sandbox/option"

type bytes struct {
	inner []byte
	len   int
	index int
}

func Bytes(s string) Iterator[byte] {
	inner := []byte(s)

	return &bytes{
		inner: inner,
		len:   len(inner),
		index: 0,
	}
}

func (b *bytes) Next() option.Option[byte] {
	if b.index >= b.len {
		return option.None[byte]()
	}

	result := b.inner[b.index]
	b.index++
	return option.Some(result)
}

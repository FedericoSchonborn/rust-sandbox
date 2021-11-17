package iter

type Peeker[Item any] interface {
	Iterator[Item]

	Peek() (_ Item, ok bool)
}

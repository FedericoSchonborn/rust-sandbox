package tuple

type Tuple2[A, B any] struct {
	A A
	B B
}

func NewTuple2[A, B any](a A, b B) Tuple2[A, B] {
	return Tuple2[A, B]{A: a, B: b}
}

func (t2 Tuple2[A, B]) Unpack() (A, B) {
	return t2.A, t2.B
}

type Tuple3[A, B, C any] struct {
	A A
	B B
	C C
}

func NewTuple3[A, B, C any](a A, b B, c C) Tuple3[A, B, C] {
	return Tuple3[A, B, C]{A: a, B: b, C: c}
}

func (t3 Tuple3[A, B, C]) Unpack() (A, B, C) {
	return t3.A, t3.B, t3.C
}

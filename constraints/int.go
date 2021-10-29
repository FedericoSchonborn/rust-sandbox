package constraints

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Int interface {
	Signed | Unsigned
}

func Add[L, R, O Int](l L, r R) O {
	return O(l) + O(r)
}

func Sub[L, R, O Int](l L, r R) O {
	return O(l) - O(r)
}

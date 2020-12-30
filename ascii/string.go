package ascii

// String is an ASCII-only string.
type String struct {
	buf []byte
	len int
	cap int
}

func (s String) Len() int {
	return s.len
}

func (s String) Cap() int {
	return s.cap
}

func (s String) String() string {
	return string(s.buf)
}

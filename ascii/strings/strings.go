package strings

// String is an ASCII-only string.
type String struct {
	buf []byte
	len int
	cap int
}

func From(s String) String {
	return s
}

func FromString(s string) String {
	return FromBytes([]byte(s))
}

func FromBytes(buf []byte) String {
	return String{
		buf: buf,
		len: len(buf),
		cap: cap(buf),
	}
}

func Empty() String {
	return String{
		buf: make([]byte, 0),
		len: 0,
		cap: 0,
	}
}

func Make(size, cap int) String {
	if cap < size {
		cap = size
	}

	return String{
		buf: make([]byte, size, cap),
		len: 0,
		cap: cap,
	}
}

func (s String) Append(sb String) String {
	totalLen := len(s.buf) + len(sb.buf)
	sr := Make(totalLen, 0)
	sr.len = totalLen
	if len(s.buf) > 0 {
		for i, c := range s.buf {
			sr.buf[i] = c
		}
	}

	if len(sb.buf) > 0 {
		for i, c := range sb.buf {
			sr.buf[len(s.buf)+i] = c
		}
	}

	return sr
}

func (s String) Equals(sb String) bool {
	for i, c := range sb.buf {
		if s.buf[i] != c {
			return false
		}
	}

	return true
}

func (s String) Bytes() []byte {
	return s.buf
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

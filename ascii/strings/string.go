package strings

// String is an ASCII-only string.
type String struct {
	buf []byte
	len int
	cap int
}

func New() String {
	return String{
		buf: []byte{},
	}
}

func Make(cap int) String {
	return String{
		buf: make([]byte, 0, cap),
		cap: cap,
	}
}

func From(s string) String {
	return FromBytes([]byte(s))
}

func FromBytes(buf []byte) String {
	return String{
		buf: buf,
		len: len(buf),
		cap: cap(buf),
	}
}

func (s String) Append(sb String) String {
	totalLen := len(s.buf) + len(sb.buf)
	sr := Make(totalLen)
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

func (s String) Clone() String {
	return String{
		s.buf,
		s.len,
		s.cap,
	}
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

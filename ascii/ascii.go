package ascii

const (
	MaxByte byte = 255
)

// ISO C

func IsAlnum(c byte) bool {
	return IsAlpha(c) || IsDigit(c)
}

func IsAlpha(c byte) bool {
	return IsUpper(c) || IsLower(c)
}

func IsBlank(c byte) bool {
	return c == ' ' || c == '\t'
}

func IsCntrl(c byte) bool {
	return c == 0x7F || (c >= 0x00 && c <= 0x1F)
}

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func IsGraph(c byte) bool {
	return c != ' ' && IsPrint(c)
}

func IsLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func IsPrint(c byte) bool {
	return c != 0x7F && c >= 0x1F
}

func IsPunct(c byte) bool {
	return !IsAlnum(c) && IsGraph(c)
}

func IsSpace(c byte) bool {
	return IsBlank(c) || c == '\n' || c == '\v' || c == '\f' || c == '\r'
}

func IsUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func IsXdigit(c byte) bool {
	return IsDigit(c) || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
}

func ToLower(c byte) byte {
	if !IsAlpha(c) {
		return c
	}

	return c + 0x20
}

func ToUpper(c byte) byte {
	if !IsAlpha(c) {
		return c
	}

	return c - 0x20
}

// POSIX C

func IsAscii(r rune) bool {
	return r >= 0 && r <= 255
}

func ToAscii(r rune) (byte, bool) {
	if !IsAscii(r) {
		return 0, false
	}

	return byte(r), true
}

// Extension

type Type uint8

const (
	TypeAlnum Type = iota << 1
	TypeAlpha
	TypeBlank
	TypeCntrl
	TypeDigit
	TypeGraph
	TypeLower
	TypePrint
	TypePunct
	TypeSpace
	TypeUpper
	TypeXdigit

	TypeAll = TypeAlnum | TypeAlpha | TypeBlank | TypeCntrl | TypeDigit | TypeGraph | TypeLower | TypePrint | TypePunct | TypeSpace | TypeUpper | TypeXdigit
)

func TypeOf(c byte) Type {
	var typ Type

	if IsAlnum(c) {
		typ |= TypeAlnum
	}

	if IsAlpha(c) {
		typ |= TypeAlpha
	}

	if IsBlank(c) {
		typ |= TypeBlank
	}

	if IsCntrl(c) {
		typ |= TypeCntrl
	}

	if IsDigit(c) {
		typ |= TypeDigit
	}

	if IsGraph(c) {
		typ |= TypeGraph
	}

	if IsLower(c) {
		typ |= TypeLower
	}

	if IsPrint(c) {
		typ |= TypePrint
	}

	if IsPunct(c) {
		typ |= TypePunct
	}

	if IsSpace(c) {
		typ |= TypeSpace
	}

	if IsUpper(c) {
		typ |= TypeUpper
	}

	if IsXdigit(c) {
		typ |= TypeXdigit
	}

	return typ
}

package ascii

const (
	MaxByte byte = 255
)

// <ctype.h>

func IsAlnum(c byte) bool

func IsAlpha(c byte) bool

func IsBlank(c byte) bool

func IsCntrl(c byte) bool

func IsDigit(c byte) bool

func IsGraph(c byte) bool

func IsLower(c byte) bool

func IsPrint(c byte) bool

func IsPunct(c byte) bool

func IsSpace(c byte) bool

func IsUpper(c byte) bool

func IsXdigit(c byte) bool

func ToLower(c byte) byte

func ToUpper(c byte) byte

// POSIX

func IsAscii(r rune) bool

func ToAscii(r rune) byte

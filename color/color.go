package color

type Color int32

func RGB(red, green, blue uint8) Color {
	return ARGB(255, red, green, blue)
}

func ARGB(alpha, red, green, blue uint8) Color {
	return Color(0) // TODO
}

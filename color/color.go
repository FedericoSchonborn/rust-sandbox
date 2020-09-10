package color

type Color uint32

func RGB(red, green, blue uint8) Color {
	return RGBA(red, green, blue, 255) // TODO
}

func RGBA(red, green, blue, alpha uint8) Color {
	return Color(0x000000FF) // TODO
}

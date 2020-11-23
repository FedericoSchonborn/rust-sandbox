package html

type Element interface {
	Render(*Context) error
}

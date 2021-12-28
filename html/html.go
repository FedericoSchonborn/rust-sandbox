package html

type Element interface {
	Render() error
}

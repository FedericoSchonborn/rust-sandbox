package option

type T = interface{}
type U = interface{}
type OptionT = Option
type OptionU = Option

type Option struct {
	ok    bool
	value interface{}
}

func Some(value interface{}) Option {
	return Option{ok: true, value: value}
}

// None and Option{} are equivalent.
func None() Option {
	return Option{ok: false, value: nil}
}

func (o Option) IsSome() bool {
	return o.ok
}

func (o Option) IsNone() bool {
	return !o.ok
}

func (o Option) Map(fn func(value T) U) OptionU {
	if !o.ok {
		return None()
	}

	return Some(fn(o.value))
}

func (o Option) Or(ob OptionT) OptionT {
	if !o.ok {
		return ob
	}

	return o
}

func (o Option) OrElse(fn func() OptionT) OptionT {
	if !o.ok {
		return fn()
	}

	return o
}

package option

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

func (o Option) Map(fn func(value interface{}) Option) Option {
	if !o.ok {
		return None()
	}

	return fn(o.value)
}

func (o Option) Or(value interface{}) interface{} {
	if o.ok {
		return o.value
	}

	return value
}

func (o Option) OrElse(fn func() interface{}) interface{} {
	if o.ok {
		return o.value
	}

	return fn()
}

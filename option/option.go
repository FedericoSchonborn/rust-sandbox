package option

type Option struct {
	value interface{}
	ok    bool
}

func Some(value interface{}) Option {
	return Option{value: value, ok: true}
}

func None() Option {
	return Option{value: nil, ok: false}
}

func (o Option) IsSome() bool {
	return o.ok
}

func (o Option) IsNone() bool {
	return o.ok
}

func (o Option) MapSome(fn func(value interface{}) Option) Option {
	if !o.ok {
		return None()
	}

	return fn(o.value)
}

func (o Option) MapNone(fn func() Option) Option {
	if o.ok {
		return None()
	}

	return fn()
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

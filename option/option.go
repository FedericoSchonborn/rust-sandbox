package option

type T = interface{}
type U = interface{}
type OptionT = Option
type OptionU = Option

type Option struct {
	ok    bool
	value T
}

func Some(value T) OptionT {
	return OptionT{ok: true, value: value}
}

// None and Option{} are equivalent.
func None() OptionT {
	return OptionT{ok: false, value: nil}
}

func (o OptionT) IsSome() bool {
	return o.ok
}

func (o OptionT) IsNone() bool {
	return !o.ok
}

func (o OptionT) Unwrap() T {
	if !o.ok {
		panic("Called OptionT.Unwrap on None value")
	}

	return o.value
}

func (o OptionT) UnwrapOr(def T) T {
	if !o.ok {
		return def
	}

	return o.value
}

func (o OptionT) UnwrapOrElse(fn func() T) T {
	if !o.ok {
		return fn()
	}

	return o.value
}

func (o OptionT) Map(fn func(value T) U) OptionU {
	if !o.ok {
		return None()
	}

	return Some(fn(o.value))
}

func (o OptionT) OkOr(err error) (T, error) {
	if !o.ok {
		return nil, err
	}

	return o.value, nil
}

func (o OptionT) OkOrElse(fn func() error) (T, error) {
	if !o.ok {
		return nil, fn()
	}

	return o.value, nil
}

func (o OptionT) Filter(pred func(T) bool) OptionT {
	if !o.ok {
		return None()
	}

	if !pred(o.value) {
		return None()
	}

	return Some(o.value)
}

func (o OptionT) Or(ob OptionT) OptionT {
	if !o.ok {
		return ob
	}

	return o
}

func (o OptionT) OrElse(fn func() OptionT) OptionT {
	if !o.ok {
		return fn()
	}

	return o
}

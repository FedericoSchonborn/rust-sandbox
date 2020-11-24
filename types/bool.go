package types

type Bool bool

func (b Bool) IsTrue() Bool {
	return b
}

func (b Bool) IsFalse() Bool {
	return !b
}

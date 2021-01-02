package ops

type BoundTag int

const (
	BoundTagIncluded  BoundTag = 1
	BoundTagExcluded  BoundTag = 2
	BoundTagUnbounded BoundTag = 3
)

type BoundValue interface {
	boundValue()
}

type BoundValueIncluded int

func (BoundValueIncluded) boundValue() {}

type BoundValueExcluded int

func (BoundValueExcluded) boundValue() {}

type BoundValueUnbounded struct{}

func (BoundValueUnbounded) boundValue() {}

type Bound struct {
	t BoundTag
	v BoundValue
}

func BoundIncluded(value int) Bound {
	return Bound{
		t: BoundTagIncluded,
		v: BoundValueIncluded(value),
	}
}

func BoundExcluded(value int) Bound {
	return Bound{
		t: BoundTagExcluded,
		v: BoundValueExcluded(value),
	}
}

func BoundUnbounded() Bound {
	return Bound{
		t: BoundTagUnbounded,
		v: BoundValueUnbounded{},
	}
}

func (b Bound) Tag() BoundTag {
	return b.t
}

func (b Bound) Value() BoundValue {
	return b.v
}

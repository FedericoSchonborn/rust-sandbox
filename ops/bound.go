package ops

type BoundType int

const (
	BoundIncludedType  BoundType = 1
	BoundExcludedType  BoundType = 2
	BoundUnboundedType BoundType = 3
)

type BoundValue interface {
	boundValue()
}

type BoundIncludedValue int

func (BoundIncludedValue) boundValue() {}

type BoundExcludedValue int

func (BoundExcludedValue) boundValue() {}

type BoundUnboundedValue struct{}

func (BoundUnboundedValue) boundValue() {}

type Bound struct {
	t BoundType
	v BoundValue
}

func BoundIncluded(value int) Bound {
	return Bound{
		t: BoundIncludedType,
		v: BoundIncludedValue(value),
	}
}

func BoundExcluded(value int) Bound {
	return Bound{
		t: BoundExcludedType,
		v: BoundExcludedValue(value),
	}
}

func BoundUnbounded() Bound {
	return Bound{
		t: BoundUnboundedType,
		v: BoundUnboundedValue{},
	}
}

func (b Bound) Type() BoundType {
	return b.t
}

func (b Bound) Value() BoundValue {
	return b.v
}

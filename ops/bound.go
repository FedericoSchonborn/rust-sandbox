package ops

type BoundTag int

const (
	BoundTagIncluded  BoundTag = 1
	BoundTagExcluded  BoundTag = 2
	BoundTagUnbounded BoundTag = 3
)

type Bound interface {
	boundTag() BoundTag
}

type BoundIncluded int

func NewBoundIncluded(value int) Bound {
	return BoundIncluded(value)
}

func (BoundIncluded) boundTag() BoundTag { return BoundTagIncluded }

type BoundExcluded int

func NewBoundExcluded(value int) Bound {
	return BoundExcluded(value)
}

func (BoundExcluded) boundTag() BoundTag { return BoundTagExcluded }

type BoundUnbounded struct{}

func NewBoundUnbounded() Bound {
	return BoundUnbounded{}
}

func (BoundUnbounded) boundTag() BoundTag { return BoundTagUnbounded }

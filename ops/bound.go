package ops

type boundTag byte

const (
	boundIncluded  boundTag = 1
	boundExcluded  boundTag = 2
	boundUnbounded boundTag = 3
)

type Bound interface {
	boundTag() boundTag
}

type BoundIncluded int

func (bi BoundIncluded) boundTag() boundTag {
	return boundIncluded
}

type BoundExcluded int

func (be BoundExcluded) boundTag() boundTag {
	return boundExcluded
}

type BoundUnbounded struct{}

func (bu BoundUnbounded) boundTag() boundTag {
	return boundUnbounded
}

package ranges

type Bounds interface {
	StartBound() Bound
	EndBound() Bound
	Contains(int) bool
}

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

var _ Bounds = Range{}

type Range struct {
	Start, End int
}

func New(start, end int) Range {
	return Range{Start: start, End: end}
}

func (r Range) StartBound() Bound {
	return BoundIncluded(r.Start)
}

func (r Range) EndBound() Bound {
	return BoundExcluded(r.End)
}

func (r Range) Contains(index int) bool {
	return (r.Start <= index) && (index < r.End)
}

var _ Bounds = Inclusive{}

type Inclusive struct {
	Start, End int
}

func NewInclusive(start, end int) Inclusive {
	return Inclusive{Start: start, End: end}
}

func (ir Inclusive) StartBound() Bound {
	return BoundIncluded(ir.Start)
}

func (ir Inclusive) EndBound() Bound {
	return BoundIncluded(ir.End)
}

func (ir Inclusive) Contains(index int) bool {
	return (ir.Start <= index) && (index <= ir.End)
}

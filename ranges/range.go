package ranges

type Bounds interface {
	StartBound() Bound
	EndBound() Bound
	Contains(int) bool
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

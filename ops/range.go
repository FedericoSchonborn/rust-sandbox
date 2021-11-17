package ops

type RangeBounds interface {
	StartBound() Bound
	EndBound() Bound
	Contains(int) bool
}

var _ RangeBounds = Range{}

type Range struct {
	Start, End int
}

func NewRange(start, end int) Range {
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

var _ RangeBounds = InclusiveRange{}

type InclusiveRange struct {
	Start, End int
}

func NewInclusiveRange(start, end int) InclusiveRange {
	return InclusiveRange{Start: start, End: end}
}

func (ir InclusiveRange) StartBound() Bound {
	return BoundIncluded(ir.Start)
}

func (ir InclusiveRange) EndBound() Bound {
	return BoundIncluded(ir.End)
}

func (ir InclusiveRange) Contains(index int) bool {
	return (ir.Start <= index) && (index <= ir.End)
}

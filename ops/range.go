package ops

type RangeBounds interface {
	StartBound() Bound
	EndBound() Bound
	Contains(item int) bool
}

func defaultRangeBoundsContains(rb RangeBounds, item int) bool {
	var start bool
	switch v := rb.StartBound().(type) {
	case BoundIncluded:
		start = int(v) <= item
	case BoundExcluded:
		start = int(v) < item
	case BoundUnbounded:
		start = true
	}

	var end bool
	switch v := rb.EndBound().(type) {
	case BoundIncluded:
		end = item <= int(v)
	case BoundExcluded:
		end = item < int(v)
	case BoundUnbounded:
		end = true
	}

	return start && end
}

var _ RangeBounds = Range{}

type Range struct {
	Start int
	End   int
}

func NewRange(start, end int) Range {
	return Range{
		Start: start,
		End:   end,
	}
}

func (r Range) StartBound() Bound {
	return BoundIncluded(r.Start)
}

func (r Range) EndBound() Bound {
	return BoundExcluded(r.End)
}

func (r Range) Contains(item int) bool {
	return defaultRangeBoundsContains(r, item)
}

var _ RangeBounds = InclusiveRange{}

type InclusiveRange struct {
	Start int
	End   int
}

func NewInclusiveRange(start, end int) InclusiveRange {
	return InclusiveRange{
		Start: start,
		End:   end,
	}
}

func (ir InclusiveRange) StartBound() Bound {
	return BoundIncluded(ir.Start)
}

func (ir InclusiveRange) EndBound() Bound {
	return BoundIncluded(ir.End)
}

func (ir InclusiveRange) Contains(item int) bool {
	return defaultRangeBoundsContains(ir, item)
}

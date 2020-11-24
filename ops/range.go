package ops

type BoundType int

const (
	BoundTypeIncluded  BoundType = 1
	BoundTypeExcluded  BoundType = 2
	BoundTypeUnbounded BoundType = 3
)

type BoundValue interface {
	boundValue()
}

type BoundIncluded struct {
	Value int
}

func (bi BoundIncluded) boundValue() {}

type BoundExcluded struct {
	Value int
}

func (be BoundExcluded) boundValue() {}

type BoundUnbounded struct{}

func (bu BoundUnbounded) boundValue() {}

type Bound struct {
	t BoundType
	v BoundValue
}

func NewBoundIncluded(value int) Bound {
	return Bound{
		t: BoundTypeIncluded,
		v: BoundIncluded{
			Value: value,
		},
	}
}

func NewBoundExcluded(value int) Bound {
	return Bound{
		t: BoundTypeExcluded,
		v: BoundExcluded{
			Value: value,
		},
	}
}

func NewBoundUnbounded() Bound {
	return Bound{
		t: BoundTypeUnbounded,
		v: BoundUnbounded{},
	}
}

func (b Bound) Type() BoundType {
	return b.t
}

func (b Bound) Value() BoundValue {
	return b.v
}

type RangeBounds interface {
	StartBound() Bound
	EndBound() Bound
	Contains(item int) bool
}

func DefaultRangeBoundsContains(rb RangeBounds, item int) bool {
	var start bool
	startBound := rb.StartBound()
	switch v := startBound.Value().(type) {
	case BoundIncluded:
		start = v.Value <= item
	case BoundExcluded:
		start = v.Value < item
	case BoundUnbounded:
		start = true
	}

	var end bool
	endBound := rb.EndBound()
	switch v := endBound.Value().(type) {
	case BoundIncluded:
		end = item <= v.Value
	case BoundExcluded:
		end = item < v.Value
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
	return NewBoundIncluded(r.Start)
}

func (r Range) EndBound() Bound {
	return NewBoundExcluded(r.End)
}

func (r Range) Contains(item int) bool {
	return DefaultRangeBoundsContains(r, item)
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
	return NewBoundIncluded(ir.Start)
}

func (ir InclusiveRange) EndBound() Bound {
	return NewBoundIncluded(ir.End)
}

func (ir InclusiveRange) Contains(item int) bool {
	return DefaultRangeBoundsContains(ir, item)
}

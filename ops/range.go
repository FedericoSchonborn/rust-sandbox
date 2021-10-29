package ops

import "github.com/fdschonborn/go-sandbox/constraints"

type RangeBounds[Index constraints.Int] interface {
	StartBound() Bound
	EndBound() Bound
	Contains(index Index) bool
}

var _ RangeBounds[int] = Range[int]{}

type Range[Index constraints.Int] struct {
	Start, End Index
}

func NewRange[Index constraints.Int](start, end Index) Range[Index] {
	return Range[Index]{Start: start, End: end}
}

func (r Range[Index]) StartBound() Bound {
	return BoundIncluded(r.Start)
}

func (r Range[Index]) EndBound() Bound {
	return BoundExcluded(r.End)
}

func (r Range[Index]) Contains(index Index) bool {
	return (r.Start <= index) && (index < r.End)
}

var _ RangeBounds[int] = InclusiveRange[int]{}

type InclusiveRange[Index constraints.Int] struct {
	Start, End Index
}

func NewInclusiveRange[Index constraints.Int](start, end Index) InclusiveRange[Index] {
	return InclusiveRange[Index]{Start: start, End: end}
}

func (ir InclusiveRange[Index]) StartBound() Bound {
	return BoundIncluded(ir.Start)
}

func (ir InclusiveRange[Index]) EndBound() Bound {
	return BoundIncluded(ir.End)
}

func (ir InclusiveRange[Index]) Contains(index Index) bool {
	return (ir.Start <= index) && (index <= ir.End)
}

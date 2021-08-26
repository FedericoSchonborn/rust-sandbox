package ops

import "github.com/fdschonborn/go-sandbox/ints"

type Integer = ints.Int

type RangeBounds[T Integer] interface {
	StartBound() Bound
	EndBound() Bound
	Contains(item T) bool
}

var _ RangeBounds[int] = Range[int]{}

type Range[T Integer] struct {
	Start, End T
}

func NewRange[T Integer](start, end T) Range[T] {
	return Range[T]{Start: start, End: end}
}

func (r Range[T]) StartBound() Bound {
	return BoundIncluded(r.Start)
}

func (r Range[T]) EndBound() Bound {
	return BoundExcluded(r.End)
}

func (r Range[T]) Contains(item T) bool {
	return (r.Start <= item) && (item < r.End)
}

var _ RangeBounds[int] = InclusiveRange[int]{}

type InclusiveRange[T Integer] struct {
	Start, End T
}

func NewInclusiveRange[T Integer](start, end T) InclusiveRange[T] {
	return InclusiveRange[T]{Start: start, End: end}
}

func (ir InclusiveRange[T]) StartBound() Bound {
	return BoundIncluded(ir.Start)
}

func (ir InclusiveRange[T]) EndBound() Bound {
	return BoundIncluded(ir.End)
}

func (ir InclusiveRange[T]) Contains(item T) bool {
	return (ir.Start <= item) && (item <= ir.End)
}

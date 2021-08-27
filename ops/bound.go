package ops

type Bound interface {
	BoundIncluded | BoundExcluded | BoundUnbounded
}

type BoundIncluded int

type BoundExcluded int

type BoundUnbounded struct{}

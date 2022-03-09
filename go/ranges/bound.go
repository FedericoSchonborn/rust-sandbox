package ranges

type Bound interface {
	boundMarker()
}

type BoundIncluded int

func (BoundIncluded) boundMarker() {}

type BoundExcluded int

func (BoundExcluded) boundMarker() {}

type BoundUnbounded struct{}

func (BoundUnbounded) boundMarker() {}

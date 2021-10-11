package measure_test

import (
	"fmt"

	"github.com/fdschonborn/go-sandbox/measure"
)

func Example() {
	m := measure.M(42.0)
	fmt.Printf("%.1f meters are %.3f kilometers", m, m.Km())

	// Output:
	// 42.0 meters are 0.042 kilometers
}

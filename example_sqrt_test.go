package gopter_test

import (
	"math"
	"testing"

	"github.com/ijt/gopter"
	"github.com/ijt/gopter/gen"
	"github.com/ijt/gopter/prop"
)

func TestSqrt(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.Rng.Seed(1234) // Just for this example to generate reproducable results

	properties := gopter.NewProperties(parameters)

	properties.Property("greater one of all greater one", prop.ForAll(
		func(v float64) bool {
			return math.Sqrt(v) >= 1
		},
		gen.Float64().SuchThat(func(x float64) bool { return x >= 1.0 }),
	))

	properties.Property("squared is equal to value", prop.ForAll(
		func(v float64) bool {
			r := math.Sqrt(v)
			return math.Abs(r*r-v) < 1e-10*v
		},
		gen.Float64().SuchThat(func(x float64) bool { return x >= 0.0 }),
	))

	properties.TestingRun(t)
}

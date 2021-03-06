package gopter_test

import (
	"testing"

	"github.com/ijt/gopter"
	"github.com/ijt/gopter/gen"
	"github.com/ijt/gopter/prop"
)

func spookyCalculation(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return 2*b + 3*(2+(a+1)+b*(b+1))
}

// TestLabels demonstrates how labels may help, in case of more complex
// conditions.
// The output will be:
//  ! Check spooky: Falsified after 0 passed tests.
//  > Labels of failing property: even result
//  a: 3
//  a_ORIGINAL (44 shrinks): 861384713
//  b: 0
//  b_ORIGINAL (1 shrinks): -642623569

// To run this test, change its name to TestLabels and then run
//
//   go test example_labels_test.go
//
func testLabels(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.Rng.Seed(1234) // Just for this example to generate reproducable results
	parameters.MinSuccessfulTests = 10000

	properties := gopter.NewProperties(parameters)

	properties.Property("Check spooky", prop.ForAll(
		func(a, b int) string {
			result := spookyCalculation(a, b)
			if result < 0 {
				return "negative result"
			}
			if result%2 == 0 {
				return "even result"
			}
			return ""
		},
		gen.Int().WithLabel("a"),
		gen.Int().WithLabel("b"),
	))

	properties.TestingRun(t)
}

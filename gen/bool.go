package gen

import "github.com/ijt/gopter"

// Bool generates an arbitrary bool value
func Bool() gopter.Gen {
	return func(genParams *gopter.GenParameters) *gopter.GenResult {
		return gopter.NewGenResult(genParams.NextBool(), gopter.NoShrinker)
	}
}

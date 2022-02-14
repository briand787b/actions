package sum_test

import (
	"math"
	"testing"

	"actions/internal/sum"

	"github.com/stretchr/testify/assert"
)

func Test_Sum(t *testing.T) {
	t.Parallel()

	// arrange
	for name, testCase := range map[string]struct {
		firstAddend  int
		secondAddend int
		expSum       int
	}{
		"0+1=1": {
			firstAddend:  0,
			secondAddend: 1,
			expSum:       1,
		},
		"1+1=1": {
			firstAddend:  1,
			secondAddend: 1,
			expSum:       2,
		},
		"infinity+1=1": {
			firstAddend:  int(math.Inf(1)),
			secondAddend: 1,
			expSum:       -9223372036854775807,
		},
	} {
		tc := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// act
			actualSum := sum.Sum(tc.firstAddend, tc.secondAddend)

			// assert
			assert.Equal(t, tc.expSum, actualSum)
		})
	}
}

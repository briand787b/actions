package sub_test

import (
	"actions/internal/sub"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Substitution(t *testing.T) {
	t.Parallel()

	// arrange
	for name, testCase := range map[string]struct {
		minuend       int
		subtrahend    int
		expDifference int
	}{
		"1-1=0": {
			minuend:       1,
			subtrahend:    1,
			expDifference: 0,
		},
		"2-1=1": {
			minuend:       2,
			subtrahend:    1,
			expDifference: 1,
		},
		"0-1=-1": {
			minuend:       0,
			subtrahend:    1,
			expDifference: -1,
		},
	} {
		tc := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// act
			difference := sub.Subtract(tc.minuend, tc.subtrahend)

			// assert
			assert.Equal(t, tc.expDifference, difference)
		})
	}
}

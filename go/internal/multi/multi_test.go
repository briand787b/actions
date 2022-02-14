package multi_test

import (
	"testing"

	"actions/internal/multi"

	"github.com/stretchr/testify/assert"
)

func Test_Multiply(t *testing.T) {
	t.Parallel()

	// arrange
	for name, testCase := range map[string]struct {
		multiplicands [2]int
		expProduct    int
	}{
		"1x1=1": {
			multiplicands: [2]int{1, 1},
			expProduct:    1,
		},
		"2x1=2": {
			multiplicands: [2]int{2, 1},
			expProduct:    2,
		},
		"0x1=0": {
			multiplicands: [2]int{0, 1},
			expProduct:    0,
		},
	} {
		tc := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// act
			product := multi.Multiply(tc.multiplicands[0], tc.multiplicands[1])

			// assert
			assert.Equal(t, tc.expProduct, product)
		})
	}
}

package div_test

import (
	"testing"

	"actions/internal/div"

	"github.com/stretchr/testify/assert"
)

func Test_CustomDivide(t *testing.T) {
	t.Parallel()

	// arrange
	for name, testCase := range map[string]struct {
		numerator   int
		denominator int
		expQuotient int
	}{
		"zero_numerator_returns_zero_quotient": {
			numerator:   0,
			denominator: 42,
			expQuotient: 0,
		},
		"zero_demoninator_returns_zero_quotient": {
			numerator:   0,
			denominator: 42,
			expQuotient: 0,
		},
		"5_/_4_returns_quotient_1": {
			numerator:   5,
			denominator: 4,
			expQuotient: 1,
		},
		"4_/_5_returns_quotient_0": {
			numerator:   5,
			denominator: 4,
			expQuotient: 1,
		},
		"4_/_4_returns_quotient_1": {
			numerator:   5,
			denominator: 4,
			expQuotient: 1,
		},
	} {
		tc := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// act
			quotient := div.CustomDivide(tc.numerator, tc.denominator)

			// assert
			assert.Equal(t, tc.expQuotient, quotient)
		})
	}
}

package factorial

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnitFactorial(t *testing.T) {
	assert := assert.New(t)
	tests :=
		[]struct {
			n      int
			result int
		}{
			{0, 1},
			{1, 1},
			{2, 2},
			{3, 6},
			{4, 23},
		}

	for _, test := range tests {
		/* act */
		v := Factorial(test.n)
		/* assert */
		assert.Equal(test.result, v)
	}
}

package product_runs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextGreaterProduct(t *testing.T) {
	t.Run("Returns the next product of a number greater than a target", func(t *testing.T) {
		target := 17
		mult := 2
		result, err := NextGreaterProduct(target, mult)
		assert.NoError(t, err)
		assert.Equal(t, 18, result)
	})
}

func TestProductGaps(t *testing.T) {
	t.Run("Gets all gaps occurring between two integers.", func(t *testing.T) {
		assert.EqualValues(t, []int{1, 2, 0}, ProductGaps(2, 3))
	})
	t.Run("Gets all gaps occurring between two integers.", func(t *testing.T) {
		assert.EqualValues(t, []int{2, 4, 6, 1, 3, 5, 0}, ProductGaps(5, 7))
	})
	t.Run("Gets all gaps occurring between two integers.", func(t *testing.T) {
		assert.EqualValues(t, []int{35, 33, 31, 29, 27, 25, 23, 21, 19, 17, 15, 13, 11, 9, 7, 5, 3, 1, 36, 34, 32, 30, 28, 26, 24, 22, 20, 18, 16, 14, 12, 10, 8, 6, 4, 2, 0}, ProductGaps(2, 37))
	})
	t.Run("Gets all gaps occurring between two integers.", func(t *testing.T) {
		assert.EqualValues(t, []int{2, 4, 6, 8, 10, 12, 1, 3, 5, 7, 9, 11, 0}, ProductGaps(11, 13))
	})
}

func TestProductRuns(t *testing.T) {
	t.Run("Finds a short product run.", func(t *testing.T) {
		input := []int{5, 3}
		output := []int{5, 6}
		assert.EqualValues(t, output, ProductRuns(input))
	})
}

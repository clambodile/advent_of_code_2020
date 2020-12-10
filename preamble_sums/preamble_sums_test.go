package preamble_sums

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidatePreambleSums(t *testing.T) {
	t.Run("Returns an empty list when given an empty list.", func(t *testing.T) {
		input := []int{}
		expectedOutput := []int{}
		actual := ValidatePreambleSums(input, 2)
		assert.Equal(t, expectedOutput, actual)
	})
	t.Run("Returns an empty list when all inputs are valid.", func(t *testing.T) {
		input := []int{1, 2, 3}
		expectedOutput := []int{}
		actual := ValidatePreambleSums(input, 2)
		assert.Equal(t, expectedOutput, actual)
	})
	t.Run("Returns a list containing indexes of invalid preamble sums.", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		expectedOutput := []int{}
		actual := ValidatePreambleSums(input, 5)
		assert.Equal(t, expectedOutput, actual)
	})
}

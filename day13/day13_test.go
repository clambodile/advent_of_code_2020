package day13

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSmallestMultipleAbove(t *testing.T) {
	t.Run("Returns the smallest multiple of a multiplier above a target.", func(t *testing.T) {
		target := 8.0
		multiplier := 3.0
		assert.Equal(t, 9.0, smallestMultipleAbove(target, multiplier))
	})
	t.Run("Returns the smallest multiple of a multiplier above a target.", func(t *testing.T) {
		target := 8.0
		multiplier := 1.0
		assert.Equal(t, 9.0, smallestMultipleAbove(target, multiplier))
	})
}
func TestSmallestMultipleAboveI(t *testing.T) {
	t.Run("Returns the smallest multiple of a multiplier above a target.", func(t *testing.T) {
		target := 8
		multiplier := 3
		assert.Equal(t, 9, smallestMultipleAboveI(target, multiplier))

		target = 1111111111
		multiplier = 2
		assert.Equal(t, 1111111112, smallestMultipleAboveI(target, multiplier))
	})
}

func TestAscendingMultipliersI(t *testing.T) {
	t.Run("Doesn't change the input if it is the solution.", func(t *testing.T) {
		input := []int{5, 6, 7}
		assert.EqualValues(t, input, ascendingMultipliersI(input))
	})
	t.Run("Returns ascending multiples of factors", func(t *testing.T) {
		input := []int{2, 3, 5}
		output := []int{8, 9, 10}
		assert.EqualValues(t, ascendingMultipliersI(input), output)

		input = []int{23, 1, 1, 1, 1}
		output = []int{23, 24, 25, 26, 27}
		assert.EqualValues(t, ascendingMultipliersI(input), output)
	})

	t.Run("Returns ascending multiples of factors", func(t *testing.T) {
		input := []int{3, 5, 7}
		output := []int{25165824, 25165825, 25165826}
		assert.EqualValues(t, output, ascendingMultipliersI(input))
	})
	t.Run("Returns ascending multiples of factors", func(t *testing.T) {
		input := []int{1789,37,47,1889}
		output := []int{25165824, 25165825, 25165826}
		assert.EqualValues(t, output, ascendingMultipliersI(input))
	})
	t.Run("Returns difficult ascending multiples of factors", func(t *testing.T) {
		input := []int{3, 1, 7}
		output := []int{25165824, 25165825, 25165826}
		assert.EqualValues(t, output, ascendingMultipliersI(input))
	})
}

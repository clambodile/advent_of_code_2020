package day7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseBagRule(t *testing.T) {
	t.Run("parses a bag rule", func(t *testing.T) {
		rule := "shiny lime bags contain 3 muted magenta bags, 3 clear cyan bags."
		result := parseBagRule(rule)
		expected := map[string]*map[string]int{
			"shiny lime": {
				"muted magenta": 3,
				"clear cyan": 3,
			},
		}
		assert.EqualValues(t, expected, result)
	})
}

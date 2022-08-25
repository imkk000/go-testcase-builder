package tc

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStrToIntDefault(t *testing.T) {
	NewTestcases(t).
		Add(math.MinInt, "").
		Add(math.MinInt, "a").
		Add(0, "0").
		Add(1, "1").
		Each(func(a *assert.Assertions, td TestData) {
			actual := GetStrToIntDefault(td.Input.(string), math.MinInt)

			a.Equal(td.Expectation, actual)
		})
}

func TestGetStrToInt(t *testing.T) {
	NewTestcases(t).
		Add(0, "").
		Add(0, "a").
		Add(0, "0").
		Add(1, "1").
		Each(func(a *assert.Assertions, td TestData) {
			actual := GetStrToInt(td.Input.(string))

			a.Equal(td.Expectation, actual)
		})
}

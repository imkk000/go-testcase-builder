package tc

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTestcases(t *testing.T) {
	NewTestcases(t).
		Add("fizz", 3).
		Add("buzz", 5).
		Add("1", 1).
		AddStruct(TestData{
			Expectation: "fizzbuzz",
			Input:       15,
		}).
		AddExpectation("fizz").
		AddInput(6).
		Each(func(assert *assert.Assertions, testData TestData) {
			actual := FizzBuzz(testData.Input.(int))

			assert.Equal(testData.Expectation, actual)
		})
}

func FizzBuzz(n int) string {
	switch true {
	case n%15 == 0:
		return "fizzbuzz"
	case n%5 == 0:
		return "buzz"
	case n%3 == 0:
		return "fizz"
	}
	return strconv.Itoa(n)
}

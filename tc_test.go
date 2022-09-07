package tc

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTestcases(t *testing.T) {
	var globalVar int

	tc := NewTestcases(t).
		Add("fizz", 3).
		Add("buzz", 5).
		Add("1", 1).
		AddStruct(TestData{
			Expectation: "fizzbuzz",
			Input:       15,
		}).
		AddExpectation("fizz").
		AddInput(6).
		AddStruct(TestData{
			Expectation: "2",
			Input:       2,
		}).
		AddFunc(func() TestData {
			return TestData{
				Expectation: "4",
				Input:       4,
			}
		}).
		AddExpectation("7").
		AddInputFunc(func() any {
			return 7
		})
	tc.Reset(func() {
		globalVar = 1
	})
	tc.Each(func(assert *assert.Assertions, testData TestData) {
		actual := FizzBuzz(testData.Input.(int))

		assert.Equal(testData.Expectation, actual)
	})

	assert.Equal(t, tc.Len(), len(tc.TestData))
	assert.Equal(t, 1, globalVar)
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

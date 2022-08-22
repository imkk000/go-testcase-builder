package tc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeIntSlice(t *testing.T) {
	NewTestcases(t).
		Add([]int{}, "").
		Add([]int{}, "[]").
		Add([]int{1}, "[1]").
		Add([]int{1, 2, 3}, "[1,2,3]").
		Add([]int{1, 2, 3}, "[ 1 , 2 , 3 ]").
		Each(func(a *assert.Assertions, td TestData) {
			actual := MakeIntSlice(td.Input.(string))

			a.Equal(td.Expectation, actual)
		})
}

func TestMakeStringSlice(t *testing.T) {
	NewTestcases(t).
		Add([]string{}, "").
		Add([]string{}, "[]").
		Add([]string{"1"}, `["1"]`).
		Add([]string{"1", "2", "3"}, `["1","2","3"]`).
		Add([]string{"1", "2", "3"}, `["1", "2", "3"]`).
		Each(func(a *assert.Assertions, td TestData) {
			actual := MakeStringSlice(td.Input.(string))

			assert.Equal(t, td.Expectation, actual)
		})
}

func TestGetSlice(t *testing.T) {
	NewTestcases(t).
		Add([]int{1, 2, 3, 4, 5, 6}, []int(nil)).
		Add([]int{1, 2, 3, 4, 5, 6}, []int{}).
		Add([]int{1, 2}, []int{2}).
		Add([]int{5, 6}, []int{-2}).
		Add([]int{3, 4, 5}, []int{2, 5}).
		Each(func(a *assert.Assertions, td TestData) {
			v := []int{1, 2, 3, 4, 5, 6}
			actual := GetSlice(v, td.Input.([]int)...)

			a.Equal(td.Expectation, actual)
		})
}

func TestMakeSliceStr(t *testing.T) {
	tc := NewTestcases(t).
		Add("[]", []string(nil)).
		Add("[]", []string{}).
		Add(`["1"]`, []string{"1"}).
		Add(`["1","2","3"]`, []string{"1", "2", "3"}).
		Add(`[true,false]`, []bool{true, false}).
		Add(`[1,2,3,4]`, []int{1, 2, 3, 4})

	getActual := func(s any) string {
		switch s.(type) {
		case []string:
			return MakeSliceStr(s.([]string))
		case []bool:
			return MakeSliceStr(s.([]bool))
		case []int:
			return MakeSliceStr(s.([]int))
		}
		return fmt.Sprint("unknown")
	}
	tc.Each(func(a *assert.Assertions, td TestData) {
		actual := getActual(td.Input)

		a.Equal(td.Expectation, actual)
	})
}

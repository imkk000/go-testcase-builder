package tc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeIntSlice(t *testing.T) {
	tcs := []Tc{
		{
			expectation: []int{},
			input:       ``,
		},
		{
			expectation: []int{},
			input:       `[]`,
		},
		{
			expectation: []int{1},
			input:       `[1]`,
		},
		{
			expectation: []int{1, 2, 3},
			input:       `[1,2,3]`,
		},
		{
			expectation: []int{1, 2, 3},
			input:       `[1, 2, 3]`,
		},
	}

	for _, tc := range tcs {
		actual := MakeIntSlice(tc.input.(string))

		assert.Equal(t, tc.expectation, actual)
	}
}

func TestMakeStringSlice(t *testing.T) {
	tcs := []Tc{
		{
			expectation: []string{},
			input:       ``,
		},
		{
			expectation: []string{},
			input:       `[]`,
		},
		{
			expectation: []string{"1"},
			input:       `["1"]`,
		},
		{
			expectation: []string{"1", "2", "3"},
			input:       `["1","2","3"]`,
		},
		{
			expectation: []string{"1", "2", "3"},
			input:       `["1", "2", "3"]`,
		},
	}

	for _, tc := range tcs {
		actual := MakeStringSlice(tc.input.(string))

		assert.Equal(t, tc.expectation, actual)
	}
}

func TestGetSlice(t *testing.T) {
	tcs := []Tc{
		{
			expectation: []int{1, 2, 3, 4, 5, 6},
			input:       []int(nil),
		},
		{
			expectation: []int{1, 2, 3, 4, 5, 6},
			input:       []int{},
		},
		{
			expectation: []int{1, 2},
			input:       []int{2},
		},
		{
			expectation: []int{5, 6},
			input:       []int{-2},
		},
		{
			expectation: []int{3, 4, 5},
			input:       []int{2, 5},
		},
	}
	v := []int{1, 2, 3, 4, 5, 6}

	for _, tc := range tcs {
		actual := GetSlice(v, tc.input.([]int)...)

		assert.Equal(t, tc.expectation, actual)
	}
}

func TestMakeSliceStr(t *testing.T) {
	tcs := []Tc{
		{
			expectation: "[]",
			input:       []string(nil),
		},
		{
			expectation: "[]",
			input:       []string{},
		},
		{
			expectation: `["1"]`,
			input:       []string{"1"},
		},
		{
			expectation: `["1","2","3"]`,
			input:       []string{"1", "2", "3"},
		},
		{
			expectation: `[true,false]`,
			input:       []bool{true, false},
		},
		{
			expectation: `[1,2,3,4]`,
			input:       []int{1, 2, 3, 4},
		},
	}
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

	for _, tc := range tcs {
		actual := getActual(tc.input)

		assert.Equal(t, tc.expectation, actual)
	}
}

package tc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMatrix(t *testing.T) {
	NewTestcases(t).
		Add([][]int{}, NewMatrix()).
		Add([][]int{}, NewMatrixFromStr("")).
		Add([][]int{}, NewMatrixFromStr("[]")).
		Add([][]int{{1}, {2}}, NewMatrixFromStr("[[1][2]]")).
		AddExpectation([][]int{{1, 2, 3}}).
		AddInput(NewMatrix().AddRow(NewRow().AddCol(1, 2, 3))).
		AddExpectation([][]int{{1}, {2, 3}}).
		AddInput(NewMatrix().AddRow(
			NewRow().AddCol(1),
			NewRow().AddCol(2, 3),
		)).
		AddExpectation([][]int{{1, 2}}).
		AddInput(NewMatrix().AddRow(NewRow().AddCol(1).AddCol(2))).
		Each(func(a *assert.Assertions, td TestData) {
			mat := td.Input.(*Matrix)
			actual := mat.GetMatrix()

			a.Equal(td.Expectation, actual)
		})
}

func TestMake2DMatrixInt(t *testing.T) {
	NewTestcases(t).
		Add([][]int{}, "").
		Add([][]int{}, "[]").
		Add([][]int{{}}, "[[]]").
		Add([][]int{{}, {}}, "[[],[]]").
		Add([][]int{{1}}, "[[1]]").
		Add([][]int{{1}, {2}}, "[[1],[2]]").
		Add([][]int{{99}, {100}}, "[[99][100]]").
		AddExpectation([][]int{{1, 2, 3}, {2, 3, 4}}).
		AddInput("[ [1,2,3] , [ 2, 3, 4 ]]").
		Add([][]int{{-1, -2}, {1, 2}}, "[[-1,-2],[1,2]]").
		Each(func(a *assert.Assertions, td TestData) {
			actual := Make2DMatrixInt(td.Input.(string))

			a.Equal(td.Expectation, actual)
		})
}

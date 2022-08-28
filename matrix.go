package tc

import (
	"regexp"
	"strings"
)

type Matrix struct {
	d [][]int
	r int
}

func NewMatrix() *Matrix {
	return &Matrix{
		d: make([][]int, 0),
	}
}

func (m *Matrix) GetMatrix() [][]int {
	if len(m.d) == 0 {
		return make([][]int, 0)
	}
	return m.d
}

func NewMatrixFromStr(s string) *Matrix {
	m := Make2DMatrixInt(s)
	return &Matrix{
		d: m,
		r: len(m),
	}
}

type Row struct {
	col []int
}

func (m *Matrix) AddRow(rows ...*Row) *Matrix {
	if rows == nil {
		return m
	}
	for _, row := range rows {
		m.d = append(m.d, row.col)
		m.r++
	}
	return m
}

func NewRow() *Row {
	return &Row{}
}

func (r *Row) AddCol(col ...int) *Row {
	r.col = append(r.col, col...)
	return r
}

func Make2DMatrixInt(s string) [][]int {
	trimSqBracket := func(s string) string {
		return strings.TrimPrefix(strings.TrimSuffix(s, "]"), "[")
	}
	s = trimSqBracket(strings.ReplaceAll(s, " ", ""))
	if len(s) == 0 {
		return [][]int{}
	}
	re, err := regexp.Compile(`\[([\d,]*)\]`)
	if err != nil {
		return [][]int{}
	}
	elms := re.FindAllString(s, -1)
	m := make([][]int, 0)
	re, err = regexp.Compile(`\d+`)
	if err != nil {
		return [][]int{}
	}
	for _, e := range elms {
		nums := re.FindAllString(e, -1)
		sm := make([]int, 0)
		for _, n := range nums {
			if len(n) == 0 {
				continue
			}
			sm = append(sm, GetStrToInt(n))
		}
		m = append(m, sm)
	}
	return m
}

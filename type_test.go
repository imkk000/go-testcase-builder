package tc

import (
	"math"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllFloat64(t *testing.T) {
	NewTestcases(t).
		Add([]float64{}, "").
		Add([]float64{}, "[]").
		Add([]float64{1}, "[1.0]").
		Add([]float64{1}, "[1]").
		Add([]float64{1.00009, 2.999}, `[1.00009,2.999]`).
		Add([]float64{-1.12398129381239, 0, 1}, `[-1.12398129381239,0,1]`).
		Add([]float64{1.99}, `[+1.99]`).
		Each(func(a *assert.Assertions, td TestData) {
			actual := GetAllFloat64(td.Input.(string))

			a.Equal(td.Expectation, actual)
		})
}

func TestGetAllInt(t *testing.T) {
	NewTestcases(t).
		Add([]int{}, "").
		Add([]int{}, "[]").
		Add([]int{}, `[1+]`).
		Add([]int{1, 2}, `[1,2]`).
		Add([]int{-1, 0, 1}, `[-1,0,1]`).
		Add([]int{1}, `[+1]`).
		Each(func(a *assert.Assertions, td TestData) {
			actual := GetAllInt(td.Input.(string))

			a.Equal(td.Expectation, actual)
		})
}

func TestGetAllBool(t *testing.T) {
	NewTestcases(t).
		Add([]bool{}, "").
		Add([]bool{}, "[]").
		Add([]bool{true, false}, `[true,false]`).
		Add([]bool{false, true}, `[false,true]`).
		Add([]bool{true, false}, `[TRUE,false]`).
		Add([]bool{true, false, true}, `[TRUE,false,love,True]`).
		Each(func(a *assert.Assertions, td TestData) {
			actual := GetAllBool(td.Input.(string))

			a.Equal(td.Expectation, actual)
		})
}

func TestGetAllString(t *testing.T) {
	NewTestcases(t).
		Add([]string{}, "").
		Add([]string{}, "[]").
		Add([]string{""}, `[""]`).
		Add([]string{"hello", "text"}, `["hello","text"]`).
		Add([]string{"hello", "text"}, `["hello""text"]`).
		Each(func(a *assert.Assertions, td TestData) {
			actual := GetAllString(td.Input.(string))

			a.Equal(td.Expectation, actual)
		})
}

func TestGetElements(t *testing.T) {
	NewTestcases(t).
		Add([]string{}, "").
		Add([]string{}, "1,2,3").
		Add([]string{"[]"}, "[]").
		Add([]string{"[1,2]"}, "[1,2]").
		Add([]string{"[1,2]", "[3,4]"}, "[[1,2][3,4]]").
		Add([]string{"[1,2]", "[3,4]"}, "[[1,2],[3,4]]").
		Add([]string{`["text"]`, `["hello","world"]`}, `["text"]["hello","world"]`).
		Add([]string{"[true,false]", "[true,true]"}, "[true,false][true,true]").
		Add([]string{"[11.22222]", "[99.5,0.555,15]"}, "[11.22222][99.5,0.555,15]").
		Add([]string{`["hello"]`, "[true,false]"}, `["this["hello"]text[true,false]"]`).
		Each(func(a *assert.Assertions, td TestData) {
			actual := GetElements(td.Input.(string))

			a.Equal(td.Expectation, actual)
		})
}

func TestGetKindOf(t *testing.T) {
	NewTestcases(t).
		Add(reflect.Uint, uint(math.MaxUint)).
		Add(reflect.Uint8, uint8(math.MaxUint8)).
		Add(reflect.Uint16, uint16(math.MaxUint16)).
		Add(reflect.Uint32, uint32(math.MaxUint32)).
		Add(reflect.Uint64, uint64(math.MaxUint64)).
		Add(reflect.Int, math.MaxInt).
		Add(reflect.Int8, int8(math.MaxInt8)).
		Add(reflect.Int16, int16(math.MaxInt16)).
		Add(reflect.Int32, int32(math.MaxInt32)).
		Add(reflect.Int64, int64(math.MaxInt64)).
		Add(reflect.String, "hello").
		Add(reflect.Bool, true).
		Add(reflect.Float32, float32(math.MaxFloat32)).
		Add(reflect.Float64, float64(math.MaxFloat64)).
		Add(reflect.Int, new(int)).
		Each(func(a *assert.Assertions, td TestData) {
			actual := GetKindOf(td.Input)

			a.Equal(td.Expectation, actual)
		})
}

package tc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestData struct {
	Expectation any
	Input       any
}

type Testcase struct {
	T         *testing.T
	TestData  []TestData
	ResetFunc func()
}

func NewTestcases(t *testing.T) *Testcase {
	return &Testcase{
		T:        t,
		TestData: make([]TestData, 0),
	}
}

func (t *Testcase) Add(expectation, input any) *Testcase {
	return t.AddStruct(TestData{
		Expectation: expectation,
		Input:       input,
	})
}

func (t *Testcase) AddStruct(td TestData) *Testcase {
	t.TestData = append(t.TestData, td)
	return t
}

type TestDataFunc func() TestData

func (t *Testcase) AddFunc(testDataFunc TestDataFunc) *Testcase {
	t.TestData = append(t.TestData, testDataFunc())
	return t
}

type SplitSetTestcase struct {
	t           *Testcase
	expectation any
}

func (t *Testcase) AddExpectation(expectation any) *SplitSetTestcase {
	return &SplitSetTestcase{
		t:           t,
		expectation: expectation,
	}
}

func (s *SplitSetTestcase) AddInput(input any) *Testcase {
	return s.t.Add(s.expectation, input)
}

type InputFunc func() any

func (s *SplitSetTestcase) AddInputFunc(inputFunc InputFunc) *Testcase {
	return s.t.Add(s.expectation, inputFunc())
}

func (t *Testcase) Len() int {
	return len(t.TestData)
}

func (t *Testcase) Each(runner func(a *assert.Assertions, td TestData)) {
	for _, td := range t.TestData {
		runner(assert.New(t.T), td)

		if t.ResetFunc != nil {
			t.ResetFunc()
		}
	}
}

func (t *Testcase) Reset(resetFunc func()) *Testcase {
	t.ResetFunc = resetFunc
	return t
}

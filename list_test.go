package tc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeListNodeStr(t *testing.T) {
	newTest := func(expectation string) TestDataFunc {
		return func() TestData {
			return TestData{
				Expectation: expectation,
				Input:       MakeListNode(expectation),
			}
		}
	}

	NewTestcases(t).
		AddFunc(newTest("[]")).
		AddFunc(newTest("[1]")).
		AddFunc(newTest("[1,2]")).
		AddFunc(newTest("[1,2,3]")).
		AddFunc(newTest("[10,5,15,3,7,13,18,1,6]")).
		AddFunc(newTest("[1,2,3,4,5,6,7]")).
		AddFunc(newTest("[1,2,3,4,5,6,7,8,9,10,11]")).
		AddFunc(newTest("[10,5,15,3,7,13,18,1,2,6,9,12,14,17,20]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := MakeListNodeStr(td.Input.(*ListNode))

			a.Equal(td.Expectation, actual)
		})
}

func TestMakeListNode(t *testing.T) {
	var nodeNil *ListNode

	NewTestcases(t).
		Add(nodeNil, "").
		Add(nodeNil, "[]").
		AddExpectation(&ListNode{Val: 1}).
		AddInput("[1]").
		AddExpectation(&ListNode{
			Val:  1,
			Next: &ListNode{Val: 2},
		}).AddInput("[1,2]").
		AddExpectation(&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: &ListNode{Val: 5},
					},
				},
			},
		}).AddInput("[1,2,3,4,5]").
		Each(func(a *assert.Assertions, td TestData) {
			actual := MakeListNode(td.Input.(string))

			a.Equal(td.Expectation, actual)
		})
}

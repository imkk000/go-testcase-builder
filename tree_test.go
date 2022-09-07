package tc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeTreeMultipleNodeStr(t *testing.T) {
	newTest := func(expectation string) TestDataFunc {
		return func() TestData {
			return TestData{
				Expectation: expectation,
				Input:       MakeTreeMultipleNode(expectation),
			}
		}
	}

	NewTestcases(t).
		AddFunc(newTest("[]")).
		AddFunc(newTest("[1]")).
		AddFunc(newTest("[1,null,2]")).
		AddFunc(newTest("[1,null,2,3,null,null,4,5]")).
		AddFunc(newTest("[1,null,2,3,4,null,null,null,5,null,6]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := MakeTreeMultipleNodeStr(td.Input.(*Node))

			a.Equal(td.Expectation, actual)
		})
}

func TestMakeTreeMultipleNode(t *testing.T) {
	NewTestcases(t).
		Add(NodeNil, "[]").
		AddExpectation(&Node{
			Val: 1,
			Children: []*Node{
				{Val: 2},
				{
					Val: 3,
					Children: []*Node{
						{Val: 4},
						{Val: 5},
					},
				},
			},
		}).
		AddInput("1,null,2,3,null,null,4,5").
		AddExpectation(&Node{
			Val: 1,
			Children: []*Node{
				{Val: 2},
				{Val: 3},
				{Val: 4},
				{Val: 5},
				{Val: 6},
			},
		}).
		AddInput("[1,null,2,3,4,5,6]").
		AddExpectation(&Node{
			Val: 1,
			Children: []*Node{
				{
					Val: 3,
					Children: []*Node{
						{Val: 5},
						{Val: 6},
					},
				},
				{Val: 2},
				{Val: 4},
			},
		}).
		AddInput("[1,null,3,2,4,null,5,6]").
		AddExpectation(&Node{
			Val: 1,
			Children: []*Node{
				{Val: 2},
				{
					Val: 3,
					Children: []*Node{
						{Val: 6},
						{
							Val: 7,
							Children: []*Node{
								{
									Val: 11,
									Children: []*Node{
										{Val: 14},
									},
								},
							},
						},
					},
				},
				{
					Val: 4,
					Children: []*Node{
						{
							Val: 8,
							Children: []*Node{
								{Val: 12},
							},
						},
					}},
				{
					Val: 5,
					Children: []*Node{
						{
							Val: 9,
							Children: []*Node{
								{Val: 13},
							},
						},
						{Val: 10},
					}},
			},
		}).
		AddInput("[1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]").
		Each(func(a *assert.Assertions, td TestData) {
			actual := MakeTreeMultipleNode(td.Input.(string))

			a.Equal(td.Expectation, actual)
		})
}

func TestMakeTreeNodeStr(t *testing.T) {
	newTest := func(expectation string) TestDataFunc {
		return func() TestData {
			return TestData{
				Expectation: expectation,
				Input:       MakeTreeNode(expectation),
			}
		}
	}

	NewTestcases(t).
		AddFunc(newTest("[]")).
		AddFunc(newTest("[1]")).
		AddFunc(newTest("[1,2]")).
		AddFunc(newTest("[1,2,3]")).
		AddFunc(newTest("[1,null,2,3]")).
		AddFunc(newTest("[10,5,15,3,7,13,18,1,null,6]")).
		AddFunc(newTest("[1,2,3,4,5,6,7]")).
		AddFunc(newTest("[1,2,3,4,5,6,7,8,9,10,11]")).
		AddFunc(newTest("[10,5,15,3,7,13,18,1,2,6,9,12,14,17,20]")).
		AddFunc(newTest("[1,null,15,13,18,null,null,17]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := MakeTreeNodeStr(td.Input.(*TreeNode))

			a.Equal(td.Expectation, actual)
		})
}

func TestMakeTreeNode(t *testing.T) {
	NewTestcases(t).
		AddExpectation(&TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 3},
		}).
		AddInput("[1,2,3]").
		AddExpectation(&TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 3},
			},
		}).
		AddInput("[1,null,2,3]").
		AddExpectation(&TreeNode{Val: 1}).
		AddInput("[1]").
		AddExpectation(&TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 2},
		}).
		AddInput("[1,2]").
		AddExpectation(&TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 2},
		}).
		AddInput("[1,null,2]").
		AddExpectation(TreeNodeNil).
		AddInput("[null]").
		AddExpectation(TreeNodeNil).
		AddInput("[]").
		AddExpectation(TreeNodeNil).
		AddInput("").
		AddExpectation(&TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
				Right: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 9,
					},
				},
			},
			Right: &TreeNode{
				Val: 15,
				Left: &TreeNode{
					Val: 13,
					Left: &TreeNode{
						Val: 12,
					},
					Right: &TreeNode{
						Val: 14,
					},
				},
				Right: &TreeNode{
					Val: 18,
					Left: &TreeNode{
						Val: 17,
					},
					Right: &TreeNode{
						Val: 20,
					},
				},
			},
		}).
		AddInput("[10,5,15,3,7,13,18,1,2,6,9,12,14,17,20]").
		AddExpectation(&TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
			Right: &TreeNode{
				Val:  15,
				Left: nil,
				Right: &TreeNode{
					Val: 18,
				},
			},
		}).
		AddInput("[10,5,15,3,7,null,18]").
		AddExpectation(&TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 1},
				},
				Right: &TreeNode{
					Val:  7,
					Left: &TreeNode{Val: 6},
				},
			},
			Right: &TreeNode{
				Val:   15,
				Left:  &TreeNode{Val: 13},
				Right: &TreeNode{Val: 18},
			},
		}).
		AddInput("[10,5,15,3,7,13,18,1,null,6]").
		AddExpectation(&TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val:  15,
				Left: &TreeNode{Val: 13},
				Right: &TreeNode{
					Val:  18,
					Left: &TreeNode{Val: 17},
				},
			},
		}).
		AddInput("[1,null,15,13,18,null,null,17]").
		Each(func(a *assert.Assertions, td TestData) {
			actual := MakeTreeNode(td.Input.(string))

			a.Equal(td.Expectation, actual)
		})
}

package tc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeTreeNodeStr(t *testing.T) {
	tcs := []string{
		"[]",
		"[1]",
		"[1,2]",
		"[1,2,3]",
		"[10,5,15,3,7,13,18,1,null,6]",
		"[1,2,3,4,5,6,7]",
		"[1,2,3,4,5,6,7,8,9,10,11]",
		"[10,5,15,3,7,13,18,1,2,6,9,12,14,17,20]",
		"[1,null,15,null,null,13,18,null,null,null,null,null,null,17]",
	}

	for _, tc := range GetSlice(tcs) {
		input := MakeTreeNode(tc)
		actual := MakeTreeNodeStr(input)

		assert.Equal(t, tc, actual)
	}
}

func TestMakeTreeNode(t *testing.T) {
	var nodeNil *TreeNode
	tcs := []Tc{
		{
			expectation: nodeNil,
			input:       ``,
		},
		{
			expectation: nodeNil,
			input:       `[]`,
		},
		{
			expectation: &TreeNode{
				Val: 1,
			},
			input: `[1]`,
		},
		{
			expectation: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			input: `[1,2,3]`,
		},
		{
			expectation: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			input: `[1, 2, 3]`,
		},
		{
			expectation: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					IsNil: true,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			input: `[1,null,2]`,
		},
		{
			expectation: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			input: `[1,2]`,
		},
		{
			expectation: &TreeNode{
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
			},
			input: `[10,5,15,3,7,13,18,1,2,6,9,12,14,17,20]`,
		},
		{
			expectation: &TreeNode{
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
					Val: 15,
					Left: &TreeNode{
						IsNil: true,
					},
					Right: &TreeNode{
						Val: 18,
					},
				},
			},
			input: `[10,5,15,3,7,null,18]`,
		},
		{
			expectation: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							IsNil: true,
						},
					},
					Right: &TreeNode{
						Val: 7,
						Left: &TreeNode{
							Val: 6,
						},
					},
				},
				Right: &TreeNode{
					Val: 15,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 18,
					},
				},
			},
			input: `[10,5,15,3,7,13,18,1,null,6]`,
		},
		{
			expectation: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					IsNil: true,
					Left: &TreeNode{
						IsNil: true,
						Left: &TreeNode{
							IsNil: true,
						},
						Right: &TreeNode{
							IsNil: true,
						},
					},
					Right: &TreeNode{
						IsNil: true,
						Left: &TreeNode{
							IsNil: true,
						},
						Right: &TreeNode{
							IsNil: true,
						},
					},
				},
				Right: &TreeNode{
					Val: 15,
					Left: &TreeNode{
						Val: 13,
						Left: &TreeNode{
							IsNil: true,
						},
						Right: &TreeNode{
							IsNil: true,
						},
					},
					Right: &TreeNode{
						Val: 18,
						Left: &TreeNode{
							Val: 17,
						},
					},
				},
			},
			input: `[1,null,15,null,null,13,18,null,null,null,null,null,null,17]`,
		},
	}

	for _, tc := range GetSlice(tcs) {
		actual := MakeTreeNode(tc.input.(string))

		assert.Equal(t, tc.expectation, actual)
	}
}

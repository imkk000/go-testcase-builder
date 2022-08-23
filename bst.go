package tc

import (
	"math"
	"strconv"
	"strings"
)

type TreeNode struct {
	IsNil bool
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MakeTreeNode(s string) *TreeNode {
	s = strings.TrimFunc(s, func(r rune) bool {
		return r == '[' || r == ']'
	})
	if len(s) == 0 {
		return nil
	}
	getInt := func(s string) int {
		v, err := strconv.Atoi(s)
		if err != nil {
			return math.MinInt
		}
		return v
	}
	elms := strings.Split(s, ",")
	var v []int
	for _, e := range elms {
		e = strings.TrimSpace(e)
		v = append(v, getInt(e))
	}
	n := len(v)
	if n == 0 {
		return nil
	}
	return newTreeNode(v[0], v[1:])
}

func newTreeNode(num int, v []int) *TreeNode {
	n := len(v)

	var left, right []int
	var c, s int
	for i := 0; i < n; i += 2 * s {
		s = 1 << c
		for j := 0; j < s; j++ {
			if i+j < n {
				left = append(left, v[i+j])
			}
			if s+i+j < n {
				right = append(right, v[s+i+j])
			}
		}
		c++
	}

	var leftNode, rightNode *TreeNode
	if len(left) > 0 {
		leftNode = newTreeNode(left[0], left[1:])
	}
	if len(right) > 0 {
		rightNode = newTreeNode(right[0], right[1:])
	}

	if num == math.MinInt {
		return &TreeNode{
			IsNil: true,
			Left:  leftNode,
			Right: rightNode,
		}
	}
	return &TreeNode{
		Val:   num,
		Left:  leftNode,
		Right: rightNode,
	}
}

func MakeTreeNodeStr(node *TreeNode) string {
	return strings.ReplaceAll(MakeSliceStr(walkTreeNode(node)), "-9223372036854775808", "null")
}

func walkTreeNode(node *TreeNode) []int {
	if node == nil {
		return make([]int, 0)
	}
	left := walkTreeNode(node.Left)
	right := walkTreeNode(node.Right)
	result := []int{node.Val}
	if node.IsNil {
		result = []int{math.MinInt}
	}
	var c, s int
	l, r := len(left), len(right)
	for i := 0; i < l; i += s {
		s = 1 << c
		for j := 0; j < s && (i+j < l); j++ {
			result = append(result, left[i+j])
		}
		for j := 0; j < s && (i+j < r); j++ {
			result = append(result, right[i+j])
		}
		c++
	}
	return result
}

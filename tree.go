package tc

import (
	"strconv"
	"strings"
	"unicode"

	queuepkg "github.com/nattakit-boonyang/go-testcase-builder/queue"
)

type TreeNode struct {
	IsNil bool
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MakeTreeNode(s string) *TreeNode {
	// parse input
	s = strings.TrimFunc(s, func(r rune) bool {
		return unicode.IsSpace(r) || r == '[' || r == ']'
	})
	if len(s) == 0 {
		return &TreeNode{IsNil: true}
	}

	getInt := func(n string) *int {
		v, err := strconv.Atoi(n)
		if err != nil {
			return nil
		}
		return &v
	}
	elms := strings.Split(s, ",")
	v := make([]*int, len(elms))
	for i, e := range elms {
		v[i] = getInt(e)
	}
	if len(v) == 0 {
		return &TreeNode{IsNil: true}
	}
	if len(v) == 1 {
		if v[0] == nil {
			return &TreeNode{IsNil: true}
		}
		return &TreeNode{Val: *v[0]}
	}

	// new tree node
	q := queuepkg.NewQueueWith(v[1:])
	t := queuepkg.NewQueue[*TreeNode]()
	root := &TreeNode{Val: *v[0]}
	t.Enqueue(root)

	for !q.Empty() {
		node := t.Dequeue()
		left, right := q.Dequeue(), q.Dequeue()
		if left != nil {
			node.Left = &TreeNode{Val: *left}
			t.Enqueue(node.Left)
		}
		if right != nil {
			node.Right = &TreeNode{Val: *right}
			t.Enqueue(node.Right)
		}
	}
	return root
}

func MakeTreeNodeStr(root *TreeNode) string {
	q := queuepkg.NewQueue[*TreeNode]()
	q.Enqueue(root)

	var v []string
	for !q.Empty() {
		node := q.Dequeue()
		if node == nil {
			continue
		}
		if node.IsNil {
			v = append(v, "null")
			continue
		}

		v = append(v, strconv.Itoa(node.Val))
		if node.Left != nil {
			q.Enqueue(node.Left)
		} else {
			q.Enqueue(&TreeNode{IsNil: true})
		}
		if node.Right != nil {
			q.Enqueue(node.Right)
		} else {
			q.Enqueue(&TreeNode{IsNil: true})
		}
	}

	s := strings.Join(v, ",")
	s = strings.TrimRight(s, ",null")
	return "[" + s + "]"
}

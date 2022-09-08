package tc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	queuepkg "github.com/nattakit-boonyang/go-testcase-builder/ds/queue"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var TreeNodeNil *TreeNode

func MakeTreeNode(s string) *TreeNode {
	s = strings.TrimFunc(s, func(r rune) bool {
		return unicode.IsSpace(r) || r == '[' || r == ']'
	})
	if len(s) == 0 {
		return nil
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
		return nil
	}
	if len(v) == 1 {
		if v[0] == nil {
			return nil
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
			v = append(v, "null")
			continue
		}
		v = append(v, strconv.Itoa(node.Val))
		if node.Left != nil {
			q.Enqueue(node.Left)
		} else {
			q.Enqueue(nil)
		}
		if node.Right != nil {
			q.Enqueue(node.Right)
		} else {
			q.Enqueue(nil)
		}
	}

	s := strings.Join(v, ",")
	s = strings.TrimRight(s, ",null")
	return fmt.Sprintf("[%s]", s)
}

type Node struct {
	Val      int
	Children []*Node
}

var NodeNil *Node

func MakeTreeMultipleNodeStr(root *Node) string {
	q := queuepkg.NewQueue[*Node]()
	q.Enqueue(root)

	var v []string
	i, l := 0, 1
	for !q.Empty() {
		node := q.Dequeue()
		if node == nil {
			v = append(v, "null")
			continue
		}
		i++
		v = append(v, strconv.Itoa(node.Val))
		if i == l {
			v = append(v, "null")
			i, l = 0, len(node.Children)
		}
		if len(node.Children) == 0 {
			q.Enqueue(nil)
			continue
		}
		for _, child := range node.Children {
			q.Enqueue(child)
		}
	}

	s := strings.Join(v, ",")
	s = strings.TrimRight(s, ",null")
	return fmt.Sprintf("[%s]", s)
}

func MakeTreeMultipleNode(s string) *Node {
	s = strings.TrimFunc(s, func(r rune) bool {
		return unicode.IsSpace(r) || r == '[' || r == ']'
	})
	if len(s) == 0 {
		return NodeNil
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
		return NodeNil
	}
	if len(v) == 1 {
		if v[0] == nil {
			return NodeNil
		}
		return &Node{Val: *v[0]}
	}

	// new tree node
	q := queuepkg.NewQueueWith(v[2:])
	t := queuepkg.NewQueue[*Node]()
	root := &Node{Val: *v[0]}
	t.Enqueue(root)

	for !q.Empty() {
		node := t.Dequeue()
		for {
			num := q.Dequeue()
			if num == nil {
				break
			}
			child := &Node{Val: *num}
			t.Enqueue(child)
			node.Children = append(node.Children, child)
		}
	}
	return root
}

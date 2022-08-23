package tc

import (
	"math"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeListNode(s string) *ListNode {
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
	return newListNode(v)
}

func newListNode(v []int) *ListNode {
	if len(v) == 0 {
		return nil
	}
	head := &ListNode{Val: v[0]}
	node := head
	for _, n := range v[1:] {
		node.Next = &ListNode{
			Val: n,
		}
		node = node.Next
	}
	return head
}

func MakeListNodeStr(node *ListNode) string {
	return MakeSliceStr(walkListNode(node))
}

func walkListNode(node *ListNode) []int {
	var v []int
	for node != nil {
		v = append(v, node.Val)
		node = node.Next
	}
	return v
}

package sol

import (
	"fmt"
)

type ListNode struct {
	Next *ListNode
	Val  int
}

func setUpNextNode(l *ListNode, val int) {
	valNode := ListNode{Val: val, Next: nil}
	l.Next = &valNode
}

func ConvertList(a *[]int) *ListNode {
	head := &ListNode{}
	prev := &ListNode{}
	for i, v := range *a {
		println(i, v)
		cur := ListNode{Val: v, Next: nil}
		if i == 0 {
			head = &cur
			prev = &cur
		} else {
			prev.Next = &cur
			prev = prev.Next
		}
	}
	return head
}

func Transerval(h *ListNode) {
	for ; h != nil; h = h.Next {
		fmt.Printf("%d->", h.Val)
	}
	fmt.Println("nil")
}

func DoCountLen(h *ListNode) (int, *ListNode) {
	len := 0
	tail := &ListNode{}
	for ; h != nil; h = h.Next {
		len++
		tail = h
	}
	return len, tail
}
func FindTargetHead(h *ListNode, move int) (*ListNode, *ListNode) {
	targetHead := &ListNode{}
	prev := &ListNode{}
	pace := 0
	for ; pace < move; h = h.Next {
		if pace != move {
			prev = h
		}
		pace++
	}
	targetHead = h
	return targetHead, prev
}
func RightShiftLinkedListByKSteps(l *ListNode, k int) *ListNode {
	result := &ListNode{}
	len, tail := DoCountLen(l)
	shift := k % len
	if shift == 0 {
		return l
	}
	move := len - shift
	targetHead, prev := FindTargetHead(l, move)
	prev.Next = nil
	tail.Next = l
	result = targetHead
	return result
}

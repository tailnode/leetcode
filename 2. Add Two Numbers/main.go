package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	over := 0
	newNode := &ListNode{}
	newList := newNode
	node1, node2 := l1, l2
	zeroNode := &ListNode{}
	for {
		sum := node1.Val + node2.Val + over
		newNode.Val = sum % 10
		over = sum / 10
		node1, node2 = node1.Next, node2.Next
		if node1 != nil && node2 == nil {
			node2 = zeroNode
		} else if node1 == nil && node2 != nil {
			node1 = zeroNode
		} else if node1 == nil && node2 == nil {
			break
		}
		newNode.Next = &ListNode{}
		newNode = newNode.Next
	}
	if over == 1 {
		newNode.Next = &ListNode{
			Val: 1,
		}
	}
	return newList
}
func genList(s []int) *ListNode {
	if len(s) == 0 || s[0] < 0 || s[0] > 9 {
		panic("wrong slice")
	}
	list := &ListNode{
		Val: s[0],
	}
	node := list
	for _, v := range s[1:] {
		if v < 0 || v > 9 {
			panic("wrong slice")
		}
		node.Next = &ListNode{Val: v}
		node = node.Next
	}
	return list
}
func (l *ListNode) String() string {
	node := l
	out := make([]int, 0)
	for node != nil {
		out = append(out, node.Val)
		node = node.Next
	}
	return fmt.Sprint(out)
}
func main() {
	l1 := genList([]int{1, 2, 3, 4, 5, 6, 7, 8})
	l2 := genList([]int{1, 2, 3, 4, 5, 6, 7, 8})
	sum := addTwoNumbers(l1, l2)
	fmt.Println(sum)
}

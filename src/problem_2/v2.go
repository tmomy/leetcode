package problem_2

import "fmt"

const INT  = 10

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l = &ListNode{}
	if l1.Next == nil{
		l = addOneListNode(l2,l1.Val)
		return l
	}
	if l2.Next == nil {
		l = addOneListNode(l1,l2.Val)
		return l
	}
	l = addOneListNode(l1, l2.Val)
	l.Next = addTwoNumbers(l.Next, l2.Next)
	return l
}

func addOneListNode(l *ListNode, otherVar int) *ListNode {
	var v = l.Val + otherVar
	var diff = v - INT
	if diff < 0 {
		l.Val = v
	} else {
		l.Val = diff
		if l.Next == nil {
			l.Next = &ListNode{Val:1}
		} else {
			l.Next = addOneListNode(l.Next,1)
		}
	}
	return l
}

func printListNextNode(l *ListNode) int {
	if l.Next == nil {
		fmt.Println(l.Val)
	} else {
		fmt.Println(l.Val)
		printListNextNode(l.Next)
	}
	return 0
}
//@author: Tmomy
//@time: 2018/2/23 17:00
package problem_2

import "testing"

func TestAddTwoNumbers(t *testing.T)  {
	// test case
	oneListNode3 := ListNode{3, nil}
	oneListNode2 := ListNode{3, &oneListNode3}
	oneListNode := ListNode{3, &oneListNode2}
	towListNode3 := ListNode{3, nil}
	towListNode2 := ListNode{3, &towListNode3}
	towListNode := ListNode{3, &towListNode2}
	expected3 := ListNode{6, nil}
	expected2 := ListNode{6, &expected3}
	expected := ListNode{6, &expected2}
	actually1 := addTwoNumbers1(&oneListNode, &towListNode)
	actually := addTwoNumbers(&oneListNode, &towListNode)
	result, a := isEqual(&expected, actually)
	if result==1 {
		t.Errorf("TestAddTwoNumbers: [%v], actually: [%v]", &expected, a)
	}
	result1, a1 := isEqual(&expected, actually1)
	if result1 == 1 {
		t.Errorf("TestAddTwoNumbers1: [%v], actually: [%v]", &expected, a1)
	}
}

func isEqual(l *ListNode, a *ListNode)  (int, *ListNode) {
	i:=1
	for  i>0 {
		i++
		if l.Val != a.Val{
			i=0
			return 1, a
		} else if l.Next == nil {
			i = 0
		} else {
			l = l.Next
			a = a.Next
		}
	}
	return 0, nil
}

func TestPrintListNextNode(t *testing.T)  {
	oneListNode3 := ListNode{3, nil}
	oneListNode2 := ListNode{3, &oneListNode3}
	oneListNode := ListNode{3, &oneListNode2}
	printListNextNode(&oneListNode)
}
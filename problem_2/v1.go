package problem_2


type ListNode struct{
	Val int
	Next *ListNode
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	const i1  = 0
	const i2  = 0
	var v1 = listNodeToInt(l1,i1,i1)
	var v2= listNodeToInt(l2, i2,i2)
	var sum = v1 + v2
	var result = intToListNode(sum)
	return result
}
func listNodeToInt(l *ListNode, i int, value int64) (v int64) {
	var val, c int64
	val = int64(l.Val)
	c = int64(pow(10,i))
	value += val * c
	if l.Next == nil {
		return value
	} else {
		i += 1
		return listNodeToInt(l.Next,i,value)
	}
}

func intToListNode(ne int64) *ListNode {
	var l  = &ListNode{}
	if ne < 10 {
		l.Val = int(ne)
		return l
	}else{
		var next = ne / 10
		var value = ne % 10
		l.Val = int(value)
		l.Next = intToListNode(next)
		return l
	}

}

func max(num1, num2 int64) int64 {
	/* 声明局部变量 */
	var result int64
	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}

func pow(x, n int) int {
	ret := 1 // 结果初始为0次方的值，整数0次方为1。如果是矩阵，则为单元矩阵。
	for n != 0 {
		if n%2 != 0 {
			ret = ret * x
		}
		n /= 2
		x = x * x
	}
	return ret
}
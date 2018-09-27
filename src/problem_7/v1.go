package problem_7



func reverse(x int) int {
	max := 1<<31 -1
	min := - 1<<31

	res := 0
	flag := true
	for flag {
		if -10 < x && x < 10 {
			flag = false
			res += x
			break
		}
		remainder:= x%10
		x = x/10
		res = (res + remainder)*10
	}
	if res <min || res>max {
		res = 0
	}
	return res
}

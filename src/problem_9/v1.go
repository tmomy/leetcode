package problem_9



func isPalindrome(x int) bool {
	if x < 0 {
		return false

	}
	if -1 < x && x < 10 {
		return true
	}
	res := 0
	isOk := false
	flag := true
	rx := x
	for flag {
		if -1 < x && x < 10 {
			flag = false
			res += x
			break
		}
		remainder:= x%10
		x = x/10
		res = (res + remainder)*10
	}
	if res == rx {
		isOk = true
	}

	return isOk
}

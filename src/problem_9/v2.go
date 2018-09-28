package problem_9



func isPalindrome2(x int) bool {
	if x < 0 {
		return false

	}
	if -1 < x && x < 10 {
		return true
	}
	res := 0
	isOk := false
	rx := x
	for x != 0 {
		res = (res * 10) + (x % 10)
		x /= 10
	}

	if res == rx {
		isOk = true
	}

	return isOk
}

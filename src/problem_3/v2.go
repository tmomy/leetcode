package problem_3

import "fmt"

func lengthOfLongestSubstring(str string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	M := [1024]int{}
	s := 0
	maxLen := 0
	for i := 0; i < len(str); i++ {
		v := str[i] - 'a'

		if M[v] > 0 {
			for s < i && str[s] != str[i] {
				M[str[s] - 'a'] = 0
				s++
			}
			for s < i && str[s] == str[i] {
				M[str[s] - 'a'] = 0
				s++
			}
		}

		maxLen = max(maxLen, i - s + 1)

		M[v] = 1

		fmt.Printf("%d - %d = %d\n", i, s, maxLen)
	}

	return maxLen
}
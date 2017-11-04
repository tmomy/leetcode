package problem_3

import "fmt"


func lengthOfLongestSubstring1(s string) (int) {
	var subStingMap = make(map[string]int)
	var maxLength  = 0
	var c  = len(s)
	if c <2 {
		return c
	}
	for i:=0; i<c ; i++  {
		var charMap  = make(map[string]bool)

		for j:=i+1; j<c+1; j ++ {
			fmt.Println(i)
			var v = s[i:j]
			var v1  = s[j-1:j]
			_, ok := subStingMap[v]
			if ok {
				charMap[v1] = true
				continue
			} else {
				_, ok1 := charMap[v1]
				if ok1 {
					break
				} else {
					charMap[v1] = true
				}
				var length = j-i
				subStingMap[s[i:j]] = length
				if maxLength < length {
					maxLength = length
				}
			}
		}
	}
	fmt.Println(subStingMap)
	return maxLength
}

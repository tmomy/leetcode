//@author: Tmomy
//@time: 2018/2/24 17:29
// 以下代码摘抄自leetCode该题下提交的最优解代码
// 字符串s长度为n,算法描述:
// 先判断相同字符串回文子串，再以子串为中心判断左右字符是否相等，返回最大回文子串
package problem_5

func longestPalindrome2(s string) string {
	slen := len(s)
	if slen < 2 {
		return s
	}

	maxLen, maxLeft := 1, 0
	for i := 0; (i < slen) && (slen-i > (maxLen / 2)); {
		left, right := i, i
		for right < slen-1 && s[right] == s[right+1] {
			right++
		}
		i = right + 1

		for right < slen-1 && left > 0 && s[right+1] == s[left-1] {
			right++
			left--
		}

		if right-left+1 > maxLen {
			maxLen = right - left + 1
			maxLeft = left
		}
	}
	return s[maxLeft : maxLeft+maxLen]
}

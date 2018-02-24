//@author: Tmomy
//@time: 2018/2/24 15:35
//Longest Palindromic Substring, Given a string s, find the longest palindromic substring in s.
// You may assume that the maximum length of s is 1000.
// 字符串s长度为n,算法描述:
// 1.判断相同字符串回文子串，
// 2.以单个子串为中心判断左右字符是否存在回文子串，
// 3.以两个相同字符串为中心判断左右字符是否存在回文子串
// 4.返回最大回文子串
package problem_5

func longestPalindrome(s string) string {
	longest := struct {
		len int
		ls string
	}{0, ""}
	l := len(s)
	for i:=0; i<l; i++ {
		pl := 0
		for j:=i+1; j<=l&&s[i]==s[j-1]; j++ {
			pl ++
			if pl > longest.len {
				longest.len = pl
				if i ==0 {
					longest.ls = s[:j]
				}else {
					longest.ls = s[i:j]
				}
			}
		}
		for k:=1; i-k>=0&&i+k<l&&s[i-k] == s[i+k]; k++ {
			pl = 2*k+1
			if pl >= longest.len {
				longest.len = pl
				if i-k == 0{
					longest.ls = s[:i+k+1]
				}else {
					longest.ls = s[i-k:i+k+1]
				}
			}
		}
		for h:=0;i-h>=0&&i+h+1<l&&i+1<l&&s[i-h] == s[i+h+1]; h++ {
			pl = 2*(h+1)
			if pl >= longest.len {
				longest.len = pl
				if i-h == 0{
					longest.ls = s[:i+h+2]
				}else {
					longest.ls = s[i-h:i+h+2]
				}
			}
		}
	}
	return longest.ls
}
//@author: Tmomy
//@time: 2018/2/24 16:01

package problem_5

import "testing"

func TestLongestPalindrome(t *testing.T)  {
	for _, unit := range []struct {
		s        string
		expected string
	}{
		{"bbb", "bbb"},
		{"abbb", "bbb"},
		{"abbbbabb", "abbbba"},
		{"tattarrattat", "tattarrattat"},
	}{
		if actually:= longestPalindrome(unit.s); actually!=unit.expected {
			t.Errorf("longestPalindrome:[%v],actually:[%v]",unit, actually)
		}
		if actually:= longestPalindrome2(unit.s); actually!=unit.expected {
			t.Errorf("longestPalindrome2:[%v],actually:[%v]",unit, actually)
		}

	}
}

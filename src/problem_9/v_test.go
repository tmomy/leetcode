//@author: lizhijun
package problem_9

import (
	"testing"
)

func TestIsPalindrome(t *testing.T)  {
	for _, unit := range []struct {
		s        int
		expected bool
	}{
		{1534236469, false},
		{-123, false},
		{-320, false},
		{121, true},
		{1, true},
	}{

		if actually:= isPalindrome(unit.s); actually!=unit.expected {
			t.Errorf("longestPalindrome2:[%v],actually:[%v]",unit, actually)
		}
		if actually:= isPalindrome2(unit.s); actually!=unit.expected {
			t.Errorf("longestPalindrome2:[%v],actually:[%v]",unit, actually)
		}

	}
}


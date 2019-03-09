//@author: lizhijun
package problem_7

import (
	"testing"
)

func TestReverse(t *testing.T)  {
	for _, unit := range []struct {
		s        int
		expected int
	}{
		{1534236469, 0},
		{-123, -321},
		{-320, -23},
		{1534236469, 0},
	}{

		if actually:= reverse(unit.s); actually!=unit.expected {
			t.Errorf("longestPalindrome2:[%v],actually:[%v]",unit, actually)
		}

	}
}


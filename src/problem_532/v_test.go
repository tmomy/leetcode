package problem_532

import "testing"

func TestIsPerfectSquare(t *testing.T) {
	for _, unit := range []struct {
		input    []int
		k        int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, -1, 4},
		{[]int{1, 2, 3, 4, 5}, 1, 4},
	} {
		if actually := findPairs(unit.input, unit.k); actually != unit.expected {
			t.Errorf("isPerfectSquare:[%v],actually:[%v]", unit, actually)
		}

	}
}

//"""
//@author: Tmomy
//@time: 2018/2/23 15:34
//"""
package problem_4

import "testing"

func TestFindMedianSortedArrays(t *testing.T)  {
	for _, unit := range []struct{
		arrayOne  []int
		arrayTow  []int
		expected  float64
	} {
		{[]int{1,3,5}, []int{2,4,6}, 3.5},
		{[]int{1,3,5}, []int{2,4,6,7}, 4.0},
		{[]int{1,3}, []int{2,4,6}, 3},
	} {
		if output := findMedianSortedArrays(unit.arrayOne,unit.arrayTow); output != unit.expected {
			t.Errorf("findMedianSortedArrays: [%v], actually: [%v]", unit, output)
		}
	}
}
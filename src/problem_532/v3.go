// WL
// 2019/3/9
package problem_532

// leetCode case result: Runtime: 28 ms	 Memory: 5.9 MB
import "sort"

func findPairs3(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	ret := 0

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[j]-nums[i] == k {
				ret++
				break
			} else if nums[j]-nums[i] > k {
				break
			}
		}
	}

	return ret
}

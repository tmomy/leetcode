// WL
// 2019/3/6

package problem_532

// 查找效率最高， 但空间效率不行
// leetCode case result: Runtime: 20 ms	 Memory: 6.4 MB
func findPairs2(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	start := make(map[int]bool)
	unq := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := unq[nums[i]-k]; ok {
			start[nums[i]-k] = true
		}

		if _, ok := unq[nums[i]+k]; ok {
			start[nums[i]] = true
		}
		unq[nums[i]] = true
	}
	return len(start)
}

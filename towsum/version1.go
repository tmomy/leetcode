package towsum

func TowSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}
	valKeyMap := make(map[int]int)
	reArray := make([]int, 2)
	for key, val := range nums {
		exist, ok := valKeyMap[val]
		if ok {
			reArray[0] = exist
			reArray[1] = key
		} else {
			valKeyMap[target - val] = key
		}
	}
	return reArray
}
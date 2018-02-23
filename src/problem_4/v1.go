package problem_4

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	l := l1 + l2
	var array []*int
	var median float64
	i := 0
	j := 0
	for i<l1 && j<l2 {
		if nums1[i] <= nums2[j] {
			array = append(array, &nums1[i])
			i++
		} else {
			array = append(array, &nums2[j])
			j++
		}
	}
	for i<l1 {
		array = append(array, &nums1[i])
		i++
	}
	for j<l2 {
		array = append(array, &nums2[j])
		j++
	}
	m := l/2
	if l%2 == 0 {
		median = float64(*array[m-1] + *array[m])/2
		return median
	} else {
		median := float64(*array[m])
		return median
	}
}

package problem_532

// leetCode case result: Runtime: 28 ms	 Memory: 7.5 MB
func findPairs(num []int, k int) int {
	if k < 0 {
		return 0
	}
	pairNum := 0
	numLen := len(num)
	if numLen == 0 {
		return pairNum
	}
	numArray := newNumArray()
	for _, i := range num {
		checkPairs(numArray, &pairNum, i, k)
	}
	return pairNum
}

func checkPairs(numArray *numArray, pairNum *int, i, k int) bool {
	if k == 0 {
		dp, isFirst := numArray.get(i)
		if !isFirst && !dp.provBool {
			*pairNum++
			dp.provBool = true
		}
		return true
	}
	dp, _ := numArray.get(i)
	nextDp, _ := numArray.get(i + k)
	provDp, _ := numArray.get(i - k)
	linkNode(dp, nextDp)
	linkNode(provDp, dp)
	if !nextDp.provBool && dp.nextBool {
		*pairNum++
	}
	if !provDp.nextBool && dp.provBool {
		*pairNum++
	}
	nextDp.provBool = true
	provDp.nextBool = true
	return true
}

type numArray struct {
	Map map[int]*DiffPairs
}

func newNumArray() *numArray {
	nay := new(numArray)
	nay.Map = make(map[int]*DiffPairs)
	return nay

}

func (nay *numArray) get(key int) (*DiffPairs, bool) {
	first := false
	dp, ok := nay.Map[key]
	if !ok {
		dpf := DiffPairs{provBool: false, nextBool: false}
		dp = &dpf
		nay.Map[key] = dp
		first = true
	}
	return dp, first
}

func linkNode(prov, next *DiffPairs) {
	prov.nextDp = next
	next.provDp = prov
	return
}

type DiffPairs struct {
	provDp   *DiffPairs
	provBool bool
	nextDp   *DiffPairs
	nextBool bool
}

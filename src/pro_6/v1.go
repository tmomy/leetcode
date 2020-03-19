package pro_6

func convert(s string, numRows int) string {
	retMap := map[int]string{}
	std := 0
	end := 0
	l := len(s)
	for low := 1; end < l; low++ {
		md := low % numRows
		gap := 1
		if md == 1 || md == 0 {
			gap = numRows
			end = snd(end, gap, l)
			for i, _ := range s[std:end] {
				retMap[i] += s[std:end][i : i+1]
			}
		} else {
			end = snd(end, gap, l)
			retMap[numRows-md] += s[std:end]
		}
		std = end
		if md == 0 {
			low = 1
		}

	}
	ret := ""
	for i := 0; i < numRows; i++ {
		ret += retMap[i]
	}
	return ret
}

func snd(end, gap, l int) int {
	nEnd := end + gap
	if nEnd > l {
		return l
	}
	return nEnd
}

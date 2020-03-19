package pro_8

import "math"

var int2strMap = map[string]int{
	"-": -2,
	"+": -2,
	" ": -1,
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

func myAtoi(str string) int {
	max := float64(1<<31 - 1)
	min := float64(-1 << 31)
	index2bvMap := map[int]int{}
	index := 0
	isNegative := false
	isN := 0
	for _, bt := range str {
		bv, ok := int2strMap[string(bt)]
		if index == 0 && isN < -2 {
			return 0
		}
		if index == 0 && !ok {
			return 0
		}
		if index == 0 && bv == -1 {
			if isN == -2 {
				isN += bv
			}
			continue
		}
		if index == 0 && bv == -2 {
			if string(bt) == "-" {
				isNegative = true
			}
			isN += bv
			continue
		}
		if bv <= 0 && string(bt) != "0" {
			break
		}
		if ok {
			index2bvMap[index] = bv
			index++
		} else {
			break
		}
	}
	if index2bvMap == nil {
		return 0
	}
	ret := 0.0
	for i := 0; i < index; i++ {
		ret += float64(index2bvMap[i]) * math.Pow10(index-i-1)
	}
	if isNegative {
		ret = -ret
	}
	if ret <= min {
		return int(min)
	}
	if ret >= max {
		return int(max)
	}
	return int(ret)
}

package medium008

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	var num int
	var isNegative bool
	var isFirstValueAppeared bool
	for _, v := range s {
		if !isFirstValueAppeared {
			if v == '-' {
				isNegative, isFirstValueAppeared = true, true
				continue
			} else if v == '+' {
				isFirstValueAppeared = true
				continue
			} else if v == ' ' {
				continue
			} else if v >= '0' && v <= '9' {
				isFirstValueAppeared = true
				num, _ = runeToInt(v)
				continue
			}
			break
		}
		if v >= '0' && v <= '9' {
			tmp, _ := runeToInt(v)
			num = num*10 + tmp
		} else {
			break
		}
		if isNegative && -num < math.MinInt32 {
			return math.MinInt32
		} else if !isNegative && num > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	if isNegative {
		return -num
	}
	return num
}

func runeToInt(r rune) (int, error) {
	if r < 48 || r > 57 {
		return -1, fmt.Errorf("Invalid argument: %v", r)
	}
	return int(r - 48), nil
}

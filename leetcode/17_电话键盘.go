package leetcode

import (
	"fmt"
)

var numMap = map[uint8][]uint{
	2: {'a', 'b', 'c'},
	3: {'d', 'e', 'f'},
	4: {'g', 'h', 'i'},
	5: {'j', 'k', 'l'},
	6: {'m', 'n', 'o'},
	7: {'p', 'q', 'r', 's'},
	8: {'t', 'u', 'v'},
	9: {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	return doLetterCombine(digits, 0, len(digits)-1, "", nil)
}

func doLetterCombine(digits string, start, end int, letter string, ret []string) []string {
	if len(letter) == len(digits) {
		ret = append(ret, letter)
		return ret
	}

	arr := numMap[digits[start]-48]
	for j := 0; j < len(arr); j++ {
		letter = fmt.Sprintf("%s%c", letter, arr[j])
		ret = doLetterCombine(digits, start+1, end, letter, ret)
		letter = letter[:len(letter)-1]
	}

	return ret
}

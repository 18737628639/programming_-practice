package leetcode

var arr []string
var ret [][]string

func partition(s string) [][]string {
	arr, ret = make([]string, 0), make([][]string, 0)
	doPartition(s, 0)
	return ret
}



func doPartition(s string, start int)  {
	if start == len(s) {
		tmpArr := make([]string, len(arr))
		copy(tmpArr, arr)
		ret = append(ret, tmpArr)
		return
	}

	for i :=start; i < len(s); i ++ {
		str :=  s[start:i+1]
		// 当前子串不是会问子串
		if isPalindrome(str) {
			arr = append(arr, str)
			doPartition(s, i+1)
			arr = arr[:len(arr) -1]
		}
	}
}


func isPalindrome (subString string) bool {
	for i, j := 0, len(subString)-1; i < j; i, j = i+1, j-1 {
		if subString[i] != subString[j] {
			return false
		}
	}
	return true
}



package leetcode

func combinationSum3(k int, n int) [][]int {
	return doCombine3(k, n, nil, nil, 1, n)
}

func doCombine3(k int, n int, ret [][]int, arr []int, start, end int) [][]int {
	if len(arr) == k {
		if doSum(arr) == n {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			ret = append(ret, tmp)
			return ret
		}
	}
	if doSum(arr)+start > n {
		return ret
	}
	for i := start; i <= end && i < 10; i++ {
		arr = append(arr, i)
		ret = doCombine3(k, n, ret, arr, i+1, end)
		arr = arr[:len(arr)-1]
	}
	return ret
}

func doSum(arr []int) int {
	var sum int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

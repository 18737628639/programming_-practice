package leetcode

func permute(nums []int) [][]int {
	return doPermute(nums, 0, len(nums)-1, nil, nil)
}

func doPermute(nums []int, start, end int, arr []int, ret [][]int) [][]int {
	if len(arr) == len(nums) {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		ret = append(ret, tmp)
	}
	if start > end {
		return ret
	}

	for i := 0; i <= len(nums)-1; i++ {
		arr, flag := filterDup(arr, nums[i])
		ret = doPermute(nums, start+1, end, arr, ret)
		if !flag {
			arr = arr[:len(arr)-1]
		}
	}
	return ret
}

func filterDup(arr []int, elem int) ([]int, bool) {
	var flag bool
	for _, v := range arr {
		if v == elem {
			flag = true
		}
	}
	if flag {
		return arr, true
	}
	arr = append(arr, elem)
	return arr, false
}

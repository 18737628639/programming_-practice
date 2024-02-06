package leetcode

func permuteUnique(nums []int) [][]int {
	return doPermuteUnique(nums, 0, len(nums), nil, nil)
}

func doPermuteUnique(nums []int, start, end int, ret [][]int, arr []int) [][]int {
	if len(arr) == len(nums) {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		ret = setNonDupSubset(tmp, ret)
	}
	if start > end {
		return ret
	}
	for i := 0; i <= len(nums)-1; i++ {
		arr, flag := filterDup(arr, nums[i])
		ret = doPermuteUnique(nums, start+1, end, ret, arr)
		if !flag {
			arr = arr[:len(arr)-1]
		}
	}
	return ret
}

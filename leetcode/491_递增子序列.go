package leetcode

func findSubsequences(nums []int) [][]int {
	return buildSubsequences(nums, 0, len(nums)-1, nil, nil)
}

func buildSubsequences(nums []int, start, end int, arr []int, ret [][]int) [][]int {
	if len(arr) >= 2 {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		ret = setNonDupSubset(tmp, ret)
	}
	for i := start; i <= end; i++ {
		var isAdd bool
		if i == 0 || len(arr) == 0 {
			isAdd = true
			arr = append(arr, nums[i])
		} else if nums[i] >= arr[len(arr)-1] {
			isAdd = true
			arr = append(arr, nums[i])
		}
		ret = buildSubsequences(nums, i+1, end, arr, ret)
		if isAdd {
			arr = arr[:len(arr)-1]
		}
	}
	return ret
}

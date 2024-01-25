package leetcode

import "sort"

func subsets(nums []int) [][]int {
	var arr []int
	for i := 0 ; i < len(nums); i ++ {
		var flag bool
		for j := 0 ;j <len(arr) ; j++ {
			if arr[j] == nums[i] {
				flag = true
				break
			}
		}
		if !flag {
			arr = append(arr, nums[i])
		}
	}
	sort.Slice(nums, func(i, j int) bool {
		if nums[i] > nums[j] {
			return false
		}
		return true
	})
	return getSubsets(arr, nil, 0, len(arr)-1, nil)
}


func getSubsets (nums []int, ret[][]int, start, end int, arr []int) [][]int {
	if start > len(nums) {
		return ret
	}
	tmp := make([]int, len(arr))
	copy(tmp, arr)
	ret = append(ret, tmp)
	for i := start; i <=end; i ++ {
		num := nums[i]
		arr = append(arr, num)
		ret = getSubsets(nums, ret, i+1, end, arr)
		arr = arr[:len(arr)-1]
	}
	return ret
}

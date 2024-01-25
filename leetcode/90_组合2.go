package leetcode

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		if nums[i] > nums[j] {
			return false
		}
		return true
	})
	return doSubsetWithDup(nums, nil, 0, len(nums)-1, nil)
}

func doSubsetWithDup(nums []int, ret [][]int, start, end int, arr []int) [][]int {
	if start > len(nums) {
		return ret
	}
	tmp := make([]int, len(arr))
	copy(tmp, arr)
	ret = setNonDupSubset(tmp, ret)
	for i := start; i <= end; i++ {
		arr = append(arr, nums[i])
		ret = doSubsetWithDup(nums, ret, i+1, end, arr)
		arr = arr[:len(arr)-1]
	}
	return ret
}

func setNonDupSubset(arr []int, ret [][]int) [][]int {
	flag := false
	for _, l := range ret {
		if len(arr) != len(l) {
			continue
		}
		index := 0
		for ; index < len(l); index++ {
			if arr[index] != l[index] {
				flag = false
				break
			}
		}
		if index == len(l) {
			flag = true
			break
		}
	}
	if !flag {
		ret = append(ret, arr)
	}
	return ret
}

//[[] [1] [1 1] [1 1 2] [1 1 2 2] [1 1 2 2 3] [1 1 2 3] [1 1 3] [1 2] [1 2 2] [1 2 2 3] [1 2 3] [1 3] [2] [2 2] [2 2 3] [2 3] [3]]

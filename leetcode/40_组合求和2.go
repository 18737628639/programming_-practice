package leetcode

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		if candidates[i] < candidates[j] {
			return true
		}
		return false
	})
	return doCombinationSum2(candidates, target, 0, len(candidates)-1, nil, nil)
}

func doCombinationSum2(candidates []int, target int, start, end int, ret [][]int, arr []int) [][]int {
	sum := doSum(arr)
	if sum > target {
		return ret
	}
	if sum == target {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		ret = append(ret, tmp)
	}
	for i := start; i <= end; i++ {
		if i-1 >= start && candidates[i-1] == candidates[i] {
			continue
		}
		arr = append(arr, candidates[i])
		ret = doCombinationSum2(candidates, target, i+1, end, ret, arr)
		arr = arr[:len(arr)-1]
	}
	return ret
}

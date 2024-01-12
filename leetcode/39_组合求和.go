package leetcode

import (
	"reflect"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	ret := doCombinationSum(candidates, target, 0, len(candidates)-1, nil, nil)
	return filterDulpArr(ret)
}

func doCombinationSum(candidates []int, target int, start, end int, ret [][]int, arr []int) [][]int {
	curSum := getSum(arr)
	if curSum > target {
		return ret
	}
	if curSum == target {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		ret = append(ret, tmp)
		return ret
	}

	for i := start; i <= end; i++ {
		arr = append(arr, candidates[i])
		ret = doCombinationSum(candidates, target, i, end, ret, arr)
		arr = arr[:len(arr)-1]
	}
	return ret
}

func getSum(arr []int) int {
	var sum int
	for _, v := range arr {
		sum += v
	}
	return sum
}

func filterDulpArr(source [][]int) [][]int {
	var ret [][]int
	for _, a := range source {
		sort.Slice(a, func(i, j int) bool {
			if a[i] < a[j] {
				return true
			}
			return false
		})
	}
	for i := 0; i < len(source); i++ {
		isExist := false
		for j := i + 1; j < len(source); j++ {
			if reflect.DeepEqual(source[i], source[j]) {
				isExist = true
				break
			}
		}
		if !isExist {
			ret = append(ret, source[i])
		}
	}
	return source
}

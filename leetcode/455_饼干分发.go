package leetcode

import "sort"

func findContentChildren(g []int, s []int) int {
	var count int
	sort.Slice(g, func(i, j int) bool {
		if g[i] > g[j] {
			return true
		}
		return false
	})
	sort.Slice(s, func(i, j int) bool {
		if s[i] > s[j] {
			return true
		}
		return false
	})

	var assignIndex int
	for i := 0; i < len(g); i++ {
		for j := assignIndex; j < len(s); j++ {
			if s[j] >= g[i] {
				count++
				assignIndex = j + 1
				break
			}
		}
	}
	return count
}

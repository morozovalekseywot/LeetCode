package leetcode

import "sort"

// https://leetcode.com/problems/merge-intervals/
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	start, end := intervals[0][0], intervals[0][1]
	res := make([][]int, 0, len(intervals))

	for _, v := range intervals {
		if v[0] <= end {
			if end < v[1] {
				end = v[1]
			}
		} else {
			res = append(res, []int{start, end})
			start, end = v[0], v[1]
		}
	}
	res = append(res, []int{start, end})

	return res
}

func main() {

}

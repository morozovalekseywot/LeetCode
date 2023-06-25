package leetcode

import "sort"

func topKFrequent(nums []int, k int) []int {
	mn, mx := nums[0], nums[0]
	for _, v := range nums {
		mn = min(mn, v)
		mx = max(mx, v)
	}
	counter := make([]Pair, mx-mn+1)
	shift := mn
	for _, v := range nums {
		counter[v-shift].fr += 1
		counter[v-shift].sc = v
	}

	sort.Slice(counter, func(i, j int) bool {
		return counter[i].fr > counter[j].fr
	})
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = counter[i].sc
	}

	return ans
}

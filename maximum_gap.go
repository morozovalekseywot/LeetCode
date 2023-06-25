package leetcode

func maximumGap(nums []int) int {
	n := len(nums)
	m, mn, mx := n-1, minInSlice(nums), maxInSlice(nums)
	if n <= 2 || mn == mx {
		return mx - mn
	}

	buckets := make([]Pair, m)
	for i := 0; i < m; i++ {
		buckets[i] = Pair{INT32_MAX, INT32_MIN}
	}

	for _, v := range nums {
		idx := (v - mn) * m / (mx - mn)
		if v == mx {
			idx = m - 1
		}

		if v < buckets[idx].fr {
			buckets[idx].fr = v
		}
		if v > buckets[idx].sc {
			buckets[idx].sc = v
		}
	}

	maxGap := 0
	for i := 0; i < m-1; {
		j := i + 1
		for j < m && buckets[j].fr == INT32_MAX {
			j++
		}
		if j == m {
			break
		}
		maxGap = max(maxGap, buckets[j].fr-buckets[i].sc)
		i = j
	}

	return maxGap
}

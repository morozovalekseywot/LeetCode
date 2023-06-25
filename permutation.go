package leetcode

import "sort"

// Двигаясь справа налево, находим элемент, нарушающий убывающую последовательность.
// Меняем его с минимальным элементом, большим нашего, стоящим правее.
// Перевернем правую часть.
func nextPermutation(nums []int) {
	n := len(nums)
	pivot := n - 2
	for ; pivot >= 0; pivot-- {
		if nums[pivot] < nums[pivot+1] {
			break
		}
	}
	if pivot >= 0 {
		idx_min := pivot + 1
		for i := pivot + 1; i < n; i++ {
			if nums[pivot] < nums[i] && nums[i] <= nums[idx_min] {
				idx_min = i
			}
		}
		nums[pivot], nums[idx_min] = nums[idx_min], nums[pivot]
	}
	for i, j := pivot+1, n-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func permute(nums []int) [][]int {
	sort.Ints(nums)
	shift := -nums[0] + 1
	expected := 1
	for i := 0; i < len(nums); i++ {
		nums[i] += shift
		expected *= i + 1
	}
	count := 1
	result := make([][]int, 0, expected)
	sub := make([]int, len(nums))
	copy(sub, nums)
	result = append(result, sub)
	for count < expected {
		nextPermutation(nums)
		sub = make([]int, len(nums))
		copy(sub, nums)
		result = append(result, sub)
		count++
	}

	if shift == 0 {
		return result
	}

	for i := 0; i < expected; i++ {
		for j := 0; j < len(nums); j++ {
			result[i][j] -= shift
		}
	}

	return result
}

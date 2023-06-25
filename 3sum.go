package leetcode

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	n := len(nums)
	eps := Abs(nums[0] + nums[1] + nums[2] - target)
	ans := nums[0] + nums[1] + nums[2]
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			req := target - nums[i] - nums[j]

			rightIdx := n - 1
			leftIdx := j + 1
			for leftIdx+1 < rightIdx {
				midIdx := (leftIdx + rightIdx) / 2
				if nums[midIdx] > req {
					rightIdx = midIdx
				} else if nums[midIdx] < req {
					leftIdx = midIdx
				} else {
					leftIdx = midIdx
					rightIdx = midIdx
				}
			}

			if leftIdx == rightIdx {
				if nums[leftIdx] == req {
					return target
				}
				if Abs(req-nums[leftIdx]) < eps {
					eps = Abs(req - nums[leftIdx])
					ans = nums[i] + nums[j] + nums[leftIdx]
				}
			} else if leftIdx+1 == rightIdx {
				if Abs(req-nums[leftIdx]) < eps {
					eps = Abs(req - nums[leftIdx])
					ans = nums[i] + nums[j] + nums[leftIdx]
				}
				if Abs(req-nums[rightIdx]) < eps {
					eps = Abs(req - nums[rightIdx])
					ans = nums[i] + nums[j] + nums[rightIdx]
				}
			}

		}
	}

	return ans
}

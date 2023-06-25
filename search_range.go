package leetcode

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	low, high := 0, len(nums)-1
	var mid int
	for low+1 < high {
		mid = (low + high) / 2
		if nums[mid] < target {
			low = mid
		} else if nums[mid] > target {
			high = mid
		} else {
			low = mid
			high = mid
			break
		}
	}
	if nums[low] == target {
		high = low
	} else if nums[high] == target {
		low = high
	}
	if low == high && nums[low] == target {
		for ; high < len(nums)-1 && nums[high] == nums[high+1]; high++ {

		}
		for ; low > 0 && nums[low] == nums[low-1]; low-- {

		}

		return []int{low, high}
	}

	return []int{-1, -1}
}

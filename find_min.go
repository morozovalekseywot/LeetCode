package leetcode

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/
//func findMin(nums []int) int {
//	if len(nums) == 1 {
//		return nums[0]
//	}
//	low, high := 0, len(nums)-1
//	var mid int
//	for low+1 < high {
//		mid = (low + high) / 2
//		if nums[mid] > nums[high] {
//			low = mid
//		} else {
//			high = mid
//		}
//	}
//
//	return min(nums[low], nums[high])
//}

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	low, high := 0, len(nums)-1
	if nums[low] < nums[high] {
		return nums[low]
	}
	var mid int
	for low+1 < high && nums[low] >= nums[high] {
		mid = (low + high) / 2
		if nums[mid] > nums[high] {
			low = mid
		} else if nums[mid] < nums[high] {
			high = mid
		} else {
			low++
			high--
		}
	}

	return min(nums[low], nums[high])
}

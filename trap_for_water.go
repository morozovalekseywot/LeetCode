package leetcode

// https://leetcode.com/problems/trapping-rain-water/description/
//func trap(height []int) int {
//	if len(height) == 0 {
//		return 0
//	}
//	idxMax := 0
//	idxMin := 0
//	sum := 0
//	for i := 1; i < len(height); i++ {
//		if height[i] < height[idxMin] {
//			sum += height[i]
//		}
//	}
//
//	return sum
//}

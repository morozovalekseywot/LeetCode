package leetcode

// https://leetcode.com/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	curSum := 0
	for _, num := range nums {
		if curSum < 0 {
			curSum = 0
		}
		curSum += num
		if maxSum < curSum {
			maxSum = curSum
		}
	}

	return maxSum
}

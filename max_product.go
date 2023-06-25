package leetcode

func maxProduct(nums []int) int {
	begin := 0
	for ; begin < len(nums) && nums[begin] == 0; begin++ {
	}
	if begin == len(nums) {
		return 0
	}
	maxProd := nums[begin]
	maxPositive := nums[begin]
	minNegative := nums[begin]
	for i := begin + 1; i < len(nums); i++ {
		if nums[i] > 0 {
			maxPositive = max(nums[i], maxPositive*nums[i])
			minNegative = min(nums[i], minNegative*nums[i])
		} else {
			maxPositive, minNegative = minNegative*nums[i], min(maxPositive*nums[i], nums[i])
		}
		maxProd = max(maxProd, maxPositive)
	}

	return maxProd
}

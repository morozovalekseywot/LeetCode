package main

import (
 "fmt"
 "strconv"
)

func findSolutions(nums []int, target int, currentSum int, index int, expression string, last int) {
	if index == len(nums) {
		if currentSum == target && last == index {
			fmt.Println(expression)
		}
		return
	}

	end := index
	num := nums[end]
	k := 10
	for last < end{
		num += k * nums[end - 1]
		k *= 10
		end--
	}

	// По условию нельзя ставить знак перед первым числом
	if(last != 0){
		findSolutions(nums, target, currentSum+num, index+1, expression+"+"+strconv.Itoa(num), index+1)
		findSolutions(nums, target, currentSum-num, index+1, expression+"-"+strconv.Itoa(num), index+1)
	} else{
		findSolutions(nums, target, currentSum+num, index+1, strconv.Itoa(num), index+1)
	}
	findSolutions(nums, target, currentSum, index+1, expression, last)
}

func main() {
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	target := 200

	findSolutions(nums, target, 0, 0, "", 0)
}

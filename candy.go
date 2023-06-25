package leetcode

// https://leetcode.com/problems/candy/
func candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)

	// Проверяем слева направо
	for i := 1; i < n; i++ {
		if ratings[i-1] < ratings[i] {
			candies[i] = candies[i-1] + 1
		}
	}

	// Проверяем справа налево
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}

	res := n // всем дополнительно по одной конфете, так как массив изначально инициализирован нулями
	for i := 0; i < n; i++ {
		res += candies[i]
	}

	return res
}

package leetcode

// https://leetcode.com/problems/edit-distance/description/
func minDistance1(word1 string, word2 string) int {
	a, b := []rune(word1), []rune(word2)
	m, n := len(a), len(b)

	cur := make([]int, n+1)
	for j := 1; j <= n; j++ {
		cur[j] = j
	}

	for i := 1; i <= m; i++ {
		pred := i // минимальное количество операций, чтобы перевести полную первую строку в пустую вторую
		for j := 1; j <= n; j++ {
			if a[i-1] == b[j-1] {
				cur[j] = cur[j-1]
			} else {
				cur[j] = min(cur[j-1], min(pred, cur[j])) + 1
			}
			pred = cur[j]
		}
	}

	return cur[n]
}

func minDistance2(word1 string, word2 string) int {
	a, b := []rune(word1), []rune(word2)
	m, n := len(a), len(b)

	dp := make([][]int, m+1) // minimum number of operations to convert word1[0..i) to word2[0..j).
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j])) + 1
			}
		}
	}

	return dp[m][n]
}

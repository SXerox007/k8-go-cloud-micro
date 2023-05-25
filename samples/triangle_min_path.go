package samples

//time complexity of O(n^2) and a space complexity of O(n)
func MinimumTotal(triangle [][]int) int {
	rows := len(triangle)
	if rows == 0 {
		return 0
	}

	// initialize the memoization table with the values of the last row
	dp := make([]int, len(triangle[rows-1]))
	copy(dp, triangle[rows-1])

	// start from second to last row
	for i := rows - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}

	return dp[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

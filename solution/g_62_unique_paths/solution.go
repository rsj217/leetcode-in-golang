package g_62_unique_paths

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	// 第一列初始化为1
	for y := 0; y < m; y++ {
		dp[y][0] = 1
	}
	// 第一行初始化为1
	for x := 0; x < n; x++ {
		dp[0][x] = 1
	}

	for y := 1; y < m; y++ {
		for x := 1; x < n; x++ {
			dp[y][x] = dp[y-1][x] + dp[y][x-1]
		}
	}
	return dp[m-1][n-1]
}

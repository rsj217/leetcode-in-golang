package g_509_fibonacci_number

func fib(n int) int {
	if n <= 1 {
		return n
	}

	dp := []int{0, 1}
	for i := 2; i <= n; i++ {
		dpn := dp[0] + dp[1]
		dp[0], dp[1] = dp[1], dpn
	}
	return dp[1]
}

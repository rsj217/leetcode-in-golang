package g_59_spiral_matrix_ii

func generateMatrix(n int) [][]int {

	var ans = make([][]int, 0, n)
	for i := 0; i < n; i++ {
		ans = append(ans, make([]int, n, n))
	}

	l, r := 0, n-1
	u, d := 0, n-1
	number := 1
	for l < r && u < d {
		// l -> r
		for x := l; x < r; x++ {
			ans[u][x] = number
			number++
		}

		// u -> d
		for y := u; y < d; y++ {
			ans[y][r] = number
			number++
		}

		// l <- r
		for x := r; l < x; x-- {
			ans[d][x] = number
			number++
		}

		// u <- d
		for y := d; u < y; y-- {
			ans[y][l] = number
			number++
		}

		l++
		r--
		u++
		d--
	}
	if n%2 != 0 {
		ans[n/2][n/2] = number
	}
	return ans
}

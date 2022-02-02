package main

import "fmt"

func main() {
	n := 5
	var ans = make([][]int, 0, n)
	for i := 0; i < n; i++ {
		ans = append(ans, make([]int, n, n))
	}
	fmt.Println(ans)
}

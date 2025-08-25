package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	grid1 := make([]string, N)

	grid2 := make([]string, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&grid1[i])

	}

	for i := 0; i < N; i++ {
		fmt.Scan(&grid2[i])
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid1[i][j] != grid2[i][j] {
				fmt.Print(i+1, j+1)
				return
			}
		}
	}
}

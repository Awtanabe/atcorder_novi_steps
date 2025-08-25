package main

import "fmt"

func main() {
	var N, M int

	fmt.Scan(&N, &M)

	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a[i])
	}

	b := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Scan(&b[i])
	}

	pastaAlreadyTakeMap := make(map[int]int, N)

	for _, v := range a {
		if _, ok := pastaAlreadyTakeMap[v]; ok {
			pastaAlreadyTakeMap[v] += 1
		} else {
			pastaAlreadyTakeMap[v] = 1
		}
	}

	for _, d := range b {
		if _, ok := pastaAlreadyTakeMap[d]; ok {
			if pastaAlreadyTakeMap[d] <= 0 {
				fmt.Print("No")
				return
			}
			pastaAlreadyTakeMap[d]--
		} else {
			fmt.Print("No")
			return
		}
	}

	fmt.Print("Yes")
}

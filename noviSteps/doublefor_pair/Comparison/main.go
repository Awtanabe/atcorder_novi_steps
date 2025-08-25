package main

import "fmt"

func SetData(d []int, times int) {
	for i := 0; i < times; i++ {
		fmt.Scan(&d[i])
	}

}

func main() {
	var N, M, c int
	fmt.Scan(&N, &M)

	d1 := make([]int, N)
	d2 := make([]int, M)

	SetData(d1, N)
	SetData(d2, M)

	for _, d := range d1 {
		for _, dd := range d2 {
			if d <= dd {
				c++
			}
		}
	}
	fmt.Print(c)
}

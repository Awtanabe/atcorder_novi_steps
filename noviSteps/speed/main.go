package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4, 5}

	d := []int{}
	d = append(d, 0)

	for i := 0; i < len(data)-1; i++ {
		d = append(d, data[i]+data[i+1])
	}

	fmt.Print(d)
}

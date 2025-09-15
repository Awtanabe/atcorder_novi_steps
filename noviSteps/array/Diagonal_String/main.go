package main

import "fmt"

func main() {
	data := make([]string, 3)

	for i := 0; i < 3; i++ {
		fmt.Scan(&data[i])
	}

	for i := 0; i < 3; i++ {
		fmt.Print(string(data[i][i]))
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		text := scanner.Text()
		// fmt.Println(text)

		for i := range text {
			fmt.Println(string(text[i]))
			fmt.Println("===")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み取りエラー:", err)
	}
}

### メモ

- stringはイミュータブルなので []runeにキャストする必要がある
- 文字列はrangeできる

```
package main

import "fmt"

func main() {
	var S string
	var res int

	fmt.Scan(&S)

	for _, c := range S {

		if c == 'v' {
			res += 1
		}
		if c == 'w' {
			res += 2
		}
	}

	fmt.Print(res)
}

```
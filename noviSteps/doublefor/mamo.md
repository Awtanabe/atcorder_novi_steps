### 多次元配列のデータの確保の仕方

```
func main() {
	var H, W, d int

	fmt.Scan(&H, &W)

	data := make([][]int, H)

	for i := 0; i < H; i++ {
		data[i] = make([]int, W)
		for j := 0; j < W; j++ {
			fmt.Scan(&d)
			data[i][j] = d
		}
	}
}
```

### 単純な配列なら appendなしでindexでデータを入れられる

```
package main

import "fmt"

func main() {
	var H, W, c int

	fmt.Scan(&H, &W)
	data := make([]string, H)

	for i := 0; i < H; i++ {
		fmt.Scan(&data[i])
	}

	for _, d := range data {

		for _, cc := range d {
			if cc == '#' {
				c++
			}
		}
	}

	fmt.Print(c)
}
```
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	in := bufio.NewReader(os.Stdin)

// 	var H, W int
// 	fmt.Fscan(in, &H, &W)

// 	cnt := make([]int, W)

// 	for i := 0; i < H; i++ {
// 		var line string
// 		fmt.Fscan(in, &line) // 行全体を1トークンとして読む（長さはW）
// 		for j := 0; j < W; j++ {
// 			if line[j] == '#' {
// 				cnt[j]++
// 			}
// 		}
// 	}

// 	// 出力（半角スペース区切り）
// 	for j := 0; j < W; j++ {
// 		if j > 0 {
// 			fmt.Print(" ")
// 		}
// 		fmt.Print(cnt[j])
// 	}
// 	fmt.Println()
// }

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var H, W int

	fmt.Scan(&H, &W)

	data := make([]string, H)
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < H; i++ {
		scanner.Scan()
		data[i] = scanner.Text()
	}

	for i := 0; i < W; i++ {
		c := 0
		for j := 0; j < H; j++ {
			if data[j][i] == '#' {
				c++
			}
		}
		fmt.Print(c, " ")
	}
}

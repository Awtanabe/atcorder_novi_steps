package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var H, W int
	fmt.Fscan(in, &H, &W)

	// 2次元配列（byte のグリッド）
	grid := make([][]byte, H)
	for i := 0; i < H; i++ {
		var line string
		fmt.Fscan(in, &line)   // 1行まるごと読む（長さは W のはず）
		grid[i] = []byte(line) // []byte にして保持（各セル: '.' or '#')
		// 必要なら長さチェック
		// if len(grid[i]) != W { panic("invalid width") }
	}

	// 列ごとにカウント
	cnt := make([]int, W)
	for j := 0; j < W; j++ {
		for i := 0; i < H; i++ {
			if grid[i][j] == '#' {
				cnt[j]++
			}
		}
	}

	// 出力
	for j := 0; j < W; j++ {
		if j > 0 {
			fmt.Print(" ")
		}
		fmt.Print(cnt[j])
	}
	fmt.Println()
}

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

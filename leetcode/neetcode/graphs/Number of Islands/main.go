package main

import "fmt"

// 2重のforループは 冷静に進んでいく
// 訪問を 0で塗り潰す、途中で 1を見つけると再帰する
// https://chatgpt.com/c/68fdb471-77d0-8322-bb0b-5ac6f4bdfebf
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	var dfs func(r, c int)

	fmt.Println("grid", grid)
	dfs = func(r, c int) {
		// 範囲外 or 水なら終了
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == '0' {
			return
		}
		// 訪問済みにする
		grid[r][c] = '0'
		fmt.Println("grid2", grid)

		// 上下左右に移動
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count++
				// 以下で全て探索する
				dfs(r, c)
			}
		}
	}
	return count
}

func main() {
	grid := [][]byte{
		[]byte("01110"),
		[]byte("01010"),
		[]byte("11000"),
		[]byte("00000"),
	}

	result := numIslands(grid)
	fmt.Println("Number of islands:", result)
}

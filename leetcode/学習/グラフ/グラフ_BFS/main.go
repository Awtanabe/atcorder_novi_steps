package main

import "fmt"

// 問題
// https://interviewcat.dev/p/coding-interviewcat/graph-bfs-dfs

// BFS は隣接リストで表されたグラフを探索する
func BFS(graph [][]int) {
	N := len(graph)
	if N == 0 {
		return
	}

	visited := make([]bool, N)
	queue := []int{0} // ルートを 0 とする

	for len(queue) > 0 {
		// 先頭を取り出す（popLeft）
		u := queue[0]
		queue = queue[1:]

		// (2) すでに探索済みならスキップ
		if visited[u] {
			continue
		}

		// 探索済みにする
		visited[u] = true
		fmt.Printf("ノード %d に到達しました\n", u)

		// 隣接ノードをキューに追加
		for _, v := range graph[u] {
			if !visited[v] {
				queue = append(queue, v)
			}
		}
	}
}

//    (0)
//    / \
//   (1) (2)
//    \ /
//    (3)

func main() {
	// 隣接リストの例
	graph := [][]int{
		// ⭐️インデックスが値
		{1, 2}, // 0番からノード → 1, 2に行ける
		{0, 3}, // 1番からノード → 0, 3に行ける
		{0, 3}, // 2番からノード → 0, 3に行ける
		{1, 2}, // 3番からノード → 1, 2に行ける
	}

	BFS(graph)
}

// 連番じゃないケース

// package main

// import "fmt"

// func BFS(graph map[int][]int, start int) {
// 	visited := map[int]bool{}
// 	queue := []int{start}

// 	for len(queue) > 0 {
// 		u := queue[0]
// 		queue = queue[1:]

// 		if visited[u] {
// 			continue
// 		}
// 		visited[u] = true
// 		fmt.Println("訪問:", u)

// 		for _, v := range graph[u] {
// 			if !visited[v] {
// 				queue = append(queue, v)
// 			}
// 		}
// 	}
// }

// func main() {
// 	graph := map[int][]int{
// 		1: {3, 4}, // 1 → 3, 4
// 		3: {1},    // 3 → 1
// 		4: {1, 2}, // 4 → 1, 2
// 		2: {4},    // 2 → 4
// 	}

// 	BFS(graph, 1)
// }

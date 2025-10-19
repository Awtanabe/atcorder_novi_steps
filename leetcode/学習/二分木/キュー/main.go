package main

// https://interviewcat.dev/p/coding-interviewcat/binary-tree-bfs-dfs

import "fmt"

// Node 構造体を定義
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// BFS 関数（幅優先探索）
func BFS(root *Node) {
	if root == nil {
		return
	}

	// キューをスライスで実装
	queue := []*Node{root}

	for len(queue) > 0 {
		// 先頭を取り出す
		node := queue[0]
		queue = queue[1:]

		// 値を出力
		fmt.Println(node.Val)

		// 子ノードがあれば追加
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

func main() {
	// 木構造を作成
	root := &Node{
		Val: 1,
		Left: &Node{
			Val:   2,
			Left:  &Node{Val: 4},
			Right: &Node{Val: 5},
		},
		Right: &Node{
			Val:   3,
			Left:  &Node{Val: 6},
			Right: &Node{Val: 7},
		},
	}

	// BFS 実行
	BFS(root)
}

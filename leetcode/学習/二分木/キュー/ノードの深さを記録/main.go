package main

import "fmt"

// Node 構造体
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// depth 関数: 各ノードの深さを返す
func depth(root *Node) []int {
	if root == nil {
		return nil
	}

	// キューに (ノード, 深さ) のペアを入れる
	type pair struct {
		node  *Node
		level int
	}
	queue := []pair{{root, 0}}
	var ans []int

	for len(queue) > 0 {
		// 先頭を取り出す
		p := queue[0]
		queue = queue[1:]

		// 深さを記録
		ans = append(ans, p.level)

		// 子ノードを追加（深さ +1）
		if p.node.Left != nil {
			queue = append(queue, pair{p.node.Left, p.level + 1})
		}
		if p.node.Right != nil {
			queue = append(queue, pair{p.node.Right, p.level + 1})
		}
	}

	return ans
}

func main() {
	// テスト用の木構造
	root := &Node{
		Val: 1,
		Left: &Node{
			Val:   2,
			Left:  &Node{Val: 4},
			Right: &Node{Val: 5},
		},
		Right: &Node{
			Val:   3,
			Right: &Node{Val: 6},
		},
	}

	fmt.Println(depth(root))
}

package main

import "fmt"

// Node 構造体
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

const INF = 1 << 30

// goodNodes: 経路上で最大値以上のノード数を返す
func goodNodes(root *Node) int {
	ans := 0

	// 再帰関数 dfs
	var dfs func(node *Node, maxValue int)
	dfs = func(node *Node, maxValue int) {
		if node == nil {
			return
		}

		// 経路上の最大値を更新
		if node.Val >= maxValue {
			ans++
		}

		if node.Val > maxValue {
			maxValue = node.Val
		}

		dfs(node.Left, maxValue)
		dfs(node.Right, maxValue)
	}

	dfs(root, -INF)
	return ans
}

func main() {
	// 木構造の例
	root := &Node{
		Val: 3,
		Left: &Node{
			Val:  1,
			Left: &Node{Val: 3},
		},
		Right: &Node{
			Val:   4,
			Left:  &Node{Val: 1},
			Right: &Node{Val: 5},
		},
	}

	fmt.Println(goodNodes(root)) // 出力: 4
}

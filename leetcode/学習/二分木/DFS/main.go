package main

import "fmt"

// Node構造体
type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

// DFS関数
func dfs(node *Node) {
	fmt.Printf("%sに着きました\n", node.Val)

	if node.Left != nil {
		fmt.Printf("%sから%sへ行く\n", node.Val, node.Left.Val)
		dfs(node.Left)
		fmt.Printf("%sから%sへ帰ってきた\n", node.Left.Val, node.Val)
	}

	if node.Right != nil {
		fmt.Printf("%sから%sへ行く\n", node.Val, node.Right.Val)
		dfs(node.Right)
		fmt.Printf("%sから%sへ帰ってきた\n", node.Right.Val, node.Val)
	}

	fmt.Printf("%sから出ていきます\n", node.Val)
}

func main() {
	// ノードを作成
	nodeG := &Node{Val: "G"}
	nodeH := &Node{Val: "H"}
	nodeI := &Node{Val: "I"}
	nodeD := &Node{Val: "D", Left: nodeG, Right: nodeH}
	nodeE := &Node{Val: "E"}
	nodeF := &Node{Val: "F", Right: nodeI}
	nodeB := &Node{Val: "B", Left: nodeD, Right: nodeE}
	nodeC := &Node{Val: "C", Left: nodeF}
	nodeA := &Node{Val: "A", Left: nodeB, Right: nodeC}

	// DFSを開始
	dfs(nodeA)
}

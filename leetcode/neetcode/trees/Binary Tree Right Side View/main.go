package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func rightSideView(root *TreeNode) []int {

// 	res := []int{}
// 	var dfs func(node *TreeNode)
// 	dfs = func(node *TreeNode) {

// 		if node == nil {
// 			return
// 		}
// 		if node.Right == nil && node.Left != nil {
// 			res = append(res, node.Left.Val)
// 			return
// 		}

// 		res = append(res, node.Val)

// 		dfs(node.Right)
// 	}
// 	dfs(root)
// 	return res
// }

func rightSideView(root *TreeNode) []int {
	var res []int

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(res) {
			res = append(res, node.Val)
		}
		dfs(node.Right, depth+1)
		dfs(node.Left, depth+1)
	}

	dfs(root, 0)
	return res
}

func main() {
	l := &TreeNode{Val: 2}
	r := &TreeNode{Val: 3}
	root := &TreeNode{Val: 1, Left: l, Right: r}

	fmt.Println(rightSideView(root))
}

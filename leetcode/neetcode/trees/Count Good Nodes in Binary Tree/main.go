package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var dfs func(node *TreeNode, maxVal int) int
	dfs = func(node *TreeNode, maxVal int) int {
		if node == nil {
			return 0
		}

		// good node なら 1 を返し
		// good でないなら 0 を返す
		res := 0
		if node.Val >= maxVal {
			res = 1
		}

		// 大きかったらその値に更新してる
		maxVal = max(maxVal, node.Val)
		res += dfs(node.Left, maxVal)
		res += dfs(node.Right, maxVal)

		return res
	}

	return dfs(root, root.Val)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 2,1,1,3,null,1,5
func main() {

	r := &TreeNode{Val: 1}
	z := &TreeNode{Val: 5}
	e := &TreeNode{Val: 1}
	l := &TreeNode{Val: 3, Left: e, Right: z}
	root := &TreeNode{Val: 2, Left: l, Right: r}

	// 1,2,-1,3,4

	// t := &TreeNode{Val: 2}
	// h := &TreeNode{Val: 4}
	// l := &TreeNode{Val: 3, Left: h, Right: t}
	// root := &TreeNode{Val: 3, Left: l}

	fmt.Println(goodNodes(root))
}

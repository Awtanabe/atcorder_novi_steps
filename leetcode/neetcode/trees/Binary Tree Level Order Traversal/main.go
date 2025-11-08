package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	res := [][]int{}
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		// 新しい階層を作っている
		if len(res) == depth {
			// 			res := [][]int{}
			// res = append(res, []int{})
			// res[0] = append(res[0], 1)
			// fmt.Println("Hello, 世界", res)
			res = append(res, []int{})
		}

		res[depth] = append(res[depth], node.Val)

		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}

	dfs(root, 0)

	return res
}

func main() {

	b := &TreeNode{Val: 6}
	c := &TreeNode{Val: 7}
	d := &TreeNode{Val: 3, Right: c, Left: b}

	e := &TreeNode{Val: 4}
	f := &TreeNode{Val: 5}
	g := &TreeNode{Val: 2, Right: f, Left: e}
	h := &TreeNode{Val: 1, Right: g, Left: d}

	fmt.Println(levelOrder(h))
}

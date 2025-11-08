package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}

	if max(p.Val, q.Val) < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if min(p.Val, q.Val) > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	a := &TreeNode{Val: 2}
	b := &TreeNode{Val: 1, Right: a}
	c := &TreeNode{Val: 4}
	d := &TreeNode{Val: 3, Right: c, Left: b}

	e := &TreeNode{Val: 7}
	f := &TreeNode{Val: 9}
	g := &TreeNode{Val: 8, Right: f, Left: e}
	h := &TreeNode{Val: 5, Right: g, Left: d}

	fmt.Println(lowestCommonAncestor(h, d, g).Val)

}

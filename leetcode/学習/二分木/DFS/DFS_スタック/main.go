package main

import "fmt"

// 関数スタックをキューで実装してる

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

func main() {
	nodeG := &Node{"G", nil, nil}
	nodeH := &Node{"H", nil, nil}
	nodeI := &Node{"I", nil, nil}
	nodeD := &Node{"D", nodeG, nodeH}
	nodeE := &Node{"E", nil, nil}
	nodeF := &Node{"F", nil, nodeI}
	nodeB := &Node{"B", nodeD, nodeE}
	nodeC := &Node{"C", nodeF, nil}
	nodeA := &Node{"A", nodeB, nodeC}

	type frame struct {
		node  *Node
		state int
	}
	stack := []frame{{nodeA, 0}}

	for len(stack) > 0 {
		top := &stack[len(stack)-1]

		switch top.state {
		case 0:
			fmt.Printf("%sに着きました\n", top.node.Val)
			top.state++
			if top.node.Left != nil {
				fmt.Printf("%sから%sへ行く\n", top.node.Val, top.node.Left.Val)
				stack = append(stack, frame{top.node.Left, 0})
			}
		case 1:
			top.state++
			if top.node.Right != nil {
				fmt.Printf("%sから%sへ行く\n", top.node.Val, top.node.Right.Val)
				stack = append(stack, frame{top.node.Right, 0})
			}
		case 2:
			fmt.Printf("%sから出ていきます\n", top.node.Val)
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				parent := stack[len(stack)-1].node
				fmt.Printf("%sから%sへ帰ってきた\n", top.node.Val, parent.Val)
			}
		}
	}
}

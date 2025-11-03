package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// スライスとデクリメントで実装
func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	// ① すべてのノードをスライスに保存
	nodes := []*ListNode{}
	cur := head
	for cur != nil {
		nodes = append(nodes, cur)
		cur = cur.Next
	}

	// n - 1, n -2って -- でデクリメントしていけば実現できるのかcur
	// ② 前後から交互に繋ぎ直す
	i, j := 0, len(nodes)-1
	for i < j {
		fmt.Println("i", i, nodes[i].Next.Val, nodes[j].Val)
		nodes[i].Next = nodes[j] // 前から i 番目 → 後ろから j 番目
		i++
		fmt.Println(i, j)
		if i >= j {
			break
		}
		nodes[j].Next = nodes[i] // 後ろから j 番目 → 次の前側 i 番目
		j--
	}

	// ③ 最後のノードを nil にする（リストの終端を明示）
	nodes[i].Next = nil
}

func main() {
	head := &ListNode{Val: 2}
	head.Next = &ListNode{Val: 4}
	head.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next = &ListNode{Val: 8}
	reorderList(head)
}

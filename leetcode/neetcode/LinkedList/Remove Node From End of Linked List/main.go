package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	cur := head
	for cur != nil {
		// 追加してから cur.Nextの更新が普通か！！！！
		nodes = append(nodes, cur)
		cur = cur.Next
	}

	removeIndex := len(nodes) - n
	if removeIndex == 0 {
		return head.Next
	}

	nodes[removeIndex-1].Next = nodes[removeIndex].Next
	return head
}

func main() {
	// a := ListNode{Val: 4}
	// b := ListNode{Val: 3, Next: &a}
	// c := ListNode{Val: 2, Next: &b}
	// l := ListNode{Val: 1, Next: &c}

	a := ListNode{Val: 5}
	// fmt.Println(removeNthFromEnd(&l, 2))
	removeNthFromEnd(&a, 1)
}

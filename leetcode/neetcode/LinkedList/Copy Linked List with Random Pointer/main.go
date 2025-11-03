package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	oldToCopy := map[*Node]*Node{nil: nil}

	cur := head
	for cur != nil {
		copy := &Node{Val: cur.Val}
		oldToCopy[cur] = copy
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		copy := oldToCopy[cur]
		copy.Next = oldToCopy[cur.Next]
		copy.Random = oldToCopy[cur.Random]
		cur = cur.Next
	}

	return oldToCopy[head]
}

func printList(head *Node) {
	cur := head
	i := 0
	for cur != nil {
		r := "null"
		if cur.Random != nil {
			r = fmt.Sprintf("%d", cur.Random.Val)
		}
		fmt.Printf("[%d] Val=%d, Random=%s\n", i, cur.Val, r)
		cur = cur.Next
		i++
	}
	fmt.Println("-----")
}

func main() {
	// 手動でリスト構築
	n0 := &Node{Val: 3}
	n1 := &Node{Val: 7}
	n2 := &Node{Val: 4}
	n3 := &Node{Val: 5}

	// Next
	n0.Next = n1
	n1.Next = n2
	n2.Next = n3

	// Random
	n0.Random = nil
	n1.Random = n3
	n2.Random = n0
	n3.Random = n1

	fmt.Println("Original:")
	printList(n0)

	copied := copyRandomList(n0)

	fmt.Println("Copied:")
	printList(copied)

	// アドレス違う確認
	fmt.Printf("original n0 = %p\n", n0)
	fmt.Printf("copied head = %p\n", copied)
}

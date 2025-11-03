package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// l1 1 2 3, l2も 4,5,6
// 最終的には 5 7 9 でlistを繋げる 975が正解らしいんだけど
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		v1 := 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		v2 := 0
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		val := v1 + v2 + carry
		carry = val / 10
		cur.Next = &ListNode{Val: val % 10}
		cur = cur.Next
	}

	// dummy は使わない Next からが正解
	return dummy.Next
}

func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	cur := head
	for _, v := range nums[1:] {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return head
}

func printList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}

func main() {

	l1 := buildList([]int{1, 2, 3})
	l2 := buildList([]int{4, 5, 6})

	result := addTwoNumbers(l1, l2)
	printList(result)
}

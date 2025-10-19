package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func (l *LinkedList) Append(val int) {

	if l.Head == nil {
		l.Head = &ListNode{Value: val}
		return
	}

	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &ListNode{Value: val}

}

func (l *LinkedList) Remove(val int) {
	if l.Head == nil {
		return
	}
	if l.Head.Value == val {
		l.Head = l.Head.Next
		return
	}

	prev := l.Head
	cur := l.Head.Next
	for cur != nil {
		if cur.Value == val {
			prev.Next = cur.Next
			return
		}
		prev = cur
		cur = cur.Next
	}
}

func (l *LinkedList) Search(val int) *ListNode {

	cur := l.Head

	if cur.Value == val {
		return cur
	}

	for cur != nil {
		if cur.Value == val {
			return cur
		}
		cur = cur.Next
	}

	return nil
}

func (l *LinkedList) Display() {

	if l.Head == nil {
		fmt.Println("空")
	}

	cur := l.Head

	for cur != nil {
		fmt.Println(cur.Value)
		cur = cur.Next
	}
}

func main() {

	list := LinkedList{}

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Display()
	fmt.Printf("%+v\n", list.Search(2))
	list.Remove(2)
	list.Display()
}

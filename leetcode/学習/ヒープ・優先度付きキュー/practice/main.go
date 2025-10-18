package main

import (
	"container/heap"
	"fmt"
)

// IntHeap は Min Heap（最小ヒープ）を表す型
type IntHeap []int

// sort.Interface の実装
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // 小さい方が優先（Min Heap）
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// heap.Interface の実装
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	ary := []int{1, 5, 8, 0, 2, 6, 3, 9, 10, 4, 7, 11}

	// (1) heapify: O(N)
	h := IntHeap(ary)
	heap.Init(&h) // ← Pythonの heapify に相当
	fmt.Println("heap after Init:", h)
	// 例: [0 1 3 5 2 6 8 9 10 4 7 11]

	// (2) heappush: O(log N)
	heap.Push(&h, 5)
	fmt.Println("heap after Push:", h)
	// 例: [0 1 3 5 2 5 8 9 10 4 7 11 6]

	// (3) heappop: O(log N)
	minimum := heap.Pop(&h).(int)
	fmt.Println("minimum:", minimum)
	fmt.Println("heap after Pop:", h)
	// 例: minimum = 0
	//     heap = [1 2 3 5 4 5 8 9 10 6 7 11]
}

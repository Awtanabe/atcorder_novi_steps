package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)
	for _, n := range nums {
		heap.Push(h, n)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}

func main() {
	fmt.Println(findKthLargest([]int{2, 3, 1, 5, 4}, 2))
}

// クイックソートで実装できればベター
// func findKthLargest(nums []int, k int) int {
//     k = len(nums) - k

//     var qs func(l, r int) int
//     qs = func(l, r int) int {
//         pivot, p := nums[r], l
//         for i := l; i < r; i++ {
//             if nums[i] <= pivot {
//                 nums[p], nums[i] = nums[i], nums[p]
//                 p++
//             }
//         }
//         nums[p], nums[r] = nums[r], nums[p]

//         if p > k {
//             return qs(l, p-1)
//         } else if p < k {
//             return qs(p+1, r)
//         } else {
//             return nums[p]
//         }
//     }

//     return qs(0, len(nums)-1)
// }

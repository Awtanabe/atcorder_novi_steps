package main

import "fmt"

// MinHeap は整数用の最小ヒープ
type MinHeap struct {
	data []int
}

// 要素を追加（挿入）
func (h *MinHeap) Push(x int) {
	h.data = append(h.data, x)
	h.heapifyUp(len(h.data) - 1) // 親と比較して上へ
}

// 最小値を取り出して削除
func (h *MinHeap) Pop() int {
	if len(h.data) == 0 {
		panic("heap is empty")
	}

	root := h.data[0] // 最小値（ルート）

	last := len(h.data) - 1
	h.data[0] = h.data[last] // 最後の要素をルートに移動
	h.data = h.data[:last]   // 最後の要素を削除
	h.heapifyDown(0)         // 下方向に再構築

	return root
}

// ヒープの最小値を取得（削除しない）
func (h *MinHeap) Peek() int {
	if len(h.data) == 0 {
		panic("heap is empty")
	}
	return h.data[0]
}

// 親に向かって調整
func (h *MinHeap) heapifyUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.data[parent] <= h.data[i] {
			break
		}
		h.data[parent], h.data[i] = h.data[i], h.data[parent]
		i = parent
	}
}

// 子に向かって調整
func (h *MinHeap) heapifyDown(i int) {
	n := len(h.data)
	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i

		if left < n && h.data[left] < h.data[smallest] {
			smallest = left
		}
		if right < n && h.data[right] < h.data[smallest] {
			smallest = right
		}

		if smallest == i {
			break
		}

		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		i = smallest
	}
}

// ヒープが空か確認
func (h *MinHeap) IsEmpty() bool {
	return len(h.data) == 0
}

// デバッグ出力用
func (h *MinHeap) Print() {
	fmt.Println(h.data)
}

// ---- main ----
func main() {
	h := &MinHeap{}

	h.Push(3)
	h.Push(1)
	h.Push(5)
	h.Push(2)

	fmt.Print("ヒープ構造: ")
	h.Print() // 内部の配列構造を表示

	fmt.Println("最小値:", h.Peek()) // → 1

	fmt.Println("Pop:", h.Pop()) // → 1
	fmt.Println("Pop:", h.Pop()) // → 2
	fmt.Println("Pop:", h.Pop()) // → 3
	fmt.Println("Pop:", h.Pop()) // → 5
}

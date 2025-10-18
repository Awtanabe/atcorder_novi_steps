package main

// A ---1--- B ---2--- C
// |         |
// 4         3
// |         |
// D --------

import (
	"container/heap"
	"fmt"
	"math"
)

// ノード間の辺
type Edge struct {
	to, cost int
}

// グラフ構造: 隣接リスト
var graph = [][]Edge{
	0: {{1, 1}, {3, 4}},         // A→B(1), A→D(4)
	1: {{0, 1}, {2, 2}, {3, 3}}, // B→A, B→C, B→D
	2: {{1, 2}},                 // C→B
	3: {{0, 4}, {1, 3}},         // D→A, D→B
}

// 優先度付きキューに使う要素
type Item struct {
	node int
	dist int
}

// Min-Heap（距離が小さい順）
type PQ []Item

func (pq PQ) Len() int           { return len(pq) }
func (pq PQ) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PQ) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PQ) Push(x any)        { *pq = append(*pq, x.(Item)) }
func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

// Dijkstra法
func dijkstra(start int) []int {
	n := len(graph)

	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt // 無限大で初期化
	}
	dist[start] = 0

	pq := &PQ{{node: start, dist: 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Item)
		if cur.dist > dist[cur.node] {
			continue // 既により短い距離が見つかっているならスキップ
		}

		for _, e := range graph[cur.node] {
			if dist[e.to] > cur.dist+e.cost {
				dist[e.to] = cur.dist + e.cost
				heap.Push(pq, Item{node: e.to, dist: dist[e.to]})
			}
		}
	}

	return dist
}

func main() {
	start := 0 // A
	dist := dijkstra(start)
	fmt.Println("Aから各ノードへの最短距離:", dist)
}

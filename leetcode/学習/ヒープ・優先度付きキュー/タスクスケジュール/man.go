package main

import (
	"container/heap"
	"fmt"
	"time"
)

type Task struct {
	name     string
	priority int       // 小さいほど優先度高い（Min Heap）
	execAt   time.Time // 実行予定時刻
}

// 優先度付きキュー
type TaskPQ []*Task

func (pq TaskPQ) Len() int { return len(pq) }
func (pq TaskPQ) Less(i, j int) bool {
	if pq[i].priority == pq[j].priority {
		return pq[i].execAt.Before(pq[j].execAt)
	}
	return pq[i].priority < pq[j].priority // 優先度が低い数値ほど先
}
func (pq TaskPQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *TaskPQ) Push(x any) {
	*pq = append(*pq, x.(*Task))
}
func (pq *TaskPQ) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// max heap
// min heapがあるから登録できるのでは
func main() {
	now := time.Now()

	pq := &TaskPQ{
		&Task{"Backup", 3, now.Add(5 * time.Second)},
		&Task{"Report", 1, now.Add(3 * time.Second)},
		&Task{"Update", 2, now.Add(1 * time.Second)},
	}
	heap.Init(pq)

	// 新しいタスクを途中で追加
	heap.Push(pq, &Task{"Security Scan", 2, now.Add(2 * time.Second)})

	for pq.Len() > 0 {
		task := heap.Pop(pq).(*Task)
		fmt.Printf("実行: %-15s 優先度:%d 実行予定:%s\n",
			task.name, task.priority, task.execAt.Format("15:04:05"))
	}
}

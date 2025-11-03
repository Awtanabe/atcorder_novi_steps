
###

- [][2]intのデータ構造にしてデータを保持するとは思わなかった
  - mapで保持してるとsortができないからだと思う


#### min heapのやり方もおもしろい

https://neetcode.io/solutions/top-k-frequent-elements

- キューに入れてる
  - IntComparatorで順番を調整してくれる

```go
func topKFrequent(nums []int, k int) []int {
    count := make(map[int]int)
    for _, num := range nums {
        count[num]++
    }

    heap := priorityqueue.NewWith(func(a, b interface{}) int {
        freqA := a.([2]int)[0]
        freqB := b.([2]int)[0]
        return utils.IntComparator(freqA, freqB) // 比較関数
    })

    for num, freq := range count {
        heap.Enqueue([2]int{freq, num})
        if heap.Size() > k {
            heap.Dequeue()
        }
    }

    res := make([]int, k)
    for i := k - 1; i >= 0; i-- {
        value, _ := heap.Dequeue()
        res[i] = value.([2]int)[1]
    }
    return res
}
```

### bucket sort

- インデックスを数として、値を出現回数
  - 最初に makeで配列を初期化する必要がある

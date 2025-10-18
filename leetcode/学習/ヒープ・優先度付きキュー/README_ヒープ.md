### ヒープ

- データ構造の一種でただの配列

### 3つの操作

- heapify: 事前にO(N)O(N)の操作を行うことで、配列がヒープの構造となります。
- heappush:
ヒープの構造を保ちながらO(logN)O(logN)
新しい値の追加を行います。
- heappop:
ヒープの構造を保ちながらO(logN)O(logN)
最大値（最優先値）の取り出し（配列から削除）を行います。

- メモ
  - log(n)が出てきたな

### ヒープ完全二分木構造

- min Heap(小さい値を優先する)
  - 親ノード <= 子ノード
- max Heap
  - 親ノード >= 子ノード
- 木みたいになっているけど、配列で表現する事が可能
- 親ノード <= 子ノードの関係になっている

### ソート

- バブルソート
  - O(n2)
  - ロジック
    - 比較してスイッチを繰り返す

```go
package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{5, 2, 9, 1, 5, 6}
	fmt.Println("Before:", arr)
	bubbleSort(arr)
	fmt.Println("After: ", arr)
}
```

- マージソート
  - log(n)
  - ロジック
    - イメージ: https://medium-company.com/%E3%83%9E%E3%83%BC%E3%82%B8%E3%82%BD%E3%83%BC%E3%83%88/
    - 分割して、分割した物同士をまーじする
      - mergeSortは再帰処理で 要素が1になるまで繰り返す
      - その後にmerg されていく

```go
package main

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// 残りを追加
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func main() {
	arr := []int{5, 2, 9, 1, 5, 6}
	fmt.Println("Before:", arr)
	sorted := mergeSort(arr)
	fmt.Println("After: ", sorted)
}
```

### ヒープ有無

## 問題

数列の部分和を高速に求める

### 累積

```
package main

import (
	"fmt"
)

func main() {
	// 入力例
	A := []int{1, 2, 3, 4, 5}
	N := len(A)

	// 累積和配列を作成 (サイズ N+1)
	// S[i] = A[0] + A[1] + ... + A[i-1]
	S := make([]int, N+1)
	for i := 0; i < N; i++ {
		S[i+1] = S[i] + A[i]
	}

	// 部分和を計算する関数
	sum := func(l, r int) int {
		// 区間 [l, r) の和を返す
		return S[r] - S[l]
	}

	// 例: A[1..3] の和 = 2+3+4 = 9
	fmt.Println(sum(1, 4)) // 出力: 9
}

```

### 動的計算

```
package main

import (
	"fmt"
)

func main() {
	// 入力例
	A := []int{1, 2, 3, 4, 5}
	N := len(A)

	// 累積和配列を作成 (サイズ N+1)
	// S[i] = A[0] + A[1] + ... + A[i-1]
	S := make([]int, N+1)
	for i := 0; i < N; i++ {
		S[i+1] = S[i] + A[i]
	}

	// 部分和を計算する関数
	sum := func(l, r int) int {
		// 区間 [l, r) の和を返す
		return S[r] - S[l]
	}

	// 例: A[1..3] の和 = 2+3+4 = 9
	fmt.Println(sum(1, 4)) // 出力: 9
}

```
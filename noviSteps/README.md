### 教材

https://atcoder-novisteps.vercel.app/workbooks

### メモ

- 各セクションの要点をmemo.mdにメモしておく


### 全体メモ

fmt.Scanの使い方


```
package main

import (
	"fmt"
)

func main() {
	var N, X, tmp int

	fmt.Scan(&N, &X)

	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		if tmp == X {
			continue // breakにすると処理が終了する
		}
		fmt.Print(tmp)
		fmt.Print(" ")
	}
}

```

- Scan
  - 不要なdataに入れるためだけの変数はいらない

```
	var N, c int

	fmt.Scan(&N)

	data := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&data[i])
	}
```

### sort

sort.Slice

Ordinary Number
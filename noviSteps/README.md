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


### map

- mapをfor ループする時に順序が保証されないので厄介

- mapの初期化 make

- setはmapで表現する

```
set := make(map[string]struct{})
set["apple"] = struct{}{}
set["banana"] = struct{}{}
```

- struct{}{}
  - struct{}は「フィールドなし構造体型」
  - struct{}{}を初期化

```
type Empty struct{}

Empty{}している => ⭐️無名のフィールドを持たない構造体の初期化
```


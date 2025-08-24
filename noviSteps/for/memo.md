### for

### 数列

pi  (1<i<n)

の表記は iは最小が2から始まる

Pi-1, Pi, Pi+1 となるらしい

### index利用

- index + 1などにするときは len(data)はインデックス外なので、比較の富豪は <= ではなく <
  - ex) 1 2 3のスライスは lenだと3だけどインデックスのマックスは2

### continue break

- contineはループのスキップ
- breakはループを終了

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
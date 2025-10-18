### まとめ

- sortのアルゴリズムの理解
  - go のsortは実際は違うけど、バブルソートのイメージでコードを見る

- 0 < 2などの判定が感覚的に理解できていないので、ソートの結果が


### sort の使い方

- バブルソートのイメージを持って良い
  - bacの場合
    - ba => ab
	- bc
	- bcの3回を見るイメージで良いと思う  
- ソートの判定ロジック
  - trueの場合はそのまま
  - falseの場合は入れ替えるイメージ！！
    - ex) return ri < rjとした場合 trueであればこの位置関係は変えない
```
// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sort"
)

func main() {

	a := []int{1, 2}

	sort.Slice(a, func(i, j int) bool {
		fmt.Println("i", i)
		fmt.Println("j", j)

		// return true にすると  前に
		// return falseにすると　後ろに
		 return a[i] < a[j]
	})

	fmt.Println("Hello, 世界", a)
}

```
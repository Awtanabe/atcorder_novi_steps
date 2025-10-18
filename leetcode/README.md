### web

https://interviewcat.dev/p/coding-interviewcat

### スプレッドシート

https://docs.google.com/spreadsheets/d/1IsU68j8A3nFHriA39If5tyeYQHQzcITgGZh0xONznQA/edit?gid=0#gid=0

### 進め方

- 導入を一旦全てやって全体像
  - [x] 累積和
  - [] 再帰呼び出し/バックトラック
  - [] 連結リスト
  - [] 二分探索
    - 基本の実装はできた
  - [] 二分木
  - [] グラフ
  - [] 動的計画法
  - [] ヒープ
    - いまいち分からんのよな

### データの操作

- 下記のパターンで大体できる
  - インデックス [num:num]
  - 二次元ループ: 全探索

### 二分探索実装できちゃった

- binarySeach

### スライスのインデックス

- a[i:j]
  - j未満
  - ex)a[0:0] は「要素0から0未満までを取り出す」ので、

```
ackage main

import "fmt"

func main() {
	aaa := []int{1, 2, 3, 4, 5}
	fmt.Println(aaa[:3]) // 4がインデックス3で3未満
	fmt.Println(aaa[3:]) // インデックス3以降

}

[1 2 3]
[4 5]
```

- 分割は簡単にできる

```
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
```


### 文字の扱い stringsメソッド

- count とか便利

https://zenn.dev/kou_pg_0131/articles/go-strings-functions

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

	fmt.Println(a)
}

```

### web

https://interviewcat.dev/p/coding-interviewcat

### スプレッドシート

https://docs.google.com/spreadsheets/d/1IsU68j8A3nFHriA39If5tyeYQHQzcITgGZh0xONznQA/edit?gid=0#gid=0

### lete codeの方針

- https://zenn.dev/kj455/articles/b48c6af88ba090
  - hardは方法論
  - mediumに時間を費やしてOK

### neet code

https://zenn.dev/kiitosu/articles/dffd838222d69b

- 150問
  - https://neetcode.io/practice?tab=neetcode150

### 進め方

- 導入を一旦全てやって全体像
  - [x] 累積和
  - [x] 二分探索
    - 基本の実装はできた
	- minの値を探す。
	- 他は最大値を探すとかみたい left < rightが終了条件
  - [x] 連結リスト
    - ざっくり。実装はおおよそできた
  - [] 二分木
    - https://zenn.dev/brainyblog/articles/data-structures-intro-binary-trees
	- https://chatgpt.com/share/68f4b5d0-1f98-8004-8cb4-1a749ff14884
	  - ツリーは左小さい、右大きいと分けられていれば探索範囲がすくなくなる
	  - ツリーを作るのが大変
  - [△] グラフ
    - 触れたくらい
	- キュー(幅)とスタック(深さ)
	  - グラフのデータ
	    - LikedListのleft・rightあるバージョン
  - [△] 動的計画法
    - 触れたくらい
  - [△] ヒープ
    - いまいち分からんのよな
  - [] 再帰呼び出し/バックトラック
    - subsetsをしっかり理解するところから
      - 再帰の関数呼び出しの流れを理解したんだけどしっくりこない

### データの操作

- 下記のパターンで大体できる
  - インデックス [num:num]
  - 二次元ループ: 全探索

### 二分探索実装できちゃった

- binarySeach

### 再帰のイメージ

- スタックの理解
- プログラミングの実行順番

```go
package main

import "fmt"

func rec(n int) {
	if n == 0 {
		fmt.Println("終了")
		return
	}

	fmt.Printf("前処理: n=%d\n", n)
	rec(n - 1) // スタックフレームに積み上げていく
	// 関数が優先で呼び出した関数が終了(再帰の場合はreturnで終わらせる)したら、呼び出しの直後から再開
	// 呼ばれた関数が終わるまで待機する。通常の関数でもやっているけど別関数だから違和感がないのか
	fmt.Printf("後処理: n=%d\n", n)
}

func main() {
	rec(2)
}

```

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

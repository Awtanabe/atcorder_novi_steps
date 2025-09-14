### 解説AC

- [x] シミュレーション 1 First Query Problem
  - 入力 Scannerの勉強で進まなかった
- [x] シミュレーション 2 => 二分探索なので後でやる
- [x] 全探索 ① typo
- [x] 全探索 ②
- [x] 多重 for文 3
- [x] 多重 for文 4
  - mod 除数でインデックスをぐるぐる回る
- [x] 多重 for文 5
  - atcorder => max min 閾値はテストをする癖をみにつけないと実装が漏れる ⭐️実務にも役立つ
- [x] 二次元グリット 2
- [x] シミュレーショ3
- [] スタック
  - スライスで表現できたはず
	  - インデックス指定じゃなくてappendで追加できるみたい
- [] キュー
  - キューもスタックと使うものは同じs
- [] 再帰関数
- [] スタック DP高速化
- [] 貪欲法
  - 効率良い選択をすること
	  - 大きい順に並べた方が良いとか
- [] 再帰全探索
- [x] Frequency

### スタック

- appendで追加できる
- 取り除く
  - data[1:] 

```
package main

import "fmt"

func main() {
	data := []int{1}
	fmt.Println("Hello, 世界", data)
	data = append(data, 2, 3, 4)
	fmt.Println("Hello, 世界", data)
	data = data[1:]
	fmt.Println("Hello, 世界", data)
}

Hello, 世界 [1]
Hello, 世界 [1 2 3 4]
Hello, 世界 [2 3 4]
```

### Aggregate

- 集約 解き方
  - バケットとは
	  - スライスのこと
		  - var buckect [10]int
  - バケットに状態を記録してカウントs


### 2025/08/30

#### 解説 acでやったものを見ていく

- test ディレクトリを作業場にする

##### 全探索 1 typo
- 文字の入れ替えをどおするか？で迷った
  - []runeに文字列を変換
  - s[i], s[i+1] = s[i+1], s[i]でスワップ
- ポイント
　- 横並びのスワップで良いのはやりやすかった
    - 横並びじゃなかったとしても、一致であればsortしたりすれば解決するだろうな


##### シミュレーション 2 savings ⭐️ => 二分探索なので後で

- 貯金箱に i円を入れていく
  - 1 <= i
- N円以上入っているのは何日め？

- ポイント
  - 普通に計算すると計算量が多いので、下記の2つしかアプローチがない
    - 数列の数式による計算量 0(1)
    - 二分探索
- 問題
  - 普通にやると オーバーフローとか速度的に問題ありっぽい
  s

##### シミュレーション First Query Problem

- 問題
  - 長さN個のN個の整数が渡される
  - Q個のクエリが渡される
    - クエリ
      - 1と2がある
        - 1は k x
          - インデックス kの数字を xに変更する
        - 2は k
          - インデックス kの値を出力する

- 振り返り
  - 問題の理解に手こずった
  - Scanを Scannerに変更するのも時間がかかった
  - 文字列から数字への変更も

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N, Q int

	fmt.Scan(&N)

	data := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&data[i])
	}

	fmt.Scan(&Q)
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < Q; i++ {
		if scanner.Scan() {
			text := scanner.Text()
			strs := strings.Split(text, " ")

			if strs[0] == "1" {
				k, _ := strconv.Atoi(strs[1])
				x, _ := strconv.Atoi(strs[2])
				data[k-1] = x
			}

			if strs[0] == "2" {
				k, _ := strconv.Atoi(strs[1])
				fmt.Println(data[k-1])
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "読み取りエラー:", err)
		}

	}

}
```

##### 集計(バケット・連想配列の利用) Frequency
  - 文字列 を取得
  - アルファベットは 26個あるので、intスライスを26個確保
    - 
  - 文字の扱いは runeで rune - 'a'で差分を確保 && byte('a' + 差分)で文字列を
  - ポイント
    - カウント方法
      - バケットのインデックスがそのままアルファベットの位置になるs
    - max と idxの変数2つ使わないといけない
      - うまくいなかったのは maxを追加してしまったのでアスキーコードの範囲外で string(re)で出力しても何も表示されなかった(数字99以降とかだったはず)

```go
import "fmt"

func main() {
	diff := 'b' - 'a'
	res := 'a' + diff
	fmt.Println("Hello, 世界", string(res))
}

```


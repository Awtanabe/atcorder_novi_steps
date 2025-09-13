### go の文法

#### 入力

- fmt.Scan とbufioの使い分け
  - Scan
    - データ数が 10⁴ 以下
  - bufio
    - データ数が 10⁵〜10⁶ 以上
- 小規模 / デバッグ用 → fmt.Scan

- 大規模 / 競プロ / 性能重視 → bufio.NewReader/Writer + fmt.Fscan/Fprintln
  - Fscan
  - Fprintln

- Scanner利用の場合
  - 文字列しか返さないので、Atoiで変換
    

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		text := scanner.Text() // 文字列として読み込まれる
		num, err := strconv.Atoi(text) // 数字に変換
		if err != nil {
			fmt.Println("数値変換エラー:", err)
			return
		}
		fmt.Println(num) // 数字として出力
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み取りエラー:", err)
	}
}

```

- split

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var N int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		if scanner.Scan() {
			text := scanner.Text()
			strs := strings.Split(text, " ")

			fmt.Println(strs[0])
			fmt.Println(strs[1])
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "読み取りエラー:", err)
		}

	}

}
```

#### 変数宣言

- var
- 省略

#### スライス

- ループ
- make利用

#### マップ

- makeの使い方
- 作成
- 存在確認

#### Sort

#### ループ

- continue
- break

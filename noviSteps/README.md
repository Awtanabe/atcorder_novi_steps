### 前提

- 茶色・緑・水色
- 緑になるには
  - 「3完」がほぼ確実で、「4完」は問題を解き終わっている数か！
    - C~Dの問題を解けるようになることか
- 水色
  - 「4完」を8〜9割以上の確率で、そして「5完」を3〜4割以上の確率
    - Eの問題まで解けるようになるか

### 教材

- novistep
https://atcoder-novisteps.vercel.app/workbooks

- abc problems

https://kenkoooo.com/atcoder/#/table/

- 凡人が緑になる 50問 kindle


### atcorder系の記事

### 水色に

https://note.com/jikky1618/n/n3298c09633dc

### 緑になるために必要なアルゴリズム

https://note.com/jikky1618/n/n0814c5803391

#### 水色に
https://qiita.com/_ken_/items/c32d4b2e680058abd77a

- 水色 アルゴリズムのレベル
  - https://qiita.com/_ken_/items/c32d4b2e680058abd77a#%E3%81%93%E3%81%AE%E3%81%93%E3%82%8D%E3%81%AB%E8%A6%9A%E3%81%88%E3%81%9F%E3%82%A2%E3%83%AB%E3%82%B4%E3%83%AA%E3%82%BA%E3%83%A0-2

- 緑
  - https://qiita.com/_ken_/items/c32d4b2e680058abd77a#%E3%81%93%E3%81%AE%E3%81%93%E3%82%8D%E3%81%AB%E8%A6%9A%E3%81%88%E3%81%9F%E3%82%A2%E3%83%AB%E3%82%B4%E3%83%AA%E3%82%BA%E3%83%A0-1

### 

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

- 存在確認

```go
		if price, ok := colorPriceMap[color]; ok {
			total += price
		} else {
			total += P[0] // デフォルト価格
		}
```

#### mapの使い方 Default Price

- Default Priceで color: priceでmapを作成し、color(食べたもの)でループして合計金額を合算していく

### map コツ

https://qiita.com/ryuken/items/6ef4a33fa5ad4819af85

### 問題

#### Pasta

- 考慮もれ、パスタが同じ種類の可能性を考慮できていなかった


### atcorderを実務で使う範囲

これだけ押さえれば十分（実装は10〜20行）
二分探索：sort.Search(n, func(i int) bool { 条件が単調 })
尺取り法：連続区間の合計/件数を O(N)
累積和：日次/時間帯の合算を O(1) 参照
BFS：queueで到達可能性・最短手数
最小ヒープ：上位K件や最短ジョブ選択（container/heap）



結論：**かなり当たってるけど “8〜9割ずっと安泰” とは言い切れない**、が実務の土台としてはめっちゃ強いです。

# まず事実

* 典型的な業務コード（API実装、CRUD、集計、ETL、バッチ、ログ処理）は
  **slice / map（＋頻度カウント＝バケット）＋全探索**でだいたい形になります。
  さらに `sort`, `slices` パッケージが扱えれば**日常の8割前後**は回せる体感、妥当です。
* 実務の難所は「アルゴリズムテク」より**入出力・並行処理・信頼性**（エラーハンドリング、リトライ、タイムアウト、整合性、観測性）で詰まりやすい。

# ただし、ここで困る

1. **データ量が跳ねる場面**

   * O(N²) 全探索が通らない（例：N=10^6 の重複検査・区間集計・最長〇〇）。
   * → **二分探索 / 尺取り / 前計算（prefix sum）/ ソート済み活用**が必要。
2. **スケジューラや最短経路系・優先度制御**

   * → **ヒープ（`container/heap`）**、**BFS/DFS** 基本。
3. **レコメンドやランキングの上位K件**

   * → ヒープ or 片側ソート＋**二分探索**。
4. **ストリーミング処理・バックプレッシャ**

   * → **goroutine / channel / `context` / rate limit** の設計。
5. **性能/メモリ要件**

   * → **プロファイリング（`pprof`）**, **逃がしベクトル（バッファ/プーリング/再利用）**。

# 実務で“最低限”押さえておくと世界が広がる小技（Go寄り）

* `sort.Search`（単調性がある判定に二分探索を被せる）
* **尺取り法（two pointers）**：スライディングウィンドウで O(N)
* **累積和 / いもす法（バケットの拡張）**：大量区間加算→最終一括復元
* **ヒープ**：上位K、最短ジョブ、Dijkstraの土台

  ```go
  // 最小ヒープの骨格
  type Item int
  type MinHeap []Item
  func (h MinHeap) Len() int            { return len(h) }
  func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
  func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
  func (h *MinHeap) Push(x any)         { *h = append(*h, x.(Item)) }
  func (h *MinHeap) Pop() any           { old:=*h; x:=old[len(old)-1]; *h=old[:len(old)-1]; return x }
  ```
* **BFS/DFS**（グリッド・グラフ）："最短手数"問題に即効薬
* **map＋バケット**の型癖最小化：`map[string]int` / `map[int]struct{}`（Set代用）
* **エラーハンドリング & 文脈**：`context.Context`、タイムアウト、リトライ設計
* **I/O最適化**：Scanner/Writer、JSONストリーム、DBのバルク

# 実務での“勝ち筋”まとめ

* **スライス・map・バケット・全探索**＝強い土台。まずはここを鉄板に。
* その上で、**二分探索／尺取り／累積和／ヒープ／BFS**の“軽量キット”を持つと、
  \*\*O(N²)をO(N log N) or O(N)\*\*へ落とせて、途端に“詰まらない人”になります。
* さらに\*\*並行処理（goroutine, channel, context）**と**観測性（ログ/メトリクス/トレース）\*\*は、業務での評価直結ポイント。

なので結論は：

> 日常実装の大半は slice/map/バケット＋全探索でいける。
> ただ、**規模・性能・並行性**の壁に当たるので、**最小限のアルゴリズム手筋**と**Goの並行処理・I/O・エラー設計**は押さえておくのが実務最適解、です。

### シート

https://docs.google.com/spreadsheets/d/1IsU68j8A3nFHriA39If5tyeYQHQzcITgGZh0xONznQA/edit?gid=1623330397#gid=1623330397

### stringsパッケージ


- split
https://zenn.dev/kou_pg_0131/articles/go-strings-functions#split(s%2C-sep-string)-%5B%5Dstring

- join

https://zenn.dev/kou_pg_0131/articles/go-strings-functions#join(elems-%5B%5Dstring%2C-sep-string)-string


### Encode and Decode Strings

- 文字数の塊、#で文字と数字の塊を分ける


### Valid Sudoku

- スクエアのやり方
  - ポイント
    - queare / 3 * 3 + i がポイント
    - 他は全部探索

### Longest Consecutive Sequence

- ハッシュセット利用(Set 重複ない配列)
  - 全てを探索するのに forループで探索できる

### 3Sum

- sort
  - sort.Intsもある []intのデータなら
- 重複排除 => hash Set⭐️を考えると計算量が減る
  - map setの使い方 key が複数の [3]int などでも保存できる
    - ex)tmp[[3]int{nums[i], nums[j], nums[k]}] = struct{}{}

### Container With Most Water

- どのパターンが最適かどうやってチェックするか分からなかった
  - ⭐️全て計算して、一番の最大値を取れば良い => maxAreaを更新していく
    - ⭐️これ人間は全部の棒を見ていなさそうだけど実際は、全ての棒を見て判断してる

### Valid Palindrome

- 文字のformatの方法がよくなかった
  - a ~ z 0 ~ 9で見れば記号は全部カットできた
  - 左右から移動して探索した

## slising window

- indexをずらして探索していく問題を指しそうだ

### 復讐[]Longest Substring Without Repeating Characters

substring: 部分文字列
subsequence: 連続 xyzなど

### 復讐[] Longest Repeating Character Replacement

- 長さと左
  - 長さは (left - right) + 1
  - 4本の木の話
    - 3つの空間 + 1
- 大事な数式
  - 必要な置換数 = (ウィンドウの長さ) - (最頻文字の数)
    - (right - left + 1) - maxCount


### 復讐[] Permutation in String
- むずい


### Evaluate Reverse Polish Notation

- スタックの理解でできる
- 加算したりしたら、使った要素はpopする

### Generate Parentheses

- dfs を使ってる(全探索・頭悪い)
  - 2つに毎回分岐させるやつ！
- bfsの場合
  - キューになる
- バックトラッキング(賢いDFS)
  - 探索途中で“この道は間違い”と判断したら戻って別の道を試す
  - ex) 陽光なカッコだけ作る

### heap・priority que

- ⭐️heapとは？
  - 最大値・最小値をソートを使わないで、効率よく要素を追加削除して引き出せるようにするロジック
  - 実際はライブラリ使うことが多いはず
  - 実務で使うことは少ないけど、データベース・検索・スケジュールではよく使われる
    - ジョブなどSQSを使うから使わない
      - メモリベース二分木構造を保持してる
- heaoは内部で2分木を使っている
  - heap without libraryが理解できれば良い、
    - sort を使わないで二分木のデータ構成を利用した実装している
      - ライブラリを使っている
- heap practiceのコード
  - 常に最小の値を取り出したいロジック
    - sortを使わないで、スライスを使って実装できる
  - 内部では
    - ヒープ構造（完全二分木）
- ⭐️heapライブラリ
  - パッケージ

- 解答
  - container/heapは標準なので使える
    - ライブラリ使うなら、解説できないと
  - ベストは クイックソートで実装したい

### Trie retrieval（検索） trie（トライ）

- 検索手法のアルゴリズムとして独立してる
- インクリメントsearch

### Reorder Linked List

- スライスとデクリメントで
  - n - 1, n-2などの実現はデクリメントで表現できる
    - nはスライスの長さ

### グラフ

- グラフの問題は特徴
  - 問題のデータがグラフのデータ構造になっているものを対象としてる
    - ⭐️グラフの特徴を理解しておかないと、どのアルゴリズムを使うべきか分からないから、グラフの特徴の解説が多いのか
  - よく使う解法
    - dfs 

### stack

#### dailyTemperatures

- stackで解こうとするとむずい
  - インデックスをキューに入れて行って比較する

### Valid Parentheses

- 解説のコードはもっとシンプルな回答だった
  - mapで open closeを持って popの値と比較していた


### Car Fleet


- データがfloatかどうか？で結果が変わりえた
  - /割り算をした時に、切り捨てが変化する⭐️

### binary search

### []復習 Koko Eating Bananas

### [] 復習 Search in Rotated Sorted Array

- neet codeの回答の方がシンプルかも
  - ⭐️順番がおかしくなってるところまで移動する処理を入れている

### LinkedList

- LinkedList形式をスライスに変換するパターン

### []復習 Remove Node From End of Linked List

- 普通にできなかった

### Copy Linked List with Random Pointer

- コピーしているだけだった

### Add Two Numbers

- dummyは使わない dummy.Nextで使う
- 数字の繰り上がりなどの計算

### [] 復習LRU Cache

- 大枠はできていたと思う

### Tree

- LinkdedList
  - left rightがあるバージョン

### Lowest Common Ancestor in Binary Search Tree

- むずいな、なんとなくしかわからない
  - データの規則性が 左が小さい、右が大きいて決まってるなら やりようはあるけど

### [] Binary Tree Level Order Traversal⭐️

- 新しい階層を作っている
  - [][]int{} の初期化 && appendの理解不足

#### [] Count Good Nodes in Binary Tree ⭐️

- フラグで 0 1を返して加算してる
- maxVal は大きかったら更新してる

### backtracking

### [] Combination Sum
- むずいな！
  - パターン1
    - 選ぶ・選ばないパターン!!! 前にも出てた
      - 選ばないは前のリセットをして
  - パターン2
    - return num +再帰関数パターン


### gredy 

### []復讐 Maximum Subarray


### エッジケース考えられてないな

- 


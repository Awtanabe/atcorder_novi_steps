### 連結リスト

- ノード: 要素
  - val
  - next: 次の要素
    - リンク: ポインタ
- 計算量:
  - アクセス O(n)
  - 追加削除: O(1)
    - ただ要素探索があるので O(n)

### Search・insert・delete のコード

### 双方向リスト

- ノートが nextとprev両方を持っている

### 実装

LikedList
  - head: Nodeを持っている
  - メソッド
    - search
    - append
    - remove
    - display 
ListeNode
- 実装問題は出てこなうけど、探索、削除などは出てくるらしい

### 二分木

- ノードの対して子が2つ以下の木の事

### Breadth First Search, BFS（幅優先探索）・ Depth First Search, DFS（深さ優先探索）

- 木の構造を知る
  - BFS(幅優先)
    - 階層ごとにみていく
    - 実現するにはキュー
      - キュー: https://interviewcat.dev/p/coding-interviewcat/queue-intro

- キューで実装
  - nodeはleft rightがあるので下記の
    - a: node leftにb right にc
    - b c
      - b: left にd right にc
        - d e
      - c: left に fので
        - f
### BFS（幅優先探索）が有効な時

- 親より上の情報が必要な時
- 階層ごとの情報が必要な時

### DFS

- 再帰処理を使う
  - 再帰処理の動きの理解にまるかも DFSの解説
    - https://interviewcat.dev/p/coding-interviewcat/binary-tree-bfs-dfs
  - ⭐️行きと帰りは、関数スタックのフローで実現してる

### DFSが有効な時

- あるノードから根までの経路の情報が必要なとき
### 基本

- 全探索
- グラフ
- キュー



### グラフ

https://chatgpt.com/c/68877554-d6a4-8333-89a4-b5b13a462f46

頂点と辺で構成されるデータ構造
=> slice goでは


```

頂点数 N = 4
辺
1 - 2
1 - 3
3 - 4

g := [][]int{
    {1, 2}, // 0
    {0},    // 1
    {0, 3}, // 2
    {2},    // 3
}

-- 考え方 スライスのインデックスを利用
1: 2, 3 // index1 1の頂点
2: 1    // 
3: 1, 4
4: 3



----
package main

import (
    "fmt"
)

func main() {
    var n, m int
    fmt.Scan(&n, &m)

    // グラフ (隣接リスト)
    g := make([][]int, n)
    for i := 0; i < m; i++ {
        var a, b int
        fmt.Scan(&a, &b)
        a--
        b--
        g[a] = append(g[a], b)
        g[b] = append(g[b], a) // 無向なら両方向
    }

    fmt.Println(g)
}

```

- BFS（幅優先探索）

```
package main

import (
	"fmt"
)

func bfs(graph [][]int, start int) []int {
	n := len(graph)
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = -1 // -1 は未訪問
	}

	queue := []int{start}
	dist[start] = 0

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:] // 1:はスライスのindex 1移行

		for _, next := range graph[v] {
			if dist[next] == -1 { // 未訪問なら
				dist[next] = dist[v] + 1
				queue = append(queue, next)
			}
		}
	}

	return dist
}

func main() {
	graph := [][]int{
		{1, 3},    // 0
		{0, 2, 4}, // 1
		{1},       // 2
		{0, 4},    // 3
		{1, 3},    // 4
	}

	dist := bfs(graph, 0)

	for i, d := range dist {
		fmt.Printf("0からノード%dまでの距離: %d\n", i, d)
	}
}

```

- dfs

```
func dfs(v int, g [][]int, seen []bool) {
    seen[v] = true
    for _, nv := range g[v] {
        if seen[nv] {
            continue
        }
        dfs(nv, g, seen)
    }
}
```
### set

go には Setがない
- 下記2つが基本
  - mapを使う
  - ライブラリを使う

### map

- mapをfor ループする時に順序が保証されないので厄介

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


### ライブラリ golang-set


```
go get github.com/deckarep/golang-set/v2


import mapset "github.com/deckarep/golang-set/v2"

set := mapset.NewSet[int]()
set.Add(1)
set.Add(2)

fmt.Println(set.Contains(1)) // true
fmt.Println(set.Cardinality()) // 要素数
```


### グラフ

- 対象とそれらの間でのつながりを表すデータ構造
- グラフのデータ構造
  - val
  - left
  - right

### 実装

- スタックかキューか
  - スタック(深さ)
    - 後入れ先出し（LIFO：Last In, First Out）
      - 入れて、お尻から出す
  - ⭐️ここ理解必要かも
    - https://chatgpt.com/s/t_68f4df872b388191abb204bf1e4f4327

```go
stack := []int{}
stack = append(stack, 1)
stack = append(stack, 2)
stack = append(stack, 3)
fmt.Println(stack) // [1 2 3]

// pop（最後の要素を取り出す）
top := stack[len(stack)-1]
stack = stack[:len(stack)-1]
fmt.Println("取り出した:", top) // 3

```

- キュー(幅)
  - 先入れ先出し（FIFO：First In, First Out）
    - 入れたものから出す

```go
queue := []int{}
queue = append(queue, 1)
queue = append(queue, 2)
queue = append(queue, 3)
fmt.Println(queue) // [1 2 3]

// dequeue（先頭を取り出す）
front := queue[0]
queue = queue[1:]
fmt.Println("取り出した:", front) // 1

```
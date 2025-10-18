### 再帰処理

#### 実装
- 関数内で呼ぶ
  - 関数がスタックされる(再帰スタック)
  - O(N)のメモリが必要になる
- 終了条件を設置

### 再帰処理は問題点

- 直感的でわかりやすいらしい
- 問題
  - 重複計算
  - 計算量が O(2^n)
  - メモリスタックのOOMの可能性

### バックトラック方法

#### フィボナッチ

```go
func fib(n int) {
    if n == 0 || n == 1 {
        return 1
    }

    return fib(n - 1) + fib(n - 2)
}
```

- 計算量
  - 関数のスタックが2つに分かれることによって、O(2^n)

- 無駄をなくしておく方法
  - メモ化
    - https://interviewcat.dev/p/coding-interviewcat/dynamic-programming-forward-backward#bc008c3ccb0b45cdaa32e94f4fb466cb
      - memo[n] = 結果で記録して


### 中立実装 メモ化

- 再帰を直感的で分かりやすいとした程で
  - デメリットの重複計算を避ける

```go
func fib(n int, memo map[int]int) int {
	if v, ok := memo[n]; ok {
		return v
	}
	if n < 2 {
		return 1
	}
	memo[n] = fib(n-1, memo) + fib(n-2, memo)
	return memo[n]
}

func main() {
	memo := make(map[int]int)
	fmt.Println(fib(10, memo))
}

```
package main

import "fmt"

// 考え方
// 全部の ()を作成する )) や ((もある
// validで不正パターンは除外する
func generateParenthesis(n int) []string {

	res := make([]string, 0)

	var valid func(string) bool
	valid = func(s string) bool {
		open := 0
		for _, c := range s {
			if c == '(' {
				open++
			} else {
				open--
			}
			if open < 0 {
				return false
			}
		}
		// + - 0になるとOKってことか -1 でも 1でも偏ってダメ
		return open == 0
	}

	var dfs func(string)
	dfs = func(s string) {
		if len(s) == n*2 {

			if valid(s) {
				res = append(res, s)

			}
			return
		}
		dfs(s + "(")
		dfs(s + ")")
	}

	dfs("")
	return res
}

func main() {

	fmt.Println(generateParenthesis(1))
}

// bfs 幅 の場合はキューになる

// package main

// import "fmt"

// type Node struct {
// 	s     string // 現在の文字列
// 	open  int    // "(" の数
// 	close int    // ")" の数
// }

// func generateParenthesisBFS(n int) []string {
// 	var res []string

// 	queue := []Node{{s: "", open: 0, close: 0}}

// 	for len(queue) > 0 {
// 		// 先頭の要素を取り出す（BFSなのでFIFO）
// 		node := queue[0]
// 		queue = queue[1:]

// 		// 完成したら結果に追加
//         // 括弧は n 組 → 長さ 2*n
// 		if len(node.s) == 2*n {
// 			res = append(res, node.s)
// 			continue
// 		}

// 		// "(" を追加できる条件
// 		if node.open < n {
// 			queue = append(queue, Node{
// 				s:     node.s + "(",
// 				open:  node.open + 1,
// 				close: node.close,
// 			})
// 		}

// 		// ")" を追加できる条件（"(" より多くはできない）
// 		if node.close < node.open {
// 			queue = append(queue, Node{
// 				s:     node.s + ")",
// 				open:  node.open,
// 				close: node.close + 1,
// 			})
// 		}
// 	}

// 	return res
// }

// func main() {
// 	fmt.Println(generateParenthesisBFS(3))
// }

// package main

// order は文字列の順番
// s string 並び替えしたい文字列
// return 並び替えた文字列

// 基本ば全探索でガンあげる
// cba
// abcd
// 全探索で

// ロジック
// res []runeに一致したら入れていく
// 残りもあとで入れる

// func remove(index int, s string) string {
// 	i := index - 1
// 	res := ""
// 	front := s[0:i]
// 	back := s[i+1:]

// 	res += front
// 	res += back
// 	return res

// }

// func customSortString(order string, s string) string {
// 	res := make([]rune, len(s))
// 	target := []rune(s)
// 	for _, o := range order {
// 		for i, ss := range target {
// 			if o == ss {
// 				res = append(res, ss)
// 				re
// 			}

// 		}
// 	}

// }

// func main() {

// }

// gpt

// sliceを いい感じにできそう
// https://chatgpt.com/c/68ee4a44-73ec-8324-baa3-ea186d9673f5

package main

import (
	"fmt"
	"sort"
)

func customSortString(order, s string) string {
	// order の順序をマップ化
	rank := make(map[rune]int)
	// cba
	for i, ch := range order {
		// c 0
		// b 1
		// a 2
		rank[ch] = i
	}

	for k, v := range rank {
		fmt.Printf("ch %s i %d\n", string(k), v)
	}
	// s の文字をスライスに変換
	runes := []rune(s)

	// ソート
	sort.Slice(runes, func(i, j int) bool {
		ri, okI := rank[runes[i]]
		rj, okJ := rank[runes[j]]

		// order に含まれていない文字は最後に
		if !okI && !okJ {
			return runes[i] < runes[j] // 両方ないなら通常の辞書順でもいい
		}
		if !okI {
			return false // i は order にない → 後ろへ
		}
		if !okJ {
			return true // j は order にない → i を前へ
		}
		// 0 < 2 は true
		return ri < rj // rank の小さいほうが先
	})

	return string(runes)
}

func main() {
	fmt.Println(customSortString("cba", "abcd")) // cbad など
}

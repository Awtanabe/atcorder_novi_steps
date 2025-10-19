package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	res := make(map[string][]string)

	for _, s := range strs {
		sortedStr := sortedStr(s)
		res[sortedStr] = append(res[sortedStr], s)
	}

	result := [][]string{}

	for _, val := range res {
		result = append(result, val)
	}

	return result
}

func sortedStr(s string) string {
	res := []rune(s)

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	return string(res)
}

func main() {

	fmt.Println(groupAnagrams([]string{"act", "pots", "tops", "cat", "stop", "hat"}))
}

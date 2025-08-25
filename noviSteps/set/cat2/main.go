package main

// 組み合わせ
import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	s := make([]string, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&s[i])
	}

	m := make(map[string]struct{})
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i != j {
				m[s[i]+s[j]] = struct{}{}
			}
		}
	}
	fmt.Print(len(m))
}

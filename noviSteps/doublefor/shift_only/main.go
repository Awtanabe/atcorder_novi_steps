// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	in := bufio.NewReader(os.Stdin)

// 	var N int
// 	fmt.Fscan(in, &N)

// 	ans := int(1 << 60) // 十分大きい値
// 	for i := 0; i < N; i++ {
// 		var x int
// 		fmt.Fscan(in, &x)
// 		c := 0
// 		for x%2 == 0 {
// 			x /= 2
// 			c++
// 		}
// 		if c < ans {
// 			ans = c
// 		}
// 	}
// 	fmt.Println(ans)
// }

// 下記は Scanでデータを直で取ることによって、変換処理が不要
// ok のフラグも returnで終了することで終了可能
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	a := scanner.Text()
	N, _ := strconv.Atoi(a)

	scanner.Scan()
	lines := scanner.Text()

	parts := strings.Split(lines, " ")

	i_parts := make([]int, N)
	for i := range parts {
		a, _ := strconv.Atoi(parts[i])
		i_parts[i] = a
	}

	c := 0

	// ok := true

	for {
		for i := 0; i < N; i++ {
			if i_parts[i]%2 != 0 {
				fmt.Println(c) // ここに入る時終了だからbreakしないでreturnで良いのか
				return
				// ok = false
				// break
			}
			i_parts[i] = i_parts[i] / 2
		}
		// if ok {
		// 	c++
		// } else {
		// 	break
		// }
		c++
	}
	fmt.Println(c)
}

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

// 	var sum, sq int64
// 	for i := 0; i < N; i++ {
// 		var x int64
// 		fmt.Fscan(in, &x)
// 		sum += x
// 		sq += x * x
// 	}

// 	// (sum^2 - sum of squares) / 2
// 	ans := (sum*sum - sq) / 2
// 	fmt.Println(ans)
// }

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	var N int
// 	fmt.Scan(&N)

// 	data := make([]int, N)

// 	for i := 0; i < N; i++ {
// 		fmt.Scan(&data[i])
// 	}

// 	sort.Slice(data, func(i, j int) bool {
// 		return data[i] > data[j]
// 	})

// 	sum := 0
// 	for i := 0; i < N-1; i++ {
// 		sum += data[i] * data[i+1]

// 	}
// 	sum += data[0] * data[len(data)-1]

// 	fmt.Println(sum)
// }

// GTPの回答
package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	data := make([]int, N)
	for i := range data {
		fmt.Scan(&data[i])
	}

	// 降順ソート
	sort.Sort(sort.Reverse(sort.IntSlice(data)))

	sum := 0
	for i := 0; i < N; i++ {
		sum += data[i] * data[(i+1)%N]
	}

	fmt.Println(sum)
}

package main

import "fmt"

// func dailyTemperatures(temperatures []int) []int {

// 	res := make([]int, 0, len(temperatures))

// 	for i := 0; i < len(temperatures); i++ {
// 		flag := false
// 		for j := i + 1; j < len(temperatures); j++ {
// 			if temperatures[j] > temperatures[i] {
// 				res = append(res, j-i)
// 				flag = true
// 				break
// 			}
// 		}

// 		if !flag {
// 			res = append(res, 0)
// 		}
// 	}

// 	return res
// }

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := []int{} // インデックスを入れる

	for i, t := range temperatures {
		fmt.Printf("\nDay %d, Temp %d\n", i, t)
		fmt.Println("Start Stack:", stack)

		// 現在の気温がスタック上の日より高いなら処理
		for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
			stackInd := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop

			fmt.Println("---i", i, "stackInd", stackInd)
			res[stackInd] = i - stackInd

			fmt.Printf("  Found warmer day for index %d (%d℃) → wait = %d\n",
				stackInd, temperatures[stackInd], res[stackInd])
			fmt.Println("  Stack after pop:", stack)
		}

		// 今の日を追加
		stack = append(stack, i)
		fmt.Println("End Stack:", stack)
		fmt.Println("Result:", res)
	}

	return res
}

func main() {
	temps := []int{73, 72, 75}
	result := dailyTemperatures(temps)
	fmt.Println("\nFinal Answer:", result)
}

// func main() {

// 	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))
// }

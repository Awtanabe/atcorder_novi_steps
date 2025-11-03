package main

import (
	"fmt"
	"sort"
)

// func carFleet(target int, position []int, speed []int) int {

// 	type Car struct {
// 		Position int
// 		Time     int
// 	}

// 	cars := make([]Car, 0, len(position))

// 	for i := 0; i < len(position); i++ {
// 		t := (target - position[i]) / speed[i]
// 		cars = append(cars, Car{Position: position[i], Time: t})
// 	}

// 	sort.Slice(cars, func(i, j int) bool {
// 		return cars[i].Position > cars[j].Position
// 	})

// 	stack := make([]Car, 0, len(position))

// 	for i := 0; i < len(cars); i++ {
// 		stack = append(stack, cars[i])
// 		if len(stack) >= 2 && stack[len(stack)-1].Time <= stack[len(stack)-2].Time {
// 			stack = stack[:len(stack)-1]
// 		}
// 	}

// 	return len(stack)
// }

func carFleet(target int, position []int, speed []int) int {
	n := len(position)

	type Car struct {
		pos  int
		time float64
	}

	cars := make([]Car, n)
	for i := 0; i < n; i++ {
		t := float64(target-position[i]) / float64(speed[i])
		cars[i] = Car{pos: position[i], time: t}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})

	stack := []float64{}

	for _, c := range cars {
		// push
		stack = append(stack, c.time)

		// If the current car catches up (i.e., current time <= previous fleet time),
		// they merge, so remove the current time (pop),
		// effectively keeping the slower fleet leader time.
		if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1] // pop
		}
	}

	return len(stack)
}

func main() {
	// fmt.Println(carFleet(10, []int{4, 1, 0, 7}, []int{2, 2, 1, 1}))
	fmt.Println(carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}))

}

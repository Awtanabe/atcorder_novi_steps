package main

import (
	"fmt"
	"strconv"
	"strings"
)

// type Solution struct{}

// func (s *Solution) Encode(strs []string) string {
// 	res := make([]string, len(strs))

// 	for i := 0; i < len(strs); i++ {
// 		prefix := "#" + strconv.Itoa(len(strs[i]))
// 		res = append(res, prefix)
// 		res = append(res, strs[i])
// 	}

// 	return strings.Join(res, "")
// }

// func (s *Solution) Decode(encoded string) []string {
// 	res := []string{}

// 	// 23件が全体
// 	for i := 0; i < len(encoded); i++ {
// 		if encoded[i] == '#' {
// 			num := []rune{}
// 			i++
// 			for encoded[i] >= '0' && encoded[i] <= '9' {
// 				num = append(num, rune(encoded[i]))
// 				i++
// 			}
// 			fmt.Println("num", string(num))
// 			strLength, _ := strconv.Atoi(string(num))
// 			tmp := []rune{}
// 			i += 2
// 			fmt.Println("i", i)
// 			for j := 0; j < strLength; j++ {
// 				tmp = append(tmp, rune(encoded[i+j]))
// 			}
// 			res = append(res, string(tmp))
// 			i += strLength - 1
// 		}
// 	}
// 	return res
// }

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var sizes []string
	for _, str := range strs {
		sizes = append(sizes, strconv.Itoa(len(str)))
	}
	return strings.Join(sizes, ",") + "#" + strings.Join(strs, "")
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "" {
		return []string{}
	}
	parts := strings.SplitN(encoded, "#", 2)
	sizes := strings.Split(parts[0], ",")
	var res []string
	i := 0
	for _, sz := range sizes {
		if sz == "" {
			continue
		}
		length, _ := strconv.Atoi(sz)
		res = append(res, parts[1][i:i+length])
		i += length
	}
	return res
}

func main() {

	s := Solution{}

	fmt.Println(s.Encode([]string{"we", "say", ":", "yes", "!@#$%^&*()"}))
	fmt.Println(s.Decode("2,3,1,3,10#wesay:yes!@#$%^&*()"))

}

package main

import (
	"fmt"
)

func main() {
	changed := []int{1,3,4,2,6,8}
	fmt.Println(findOriginalArray(changed))
}

// func findOriginalArray(changed []int) []int {
// 	if len(changed)%2 != 0 {
// 		return []int{}
// 	}
// 	// sort.Ints(changed)
// 	res := []int{}
// 	hash := make(map[int]int)
// 	for _, v := range changed {
// 		if hash[v] > 0 {
// 			hash[v]--
// 			continue
// 		}
// 		hash[v*2]++
// 		res = append(res, v)
// 	}
// 	for _, v := range hash {
// 		if v > 0 {
// 			return []int{}
// 		}
// 	}
// 	return res
// }

func findOriginalArray(changed []int) []int {
	hash := make(map[int]int)
	if len(changed)%2 != 0 {
		return []int{}
	}
	res := []int{}
	for _, v := range changed {
		if hash[v] > 0 {
			hash[v]--
			continue
		}
		hash[v*2]++
		res = append(res, v)
	}
	for _, v := range hash {
		if v > 0 {
			return []int{}
		}
	}
	return res
}

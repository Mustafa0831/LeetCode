package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
}

// func sortArrayByParity(arr []int) []int {
// 	odd := make([]int, 0, len(arr))
// 	even := make([]int, 0, len(arr))
// 	for _, val := range arr {
// 		if val%2 == 0 {
// 			even = append(even, val)
// 		} else {
// 			odd = append(odd, val)
// 		}
// 	}
// 	even = append(even, odd...)
// 	return even
// }

func sortArrayByParity(arr []int)[]int{
	even:=make([]int, 0, len(arr))
	odd:=make([]int, 0, len(arr))
	for _,v:=range arr{
		if v%2==0{
			even = append(even, v)
		}else{
			odd = append(odd, v)
		}
	}
	even = append(even, odd...)
	return even
}

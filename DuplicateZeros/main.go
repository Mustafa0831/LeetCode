package main

import "fmt"

func main() {
	(duplicateZeros([]int{1, 0, 2, 3, 0, 4, 5, 0}))
	// duplicateZeros([]int{1, 0, 0, 0, 2, 3})
	// duplicateZeros([]int{0, 1, 7, 6, 0, 2, 0, 7})
	// duplicateZeros([]int{8, 4, 5, 0, 0, 0, 0, 7})
}

// func duplicateZeros(arr []int){
// 	oneZero:=make([]int, 0)
// 	fmt.Println(oneZero)

// 	for _, v:=range arr{
// 		if v==0{
// 			oneZero = append(oneZero, v,v)
// 			fmt.Println(oneZero)
// 		}else {
// 			oneZero = append(oneZero, v)
// 			fmt.Println(oneZero)

// 		}
// 	}
// 	for i:=0; i<len(arr); i++{
// 		arr[i]=oneZero[i]
// 	}
// 	fmt.Println(arr)
// }
func duplicateZeros(arr []int){
	oneZero:=make([]int, 0)
	for _, v:=range arr{
		if v==0{
			oneZero = append(oneZero, v,v)
		}else {
			oneZero = append(oneZero, v)
		}
	}
	for i:=0; i<len(arr); i++{
		arr[i]=oneZero[i]
	}
	fmt.Println(arr)
}
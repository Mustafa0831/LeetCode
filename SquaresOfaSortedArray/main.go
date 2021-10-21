package main

import "fmt"

func main() {
	fmt.Println(sortedSquares([]int{-4,-1,0,3,10}))
}

// func sortedSquares(nums []int) []int {
// 	result := make([]int, len(nums))
// 	startIndex, endIndex := 0, len(nums)-1
// 	remIndex := endIndex
// 	for remIndex >= 0 {
// 		startElSq, endElSq := nums[startIndex]*nums[startIndex], nums[endIndex]*nums[endIndex]
// 		if startElSq >= endElSq {
// 			result[remIndex] = startElSq
// 			// fmt.Println(result," ")
// 			startIndex++
// 		} else {
// 			result[remIndex] = endElSq
// 			// fmt.Println(result," ")
// 			endIndex--
// 		}
// 		remIndex--
// 	}
// 	return result
// }

func sortedSquares(nums []int)[]int{
	result:=make([]int,len(nums))
	startI,endI:=0, len(nums)-1
	remI:=endI
	for remI>=0{
		staElSq,endElSq:=nums[startI]*nums[startI],nums[endI]*nums[endI]
		if staElSq>=endElSq{
			result[remI]=staElSq
			startI++
		}else{
			result[remI]=endElSq
			endI--
		}
		remI--
	}
	return result
}
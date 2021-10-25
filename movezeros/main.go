package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println("Solution 1 : ", nums)
}

// func moveZeroes(nums []int) {
// 	numZero := 0
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] == 0 {
// 			numZero++
// 		} else {
// 			nums[i-numZero] = nums[i]
// 		}
// 	}
// 	for numZero > 0 {
// 		nums[len(nums)-numZero] = 0
// 		numZero--
// 	}
// 	return
// }
func moveZeroes(nums []int) {
	numZero:=0
	for i:=0; i<len(nums); i++{
		if nums[i]==0{
			numZero++
		}else {
			nums[i-numZero]= nums[i]
		}
	}
	for numZero >0{
		nums[len(nums)-numZero]=0
		numZero--
	}
	return
}

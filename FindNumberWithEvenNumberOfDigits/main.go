package main

import "fmt"

func main(){
	nums:=[]int{12, 345, 2, 6, 7896}
	fmt.Println(findNumbers(nums))
}

// func findNumbers(nums []int)int{
// 	var count int
// 	for _, num:=range nums{
// 		var numOfDigits int
// 		for num!= 0{
// 			numOfDigits++
// 			num/=10
// 		}
// 		if numOfDigits%2==0{
// 			count++
// 		}
// 	}
// 	return count
// }

func findNumbers(nums []int)int{
	var count int
	for _,num:=range nums {
		var numOfDigits int
		for num!=0{
			numOfDigits++
			num/=10
		}		
		if numOfDigits%2==0{
			count++
		}
	}
	return count
}
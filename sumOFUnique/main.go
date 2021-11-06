package main

import "fmt"

func main(){
	nums:= []int{1,2,3,4,5}
	fmt.Println(sumOfUnique(nums))
}

func sumOfUnique(nums []int) int{
	hash:= make(map[int]int)
	for _, value:= range nums{
		hash[value]++
	}
	fmt.Println(hash)
	sum:=0
	for key,value:=range hash{
		if value==1{
			sum+=key
		}
	}
	return sum
}
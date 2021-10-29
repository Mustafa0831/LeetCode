package main

import "fmt"

func main(){
	nums:=[]int{2,7,11,15}
	target:= 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int)[]int{
	hash:=make(map[int]int, len(nums))
	for i,num:=range nums{
		if j, ok:=hash[target-num]; ok{
			fmt.Println(hash)
			return []int{j, i}
		}
		hash[num]=i
	}
	return []int{}
}
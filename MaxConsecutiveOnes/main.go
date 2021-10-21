package main

import "fmt"

func main(){
	nums:= []int{1,1,0,1,1,1}
	fmt.Println(findMaxConsecutiveOnes(nums))
}

func findMaxConsecutiveOnes(nums []int)int{
	countMax, find :=0,0
	for _,v:=range nums{
		if v==1 {
			countMax++
			if countMax>find{
				find = countMax
			}
		}else {
			countMax=0
		}
	}
	return find
}
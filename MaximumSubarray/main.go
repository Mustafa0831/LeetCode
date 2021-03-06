package main

import "fmt"

func main(){
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int)int{
	max,sum:=nums[0],nums[0]
	for _,v:=range nums[1:]{
		if sum >0{
			sum+=v
		}else {
			sum=v
		}
		if max <sum{
			max = sum
		}
	}
	return max
}
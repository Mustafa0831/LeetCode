package main

import "fmt"

func main(){
	nums:=[]int{0,0,1,1,1,2,2,3,3,4}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int)int{
	count:=0
	for i:=1;i<len(nums);i++{
		if nums[count]!= nums[i]{
			count++
			if i!=count{
				nums[count],nums[i]= nums[i], nums[count]
				// fmt.Print(nums[count],nums[i])
			}
		} 
	}
	fmt.Println(nums)
	return count+1
}
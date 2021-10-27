package main

import "fmt"

func main() {
	fmt.Printf("%v\n", findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	// fmt.Printf("%v\n",findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	// fmt.Printf("%v\n",findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))

}

// func findDuplicates(nums []int)[]int{
// 	res := make([]int, 0)
// 	for i:=0; i<len(nums);{
// 		v :=nums[i]
// 		if v>0{
// 			if nums[v-1]<0{
// 				res = append(res, v)
// 				nums[i]=0
// 				i++
// 			}else {
// 				nums[i],nums[v-1]=nums[v-1], nums[i]
// 				nums[v-1]= -1
// 			}
// 		}else {
// 			i++
// 		}
// 	}
// 	return res
// }
func findDuplicates(nums []int) []int {
	numsCount := make(map[int]int)
	for _, num := range nums {
		numsCount[num]++
	}
	twice := []int{}
	for key, value:= range numsCount{
		if value ==2{
			twice = append(twice, key)
		}
	}
	return twice
}

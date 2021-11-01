package main

import "fmt"

func main(){
	// nums1:=[]int{1,1,3,2}
	// nums2:=[]int{2,3}
	// nums3:=[]int{3}
	nums1:=[]int{3,1}
	nums2:=[]int{2,3}
	nums3:=[]int{1,2}
	// nums1:=[]int{1,2,2}
	// nums2:=[]int{4,3,3}
	// nums3:=[]int{5}
	fmt.Println(twoOutOfThree(nums1, nums2, nums3))
}

// func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int)[]int{
// 	hash1:=make(map[int]int)
// 	hash2:=make(map[int]int)
// 	hash3:=make(map[int]int)
// 	for i:= range nums1{
// 		hash1[nums1[i]]++
// 	}
// 	for i:=range nums2{
// 		hash2[nums2[i]]++
// 	}
// 	for i:=range nums3{
// 		hash3[nums3[i]]++
// 	}
// 	hash:=make(map[int]int)
// 	result:=make([]int, 0)
// 	for i:=range nums1{
// 		if hash2[nums1[i]]>0{
// 			hash[nums1[i]]++
// 		}
// 		if hash3[nums1[i]]>0{
// 			hash[nums1[i]]++
// 		}
// 	}
// 	for i:=range nums2{
// 		if hash3[nums2[i]]>0{
// 			hash[nums2[i]]++
// 		}
// 	}
// 	for k:=range hash{
// 		result = append(result, k)
// 	}
// 	return result
// }
func twoOutOfThree(nums1,nums2,nums3 []int)[]int{
	hash1,hash2, hash3:=make(map[int]int),make(map[int]int),make(map[int]int)
	for i:=range nums1{
		hash1[nums1[i]]++
	}
	for i:=range nums2{
		hash2[nums2[i]]++
	}
	for i:=range nums3{
		hash3[nums3[i]]++
	}
	result:=make([]int, 0)
	hash:=make(map[int]int)
	for i:=range nums1{
		if value,ok:=hash1[nums2[i]]; ok&&value>0{
			hash[nums2[i]]++
		}
		if hash2[nums3[i]]>0{
			hash[nums3[i]]++
		}
	}
	for i:=range nums2{
		if hash3[nums2[i]]>0{
			hash[nums2[i]]++
		}
	}
	for key:=range hash{
		result = append(result, key)
	}
	return result
}
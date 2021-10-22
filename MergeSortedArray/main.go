package main

import (
	"fmt"
)

func main() {
	// nums1 := []int{1, 2, 3, 0, 0, 0}
	// // for i,v:=range nums1{
	// // 	fmt.Println(i,v)
	// // }
	// nums2 := []int{2, 5, 6}
	// merge(nums1, 3, nums2, 3)
	// fmt.Println(nums1)
	// m := 3
	// tempSl := []int{}
	// sl1 := []int{0, 2, 5, 0, 0, 0} //3

	// for i, n := range sl1 {
	// 	if i < m {
	// 		tempSl = append(tempSl, n)
	// 	}
	// }

	// sl2 := []int{2, 5, 6}
	// var c = append(tempSl, sl2...)
	// sort.Ints(c)

	// log.Println(c, "kurva")

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := 0
	for len(nums2) != 0 {
		if i == m {
			for j, v := range nums2 {
				nums1[i+j] = v
				fmt.Println(nums1, "nums1 (5)")
				fmt.Println(nums2, "nums2 (6)")

			}
			nums2 = []int{}
			fmt.Println(len(nums2), "len nums2")
			break
		}
		if nums1[i] > nums2[0] {
			num := nums2[0]
			fmt.Println(num, "num (1)")
			fmt.Println(nums1[0], "nums1[0] (2)")

			nums2 = nums2[1:]
			fmt.Println(nums2, "nums2 (3)")
			fmt.Println(nums1, "nums1 (++)")
			shift(nums1, i)

			nums1[i] = num
			fmt.Println(nums1, "nums1 (4)")
			m++
		}
		i++
	}
}

func shift(array []int, n int) {
	for i := len(array) - 2; i >= n; i-- {
		array[i+1] = array[i]
		fmt.Println(array, "array (*)")
		fmt.Println(array[i+1], "array[i+1] 123")
	}
	fmt.Println(array, "array ()")

}

//0 delete, mergerd := [arr1, ...arr2] = 1,2,3,2,5,6 - sort

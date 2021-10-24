package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := 0
	for len(nums2) != 0 {
		if nums1[i] > nums2[0] {
			num := nums2[0]
			nums2 = nums2[1:]
			shift(nums1, i)
			nums1[i] = num
			m++
		}
		if i == m {
			for j, v := range nums2 {
				nums1[i+j] = v
			}
			nums2 = []int{}
			break
		}
		i++
	}
}

func shift(nums1 []int, n int) {
	for i := len(nums1) - 2; i >= n; i-- {
		nums1[i+1] = nums1[i]
	}
}

package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums, 2))
}

func removeElement(nums []int, val int) int {
	res :=make([]int, 0)
	fmt.Println(res)

	for _,r :=range nums{
		if r!= val{
			res = append(res, r)
		}
	}
	fmt.Println(res)
	copy(nums, res)
	fmt.Println(nums)

	return len(res)
}

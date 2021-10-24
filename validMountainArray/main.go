package main

import "fmt"

func main() {
	fmt.Println(validMountainArray([]int{3, 5, 5}), "11111111")
	fmt.Println(validMountainArray([]int{1, 1, 2, 3, 4, 5, 4, 3, 2, 1}), "2222222222")
	fmt.Println(validMountainArray([]int{0, 3, 2, 1}), "33333")
}

func validMountainArray(slice []int) bool {
	i := 0
	for i<len(slice)-1 && slice[i]<slice[i+1]{
		i++
	}
	if i==0 ||i==len(slice)-1{
		return false
	}
	for i<len(slice)-1 && slice[i]>slice[i+1]{
		i++
	}
	return i==len(slice)-1
}

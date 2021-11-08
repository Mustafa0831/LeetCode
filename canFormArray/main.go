package main

import "fmt"

func main(){
	fmt.Println(canFormArray([]int{85}, [][]int{{85}}))
	// println(canFormArray([]int{15, 88}, [][]int{{88}, {15}}))
	// println(canFormArray([]int{49, 18, 16}, [][]int{{16, 18, 49}}))
	println(canFormArray([]int{91, 4, 64, 78}, [][]int{{78}, {4, 64}, {91}}))
}

func canFormArray(arr []int, pieces [][]int)bool{
	hash:=make(map[int][]int)
	for _, p:=range pieces{
		hash[p[0]]=p
	}
	j:=0
	for i:=0; i<len(arr);i+=j{
		num:=arr[i]
		piece, ok:=hash[num]
		if !ok{
			return false
		}
		for j=0; j<len(piece); j++{
			if arr[i+j]!=piece[j]{
				return false
			}
		}
		// i+=j
	}
	// println(num)
	fmt.Println(hash)
	return true
}
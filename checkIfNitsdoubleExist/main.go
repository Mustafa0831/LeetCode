package main

import "fmt"

func main() {
	arr := []int{10, 2, 6, 3}
	fmt.Println(checkIfExist(arr))
}

func checkIfExist(arr []int)bool{
	hash:=map[int]int{}
	for i:=range arr{
		if arr[i]%2==0{
			hash[arr[i]]=i
		}
	}
	if len(hash)==0{
		return false
	}
	for i:=range arr{
		temp, ok:=hash[arr[i]*2]
		if ok&&temp!=i{
			return true
		}
	}
	return false
}
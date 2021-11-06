package main

import "fmt"

func main() {
	arr := []string{"d","b","c","b","c","a"}
	k:=2
	fmt.Println(kthDistinct(arr, k))
}

// func kthDistinct(arr []string, k int)string{
// 	hash:= make(map[string]int)
// 	count:=1
// 	for _, val:= range arr{
// 		hash[val]++
// 	}
// 	for _, val:=range arr {
// 		if hash[val]==1{
// 			if count==k{
// 				return val
// 			}else if count<k{
// 				count++
// 			}
// 		}
// 	}
// 	return ""
// }
func kthDistinct(arr []string, k int)string{
	hashCount:=map[string]int{}
	for _,s:= range arr{
		hashCount[s]++
	}
	for _, s:=range arr{
		if hashCount[s]==1{
			if  k==1{
				return  s
			}
			k--
		}
	}
	return ""
}
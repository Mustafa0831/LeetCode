package main

import "fmt"

func main(){
	ransomNote:="aa"
	magazine:="aab"
	fmt.Println(canConstruct(ransomNote, magazine))
}

// func canConstruct(ransomNote string, magazine string)bool{
// 	temp := make([]int, 256)
// 	for _, v:=range magazine{
// 		temp[v]++
// 	}
// 	for _,v:=range ransomNote{
// 		temp[v]--
// 		if temp[v]<0{
// 			return false
// 		}
// 	}
// 	return true
// }

func canConstruct(ransom string, mag string)bool{
	hash:= make(map[rune]int, len(mag))
	for _, key:=range mag{
		hash[key]++
	}
	for _, key:= range ransom{
		if value, ok :=hash[key]; !ok || value<1{
			return false
		}
		hash[key]--
	}
	return true
}
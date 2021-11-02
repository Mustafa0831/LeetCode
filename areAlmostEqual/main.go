package main

import "fmt"

func main(){
	s1:="bank"
	s2:="kanb"
	fmt.Println(areAlmostEqual(s1,s2))
}

func areAlmostEqual(s1, s2 string)bool{
	if len(s1)!= len(s2){
		return false
	}
	hash:= make(map[rune]int)
	for _, letter:= range s1{
		hash[letter]++
	}
	for _, symbol:= range s2{
		if value,ok :=hash[symbol]; !ok && value<1{
			return false
		}
		hash[symbol]--
	}
	count:=0
	for i:=range s1{
		if s1[i]!=s2[i]{
			count++
		}
	}
	return count==2
}
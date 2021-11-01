package main

import "fmt"

func makeEqual(words []string) bool {
	hash:=make(map[rune]int)
	for _, word:=range words{
		for _, char:=range word{
			hash[char]++
		}
	}
	for _, numWord:= range hash{
		fmt.Println(numWord%len(words))
		if numWord != len(words){
			return false
		}
	}
	return true
}

func main(){
	words := []string{"abbab"}
	fmt.Println(makeEqual(words))
}
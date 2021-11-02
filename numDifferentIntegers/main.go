package main

import "fmt"

func main(){
	word:="a123bc34d8ef34"
	fmt.Println(numDifferentIntegers(word))
}

func numDifferentIntegers(word string)int{
	hash:= make(map[string]bool)
	for i:=0; i<len(word); i++{
		count:= word[i]
		if count== '0'&& i+1< len(word)&& isDigit(word[i+1]){
			continue
		}
		if !isDigit(count){
			continue
		}
		start:=i
		i++
		for i<len(word)&& isDigit(word[i]){
			i++
		}
		hash[word[start:i]]=true
	}
	return len(hash)
}
func isDigit(count byte)bool{
	return '0'<= count && count <='9'
}
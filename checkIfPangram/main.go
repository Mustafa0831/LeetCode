package main

import "fmt"

func main(){
	sentence:= "thequickbrownfoxjumpsoverthelazydog"
	fmt.Println(checkIfPangram(sentence))
}

func checkIfPangram(s string)bool{
	countCheck:=make([]int,26)
	for _, letter:= range s{
		countCheck[letter-'a']++
	}
	for i:=range countCheck{
		if countCheck[i]==0{
			return false
		}
	}
	return true
}
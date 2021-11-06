package main

import "fmt"

func main(){
	allowed:="ab"
	words:=[]string{"ad","bd","aaab","baa","badab"}
	fmt.Println(countConsistentStrings(allowed, words))
}

func countConsistentStrings(allowed string, words []string)int{
	hash:=make(map[rune]bool)
	var count=0
	for _, a:=range allowed{
		hash[a]=true// a b
	}
	for _, word:=range words{
		invalid:=false
		for _, w:=range word{
			if _, ok:=hash[w]; !ok{
				invalid=true
				fmt.Println(hash)
				break
			}
		}
		if !invalid{
			count++
		}
	}
	return count
}
// func countConsistentStrings(allowed string, words []string) int {
// 	alphabet := make([]bool, 26)
// 	for _, s := range allowed {
// 		alphabet[s-'a'] = true
// 	}

// 	count := 0
// 	for _, word := range words {
// 		for _, w := range word {
// 			if !alphabet[w-'a'] {
// 				count--
// 				break
// 			}
// 		}
// 		count++
// 	}
// 	return count
// }
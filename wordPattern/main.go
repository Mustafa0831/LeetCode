package main

import (
	"fmt"
	"strings"
)

func main() {
	var pattern string
	var str string
	// pattern, str = "abba", "dog cat cat dog"
	// fmt.Println(wordPattern(pattern, str))

	pattern, str = "abba", "dog cat cat fish"
	fmt.Println(wordPattern(pattern, str))

	// pattern, str = "aaaa", "dog cat cat dog"
	// fmt.Println(wordPattern(pattern, str))

	// pattern, str = "abba", "dog dog dog dog"
	// fmt.Println(wordPattern(pattern, str))

	// pattern, str = "abc", "dog cat dog"
	// fmt.Println(wordPattern(pattern, str))
}

func wordPattern(pattern, str string) bool {
	strs, hashP, hashS := strings.Split(str, " "), make(map[byte]string), make(map[string]byte)
	if len(pattern) != len(strs) {
		return false
	}
	for i:=0; i<len(pattern);i++{
		str:=strs[i]
		if value, ok:=hashP[pattern[i]]; ok && str != value{
			return false
		}
		if value, ok:=hashS[str]; ok&&pattern[i]!=value{
			return false
		}
		hashP[pattern[i]]=str
		// fmt.Println(hashP)
		hashS[str]=pattern[i]
		// fmt.Println(hashS)

	}
	return true
}

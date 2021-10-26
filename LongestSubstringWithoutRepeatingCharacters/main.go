package main

import "fmt"

func main() {
	s := "abcabcbb"
	fmt.Println(lengthLongestSubstring(s))
}

// func lengthLongestSubstring(s string) int {
// 	hash := make(map[rune]int, len(s))
// 	start, maxLen := 0, 0
// 	for i, char := range []rune(s) {
// 		if lastI, ok := hash[char]; ok && lastI >= start {
// 			start = lastI + 1

// 		}
// 		if i-start+1 > maxLen {
// 			maxLen = i - start + 1
// 		}
// 		hash[char] = i

// 	}
// 	return maxLen
// }

func lengthLongestSubstring(s string)int{
	hash:=make(map[rune]int, len(s))
	start, maxLen:=0,0
	for i, char:=range s{
		if lastI,ok:=hash[char]; ok&&lastI>=start{
			start = lastI+1
		}
		if i-start+1>maxLen{
			maxLen= i-start+1
		}
		hash[char]=i
	}
	return maxLen
}
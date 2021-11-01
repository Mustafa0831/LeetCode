package main

import "fmt"

func main() {
	s := "abacbc"
	s = "aaabb"
	fmt.Println(areOccurrencesEqual(s))
}

// func areOccurrencesEqual(s string) bool {
// 	hash := make(map[byte]int)
// 	for i := 0; i < len(s); i++ {
// 		hash[s[i]]++
// 	}
// 	setFreq:=hash[s[0]]
// 	for _, count:=range hash{
// 		if count!=setFreq{
// 			return false
// 		}
// 	}
// 	return true
// }
func areOccurrencesEqual(s string) bool{
	var record [26]int
	for _, v:=range s{
		record[v-'a']++
	}
	count:=0
	for i:=0; i<26; i++{
		if record[i]==0{
			continue
		}
		if count==0{
			count=record[i]
		}else {
			if count!= record[i]{
				return false
			}
		}
	}
	return true
}

package main

import "fmt"

type CountFreq struct {
	start, end int
	byteMp     map[byte]bool
}

func main() {
	s := "aabca"
	fmt.Println(countPalindromicSubsequence(s))
}
func countPalindromicSubsequence(s string) int {
	hash := map[byte]CountFreq{}
	for i:=0; i<len(s); i++{
		if v, found:=hash[s[i]]; found {
			v.end = i
			hash[s[i]]=v
		}else {
			hash[s[i]]=CountFreq{start: i,end :i, byteMp:map[byte]bool{}}
		}
	}
	for i:=0; i<len(s); i++{
		for _,v:=range hash{
			if v.start <i && i<v.end{
				v.byteMp[s[i]]= true
			}
		}
	}
	res:=0 
	for _, v:=range hash{
		res+=len(v.byteMp)
	}
	return res
}

// func countPalindromicSubsequence(s string) int {
// 	res := 0
// 	for i := 0; i < 26; i++ {
// 		left, right := 0, len(s)-1
// 		for left < len(s)-1 && s[left] != byte(i+'a') {
// 			left++
// 		}
// 		for right >= 0 && s[right] != byte(i+'a') {
// 			right--
// 		}
// 		hash := make(map[byte]int)
// 		for j := left + 1; j < right; j++ {
// 			hash[s[j]] = 1
// 		}
// 		fmt.Println(res)
// 		res += len(hash)
// 	}
// 	return res
// }

// func countPalindromicSubsequence(s string)int{
// 	count:=0
// 	var single [26]bool
// 	var double [26][26]bool
// 	var triple [26][26][26]bool
// 	for _,v:=range s{
// 		for i:=0;i<26;i++{
// 			if double[int(v-'a')][i]{
// 				if !triple[int(v-'a')][i][int(v-'a')]{
// 					triple[int(v-'a')][i][int(v-'a')]=true
// 					count++
// 				}
// 			}
// 		}
// 		for i:=0; i<26; i++{
// 			if single[i]{
// 				double[i][int(v-'a')]=true

// 			}
// 		}
// 		single[int(v-'a')]=true

// 	}
// 	return count
// }

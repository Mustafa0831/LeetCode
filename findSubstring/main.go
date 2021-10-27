package main

import "fmt"

func main(){
	// s:="barfoothefoobarman"
	// words:=[]string{"foo","bar"}
	s1:="wordgoodgoodgoodbestword"
	words1:=[]string{"word","good","best","word"}
	// fmt.Println(findSubstring(s,words))
	fmt.Println(findSubstring(s1,words1))
}

func findSubstring(s string, words []string)[]int{
	if len(words)==0{
		return []int{}
	}
	hash:= make(map[string]int)
	for _, word:=range words{
		hash[word]++
	}
	lenFword:=len(words[0])
	indices:=[]int{}
	for i:=range s{
		hash2:=make(map[string]int)
		count:=0
		for j:=i; j+lenFword<len(s); j+=lenFword{
			fragment:=s[j:j+lenFword]
			if _, ok:=hash[fragment];ok{
				hash2[fragment]++
				count++
				if hash2[fragment]>hash[fragment]{
					break
				}
			}else {
				break
			}
			if count==len(words){
				indices = append(indices, i)
				break
			}
		}
	}
	return indices
}

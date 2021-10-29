package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	// fmt.Println(isIsomorphic("aa", "ab")) // False
	// fmt.Println(isIsomorphic("ab", "aa")) // False
}

// func isIsomorphic(s, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}
// 	hash := make(map[byte]byte, len(s))
// 	visited := make(map[byte]bool, len(s))
// 	for i := 0; i < len(s); i++ {
// 		if key, ok := hash[s[i]]; ok {
// 			if key != t[i] {
// 				return false
// 			}
// 		} else {
// 			if _, ok := visited[t[i]]; ok {
// 				return false
// 			}
// 			hash[s[i]] = t[i]
// 			visited[t[i]] = true
// 		}
// 	}
// 	return true
// }

func isIsomorphic(s, t string)bool{
	if len(s)!= len(t){
		return false
	}
	hash:=make(map[byte]byte, len(s))
	visited:=make(map[byte]bool, len(s))
	for i:=0;i<len(s); i++{
		if key,ok:=hash[s[i]]; ok{
			if key!=t[i]{
				return false
			}
		}else {
			if _, ok:=visited[t[i]]; ok{
				return false
			}
			hash[s[i]]= t[i]
			visited[t[i]]= true
		}
	}
	return true
}

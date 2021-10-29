package main

import "fmt"

func main(){
	s:="aass"
	t:="asss"
	fmt.Println(isAnagram(s,t))
}

func isAnagram(s string, t string) bool {
	smap, tmap := make(map[rune]int), make(map[rune]int)

	if len(s) != len(t) {
		return false
	}

	for _, v := range []rune(s) {
		smap[v]++
	}
	for _, v := range []rune(t) {
		tmap[v]++
	}

	for k := range tmap {
		if smap[k] != tmap[k] {
			return false
		}
	}

	return true
}
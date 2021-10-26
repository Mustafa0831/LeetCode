package main

import "fmt"

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagram(strs))
}

func groupAnagram(strs []string) [][]string {
	hash := map[[26]byte][]string{}
	for _, str := range strs {
		key := [26]byte{}
		
		for _, char := range str {
			key[char-'a']++
		}
		hash[key] = append(hash[key], str)
	}
	result := [][]string{}
	for _, strs := range hash {
		result = append(result, strs)
	}
	return result
}

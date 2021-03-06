package main

import "fmt"

func main() {
	word1 := "abc"
	word2 := "bca"
	fmt.Println(closeStrings(word1, word2))
}

func closeStrings(word1, word2 string) bool {
	hash1 := make(map[byte]int)
	hash2 := make(map[byte]int)
	if len(word1) != len(word2) {
		return false
	}
	for i := 0; i < len(word1); i++ {
		hash1[word1[i]]++
		hash2[word2[i]]++
	}
	CntVal := make(map[int]int)
	value2 := make([]int, 0)
	for key, val1 := range hash1 {
		val2, ok := hash2[key]
		if !ok {
			return false
		}
		value2 = append(value2, val2)
		CntVal[val1]++
	}
	for _, v2 := range value2 {
		if CntVal[v2] <= 0 {
			return false
		}
		CntVal[v2]--
	}
	fmt.Println(value2, CntVal)
	return true
}

// func closeStrings(word1 string, word2 string) bool {
// 	if len(word1) != len(word2) {
// 		return false
// 	}
// 	// 2 freq maps for 2 words
// 	freq1 := make(map[byte]int)
// 	freq2 := make(map[byte]int)

// 	for i := 0; i < len(word1); i++ {
// 		freq1[word1[i]]++
// 		freq2[word2[i]]++
// 	}

// 	if len(freq1) != len(freq2) {
// 		return false
// 	}

// 	// Check if value array of word1 = value array of word2
// 	value1Map := make(map[int]int)
// 	value2 := make([]int, 0)
// 	for key, val1 := range freq1 {
// 		val2, ok := freq2[key]
// 		if !ok {
// 			return false
// 		}
// 		value2 = append(value2, val2)
// 		value1Map[val1]++
// 	}
// 	for i := 0; i < len(value2); i++ {
// 		if val1 := value1Map[value2[i]]; val1 <= 0 {
// 			return false
// 		}
// 		value1Map[value2[i]]--
// 	}
// 	return true

// }

// func mapper(word string) map[rune]int {
// 	result := make(map[rune]int)

// 	for _, r := range word {
// 		result[r]++
// 	}

// 	return result
// }

// func closeStrings(word1 string, word2 string) bool {
// 	if len(word1) != len(word2) {
// 		return false
// 	}

// 	w1 := mapper(word1)
// 	w2 := mapper(word2)

// 	for r, v := range w2 {
// 		if w1[r] == 0 {
// 			return false
// 		}

// 		if w1[r] == v {
// 			continue
// 		}

// 		for k, c := range w1 {
// 			if c == v {
// 				w1[k] = w1[r]
// 				w1[r] = c
// 				break
// 			}
// 		}

// 		return false
// 	}

// 	return true
// }

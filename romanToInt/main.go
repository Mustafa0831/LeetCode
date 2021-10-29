package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"))
	// romanToInt("IX")
	// romanToInt("LVIII")
}

func romanToInt(s string) int {
	hash := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	if len(s) == 1 {
		return hash[s]
	}
	pre := hash[string(s[0])]
	result := 0
	for i := 1; i < len(s); i++ {
		cur := hash[string(s[i])]
		if pre < cur {
			result -= pre
		} else {
			result += pre
		}
		pre = cur
	}
	result += pre
	return result
}

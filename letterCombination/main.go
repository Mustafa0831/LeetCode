package main

import "fmt"

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}

// func letterCombinations(digits string)[]string{
// 	getDigits:= [][]string{{"a", "b", "c"},{"d", "e", "f"},{"g", "h", "i"},{"j", "k", "l"},{"m", "n", "o"},{"p", "q", "r", "s"},{"t", "u", "v"},{"w", "x", "y", "z"},}
// 	if digits==""{
// 		return nil
// 	}
// 	strs := make([]string, 0)
// 	strs = append(strs, getDigits[digits[0]-'2']...)
// 	fmt.Println(getDigits[digits[0]-'2'])
// 	for _,digit:= range digits[1:]{
// 		str := getDigits[digit-'2']
// 		fmt.Println(str)
// 		temp:= make([]string, len(strs))
// 		copy(temp, strs)
// 		strs= []string{}
// 		for _, s:= range str{
// 			for _, t:=range temp{
// 				t+=s
// 				strs= append(strs, t)
// 			}
// 		}
// 	}
// 	return strs
// }

func letterCombinations(digits string) []string {
	getDigits := [][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}, {"j", "k", "l"}, {"m", "n", "o"}, {"p", "q", "r", "s"}, {"t", "u", "v"}, {"w", "x", "y", "z"}}
	if digits == "" {
		return nil
	}
	strs := make([]string, 0)
	strs = append(strs, getDigits[digits[0]-'2']...)
	for _, digit:=range digits[1:]{
		str:= getDigits[digit-'2']
		temp:=make([]string, len(strs))
		copy(temp, strs)
		strs=[]string{}
		for _, s:=range str{
			for _, t :=range temp{
				t+=s
				strs=append(strs, t)
			}
		}
	}
	return strs
}

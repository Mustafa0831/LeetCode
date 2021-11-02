package main

import "fmt"

func main() {
	s := "aababcabc"
	fmt.Println(countGoodSubstrings(s))
}

// func countGoodSubstrings(s string)int{
// 	count:=0
// 	for i:=0; i<len(s)-2; i++{
// 		if s[i]!= s[i+1] && s[i+1]!= s[i+2] && s[i+2]!=s[i]{
// 			count++
// 		}
// 	}
// 	return count
// }

//Window ...
type Window struct {
	data      []rune
	goodCount int
	left      int
	right     int
}

func countGoodSubstrings(s string) int {
	if len(s) < 3 {
		return 0
	}
	w := ConstructWindow(s)
	for w.CanMoveRight() {
		w.MoveRight()
	}
	return w.goodCount
}

//ConstructWindow ...
func ConstructWindow(s string) *Window {
	w := Window{}
	w.data = []rune(s)
	w.left = 0
	w.right = 2
	w.CheckForGoodSubstring()
	return &w
}

//CanMoveRight ...
func (w *Window) CanMoveRight() bool {
	return w.right < len(w.data)-1
}

//MoveRight ...
func (w *Window) MoveRight() {
	w.left++
	w.right++
	w.CheckForGoodSubstring()
}

//CheckForGoodSubstring ...
func (w *Window) CheckForGoodSubstring() {
	if w.ContainsGoodSubstring() {
		w.goodCount++
		// fmt.Println(w.goodCount)
	}
}

//ContainsGoodSubstring ...
func (w *Window) ContainsGoodSubstring() bool {
	hash := map[rune]int{}
	for _, r := range w.windowedData() {
		hash[r]++
		if hash[r] > 1 {
			return false
		}
	}
	return true
}

func (w *Window) windowedData() []rune {
	return w.data[w.left : w.right+1]
}

// func countGoodSubstrings(s string) int {
// 	goodSubStrings := 0
// 	subStr := map[string]int{}
// 	for wStart, wEnd, wSize := 0, 0, 0; wEnd < len(s); {
// 		for wSize < 3 && wEnd < len(s) {
// 			str := string(s[wEnd])
// 			if _, exists := subStr[str]; !exists {
// 				subStr[str] = 0
// 			}
// 			subStr[str]++
// 			wSize++
// 			wEnd++
// 		}
// 		if len(subStr) == 3 {
// 			goodSubStrings++
// 		}
// 		str := string(s[wStart])
// 		subStr[str]--
// 		if subStr[str] == 0 {
// 			delete(subStr, str)
// 		}
// 		wStart++
// 		wSize--
// 	}
// 	return goodSubStrings
// }

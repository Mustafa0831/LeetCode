package main

import "fmt"

func main(){
	// fmt.Println(getHint("1807", "7810"))
	fmt.Println(getHint("1123", "0111"))
	// fmt.Println(getHint("1", "0"))
	// fmt.Println(getHint("1", "1"))
}

func getHint(secret string, guess string)string{
	hash:=make(map[byte]int)
	var bullCnt,cowCnt int
	for i:=0; i<len(secret); i++{
		if secret[i]==guess[i]{
			bullCnt++
		}
		hash[secret[i]]++
	}
	for i:=0; i<len(guess); i++{
		if value, ok:=hash[guess[i]]; ok&&value>0{
			cowCnt++
			hash[guess[i]]--
		}
	}
	cowCnt-=bullCnt
	return fmt.Sprintf("%dA%dB", bullCnt, cowCnt)
}
package main

import "fmt"

func main(){
	s:= "dfa1542321afd"
	fmt.Println(secondHighest(s))
}

func secondHighest(s string)int{
	// var counts [10]int	
	counts:=[10]int{}
	for _,letter:=range s{
		if letter>='0' && letter<='9'{
			counts[letter-'0']=1
		}
	}
	sign:= false
	for i:=9; i>=0; i--{
		if counts[i]==1{
			if sign {
				return i
			}
			sign = true
		}
	}
	return -1
}
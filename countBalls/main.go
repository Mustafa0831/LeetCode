package main

import "fmt"

func main(){
	l:=19
	h:=28
	fmt.Println(countBalls(l,h))
}

func countBalls(lowLimit, highLimit int)int{
	var max int
	count:= make(map[int]int)
	fmt.Println(1/10)
	for i:=lowLimit; i<=highLimit; i++{
		var sum int
		tmp := i// 5
		for tmp!=0{
			sum+=tmp%10//5
			tmp /=10//0
		}
		count[sum]++//1
		// fmt.Println(count[sum])
		if count[sum]>max{
			max=count[sum]// 1
		}
	}
	return max
}
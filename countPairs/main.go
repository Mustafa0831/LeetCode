package main

import "fmt"

func main() {
	fmt.Println(countPairs([]int{1, 3, 5, 7, 9}))
}

// func countPairs(deliciousness []int) int {
// 				// 1,   2,     4
// 	twoExp := []int{1, 1 << 1, 1 << 2, 1 << 3, 1 << 4, 1 << 5, 1 << 6, 1 << 7, 1 << 8, 1 << 9, 1 << 10, 1 << 11, 1 << 12, 1 << 13, 1 << 14, 1 << 15,
// 		1 << 16, 1 << 17, 1 << 18, 1 << 19, 1 << 20, 1 << 21}
// 		hash:=map[int]int{}
// 		for _, v:=range deliciousness{
// 			hash[v]++// 1, 3, 5, 7, 9
// 		}
// 		result:=0
// 		for key,value:=range hash{
// 			for _, powerOf2:=range twoExp{
// 				if powerOf2<key{// 1 <1
// 					continue
// 				}
// 				odd:= powerOf2-key // 0
// 				if odd<key{ // 0 < 1
// 					continue
// 				}
// 				if odd!= key{ //0!= 1
// 					result += value*hash[odd] // 1*1
// 				}else {
// 					result+=value*(value-1)/2
// 				}
// 				result %=1000000007
// 			}
// 		}
// 	return  result
// }

func countPairs(deliciousness []int)int{
	m :=make(map[int]int)
	ctn:=0
	list:=make([]int,22)
	for i:=0; i<22;i++{
		list[i]= 1<<i
	}
	for _, d:=range deliciousness{
		for _, l:=range list{
			if l>=d{
				ctn+=m[l-d]
			}
		}
		m[d]++
	}
	return ctn%(1e9+7)
}

package main

import "fmt"

func main(){
	adjacentPairs:= [][]int{{2,1},{3,4},{3,2}}
	fmt.Println(restoreArray(adjacentPairs))
}

func restoreArray(adjacentPairs [][]int)[]int{
	hash:= make(map[int][]int)
	for _, link:=range adjacentPairs{
		hash[link[0]]=append(hash[link[0]], link[1])// 2:[1], 3:[4 2]
		hash[link[1]]=append(hash[link[1]], link[0])// 1:[2], 4:[3], 2:[3]
		// 1:[2] 2:[1 3] 3:[4 2] 4:[3]
	}
	result:=make([]int, len(adjacentPairs)+1)
	for k, v:=range hash{
		if len(v)==1{
			result[0]=k
			// 1, 4
			result[1]=v[0]
			// 2, 3
			break
		}
	}
	for i:=2; i<len(adjacentPairs)+1; i++{// 2 3
		// fmt.Println(result[1])
		if result[i-2]!=hash[result[i-1]][0]{// i=3 ->  2 != 
			result[i]=hash[result[i-1]][0]
		}else {
			result[i]=hash[result[i-1]][1]// i=2 ->  4 == 4 result[2]= 2
		}
	}
	return result
}
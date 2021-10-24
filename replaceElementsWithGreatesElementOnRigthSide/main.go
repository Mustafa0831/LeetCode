package main

import "fmt"

func main() {
	arr :=[]int{17,18,5,4,6,1}
	fmt.Println(replaceElements(arr))
	// for {
	// 	var numsSize, inputNum int
	// 	fmt.Scan(&numsSize)
	// 	var numsList, resultList []int
	// 	for j:=0; j<numsSize; j++{
	// 		fmt.Scan(&inputNum)
	// 		numsList=append(numsList, inputNum)
	// 	}
	// 	resultList =replaceElements(numsList)
	// 	fmt.Println(resultList)
	// }
}

func replaceElements(arr []int) []int{
	maxnum:=-1
	for x:=len(arr)-1; x>=0; x--{
		temp:=arr[x]
		arr[x]=maxnum
		if temp >maxnum{
			maxnum= temp
		}
	}
	return arr
}

// func replaceElements(arr []int) []int{
// 	newArr:=make([]int, len(arr))
// 	for i:=0;i<len(arr)-1;i++{
// 		a:=0
// 		a=arr[i+1]
// 		for j:=i+1; j<len(arr); j++{
// 			if arr[j]>a{
// 				a=arr[j]
// 			}
// 			newArr[i]=a
// 		}
// 	}
// 	newArr[len(arr)-1]= -1
// 	return newArr
// }
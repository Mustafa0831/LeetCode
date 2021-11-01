package main

import "fmt"

func main() {
	fmt.Println(isHappy(19))
	// fmt.Println(isHappy(100))
	// fmt.Println(isHappy(1))
	// fmt.Println(isHappy(0))
	fmt.Println(isHappy(2))
}

func isHappy(n int) bool {
	var exist = make(map[int]int)
	for n != 1 {
		var tmp = n
		var sum = 0
		for tmp != 0 {
			var tmp2 = tmp % 10
			sum += tmp2 * tmp2
			tmp /= 10
		}
		var _, ok = exist[sum]
		if ok {
			return false
		}
		exist[sum] = 1
		n = sum
	}
	return true
}

// func isHappy(num int) bool {
// 	hash := make(map[int]int)
// 	sum := 0
// 	for num != 1 {
// 		temp := num
// 		for temp != 0 {
// 			temp2 := temp % 10
// 			sum += temp2 * temp2
// 			temp /= 10
// 		}
// 		var _, ok = hash[sum]
// 		if ok {
// 			return false
// 		}
// 		hash[sum] = 1
// 		num = sum
// 	}
// 	return true
// }

package main

import "fmt"

func main() {
	x := 121
	// x := 123
	result := isPalindrome(x)
	fmt.Println(result)
}
func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	temp2 := x
	temp := 0

	for {
		temp += temp2 % 10
		temp2 = temp2 / 10
		if temp2 <= 0 {
			break
		}
		temp = temp * 10
	}

	return temp == x
}

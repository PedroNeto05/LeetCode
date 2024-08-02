package main

import (
	"fmt"
)

func main() {
	result := romanToInt("MCMXCIV") //1994
	fmt.Println(result)
}
func romanToInt(s string) int {

	num := 0

	nums := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i := 0; i < len(s)-1; i++ {

		if nums[string(s[i])] < nums[string(s[i+1])] {
			num -= nums[string(s[i])]
			continue
		}
		num += nums[string(s[i])]

	}

	num += nums[string(s[len(s)-1])]

	return num
}

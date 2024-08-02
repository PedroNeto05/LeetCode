package main

import (
	"fmt"
)

func main() {
	strs := []string{"dog", "racecar", "car"}

	prefix := longestCommonPrefix(strs)
	fmt.Println(prefix)
}

func longestCommonPrefix(strs []string) string {
	prefix := ""
	minStr := findMinStringInSlice(strs)

mainLoop:
	for i := 0; i < len(minStr); i++ {
		possiblePrefix := string(minStr[i])

		for j := 0; j < len(strs); j++ {

			if string(strs[j][i]) != possiblePrefix {
				break mainLoop
			}

		}

		prefix += possiblePrefix

	}

	return prefix
}

func findMinStringInSlice(strs []string) string {
	min := strs[0]

	for _, v := range strs {

		if len(v) < len(min) {
			min = v
		}

	}

	return min
}

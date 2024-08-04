package main

import (
	"fmt"
)

type Stack struct {
	elements []string
}

func (s *Stack) Push(value string) {
	s.elements = append(s.elements, value)
}

func (s *Stack) Pop() (string, bool) {
	if len(s.elements) == 0 {
		return "", false
	}
	topIndex := len(s.elements) - 1
	topElement := s.elements[topIndex]
	s.elements = s.elements[:topIndex]
	return topElement, true
}

func (s *Stack) Peek() (string, bool) {
	if len(s.elements) == 0 {
		return "", false
	}
	topIndex := len(s.elements) - 1
	topElement := s.elements[topIndex]
	return topElement, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func main() {
	s := "([)]" //false
	result := isValid(s)
	fmt.Println(result)
}

func isValid(s string) bool { //83/98 -> minha solução
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := []rune{}

	for _, r := range s {
		if _, ok := pairs[r]; ok {
			stack = append(stack, r)
		} else if len(stack) == 0 || pairs[stack[len(stack)-1]] != r {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

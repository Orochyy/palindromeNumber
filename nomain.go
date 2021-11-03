package main

import (
	"fmt"
	"strconv"
)

func isPalindrome2(x int) bool {
	s := strconv.Itoa(x)
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true
}

func main() {
	x := 22
	q := isPalindrome2(x)
	fmt.Println(q)
}

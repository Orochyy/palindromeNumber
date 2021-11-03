package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	t := strconv.Itoa(x)
	return t == Reverse(t)
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func main() {
	x := 22
	q := isPalindrome(x)
	fmt.Println(q)
}

//func main() {
//	var number, remainder, temp int
//	var reverse int = 0
//
//	fmt.Print("Enter any positive num : ")
//	fmt.Scan(&number)
//
//	temp = number
//
//	for {
//		remainder = number % 10
//		reverse = reverse*10 + remainder
//		number /= 10
//
//		if number == 0 {
//			break
//		}
//	}
//
//	if temp == reverse {
//		fmt.Printf("%d is a Palindrome", temp)
//	} else {
//		fmt.Printf("%d is not a Palindrome", temp)
//	}
//
//}

package main

import (
	"fmt"
	"strings"
)

func main() {

	var str string
	fmt.Print("Enter a string")
	fmt.Scan(&str)

	if isPalin(strings.ToUpper(str)) == true {
		fmt.Print(str, " is a palindrome.")
	} else {
		fmt.Print(str, " is not a palindrome.")
	}
}

func isPalindrome(word string) bool {
	// 	l := 0
	// 	h := len(char) - 1

	// 	for h > l {
	// 		if (string(l) != string[h]); BadExpr {
	// 			return
	// 		}
	// 	}

	for i := 0; i <= len(word)/2; i++ {
		if word[i] != word[len(word)-i-1] {
			return false
		}

	}

	return true
}

func Reverse(s string) (r string) {
	for _, v := range s {
		r = string(v) + r
	}
	return
}

func isPalin(str string) interface{} {
	if str == Reverse(str) {
		return true
	}
	return false
}

package main

import "fmt"

func main() {
	wrd := distincChar("dog")
	fmt.Println(wrd)

}

func distincChar(word string) bool {
	if len(word) > 128 {
		return false
	}

	var char_set [128]bool
	for i := 0; i < len(word); i++ {
		val := word[i]
		// char_set[i] := word[i].char
		if char_set[val] {
			return false
		}
		char_set[val] = true
	}

	return true
}

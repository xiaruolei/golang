package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	/*
	fields, split, join
	contains, index
	tolower, toupper
	trim, trimright, trimleft
	 */
	//fmt.Println(s)
	lastoccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		//fmt.Printf("(%d %c)\n", i, ch)
		lastI, ok := lastoccurred[ch]
		if ok && lastI >= start {
			start = lastoccurred[ch] + 1
		}

		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}

		lastoccurred[ch]  = i
	}

	return maxLength
}
func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("b"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdef"))
}

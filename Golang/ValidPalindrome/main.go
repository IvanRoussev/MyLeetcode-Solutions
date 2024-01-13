package main

import (
	"log"
	"unicode"
)

func isPalindrome(s string) bool {
	var alphaList []rune
	for _, c := range s {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
			continue
		} else {
			c = unicode.ToLower(c)
			log.Println(string(c))
			alphaList = append(alphaList, c)
		}
	}
	log.Println(alphaList)
	left, right := 0, len(alphaList)-1

	for left < right {
		if alphaList[left] != alphaList[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	s := "race a car"
	res := isPalindrome(s)
	println(res)

}

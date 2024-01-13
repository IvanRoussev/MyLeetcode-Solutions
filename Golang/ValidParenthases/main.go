package main

import "log"

func validParentheses(s string) bool {
	var stack []string

	closeToOpen := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}

	for idx, char := range s {
		log.Printf("Index: %v, Parantheses: %v", idx, string(char))
		val, ok := closeToOpen[string(char)]
		if ok {
			if len(stack) != 0 && stack[len(stack)-1] == val {
				log.Printf("index %v deleting Such item from stack: %v", idx, val)
				stack = stack[:len(stack)-1]
				log.Printf("Item deleted: %v, New stack: %v", val, stack)
			} else {
				return false
			}
		} else {
			stack = append(stack, string(char))
			log.Printf("Index: %v, Stack: %v", idx, stack)
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func main() {
	s := "()[]{}"
	valid := validParentheses(s)
	log.Println(valid)
}

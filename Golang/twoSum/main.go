package main

import "log"

func twoSum(numbers []int, target int) []int {
	leftPointer := 0
	rightPointer := len(numbers) - 1

	var num int
	for leftPointer != rightPointer {
		num = numbers[leftPointer] + numbers[rightPointer]
		if num == target {
			return []int{leftPointer + 1, rightPointer + 1}
		} else if num > target {
			rightPointer--
		} else if num < target {
			leftPointer++
		}
	}
	return []int{}
}

func main() {
	list := []int{0, 0, 3, 4}
	target := 0
	res := twoSum(list, target)
	log.Println(res)
}

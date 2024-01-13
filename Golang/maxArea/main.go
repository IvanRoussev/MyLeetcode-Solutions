package main

import (
	"log"
)

func maxArea(height []int) int {
	leftPointer := 0
	rightPointer := len(height) - 1

	maxArea := 0
	for leftPointer < rightPointer {
		area := (rightPointer - leftPointer) * min(height[leftPointer], height[rightPointer])
		if area > maxArea {
			maxArea = area
		}
		if height[leftPointer] < height[rightPointer] {
			leftPointer++
		} else {
			rightPointer--
		}
	}
	return maxArea
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	resp := maxArea(height)
	log.Println(resp)
}

package main

import (
	"log"
)

func trap(height []int) int {

	if height == nil {
		return 0
	}

	maxLeftIdx := 0
	maxRightIdx := len(height) - 1
	maxLeftVal := height[maxLeftIdx]
	maxRightVal := height[maxRightIdx]

	trapped := 0

	for maxLeftIdx < maxRightIdx {
		if maxLeftVal < maxRightVal {
			maxLeftIdx += 1
			maxLeftVal = max(maxLeftVal, height[maxLeftIdx])
			trapped += maxLeftVal - height[maxLeftIdx]
		} else {
			maxRightIdx -= 1
			maxRightVal = max(maxRightVal, height[maxRightIdx])
			trapped += maxRightVal - height[maxRightIdx]
		}
	}
	return trapped
}

func main() {
	water := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	resp := trap(water)
	log.Println(resp)

}

//func trap(height []int) int {
//	maxLeftIdx := 0
//	maxRightIdx := len(height) - 1
//
//	trapped := 0
//
//	maxLeft := height[maxLeftIdx]
//	maxRight := height[maxRightIdx]
//	for idx, num := range height {
//		if idx == 0 {
//			maxLeft
//		}
//		avail := min(height[maxLeftIdx], height[maxRightIdx]) - num
//		log.Println(avail)
//		if avail < 0 {
//			avail = 0
//		}
//		trapped += avail
//
//		if height[maxLeftIdx] > height[maxRightIdx] {
//			maxRightIdx--
//		} else {
//			maxLeftIdx++
//		}
//	}
//	return trapped
//}
//
//func main() {
//	water := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
//	resp := trap(water)
//	log.Println(resp)
//}

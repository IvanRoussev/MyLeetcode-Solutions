package main

import "log"

func topKFrequent(nums []int, k int) []int {
	hashmap := make(map[int]int)

	for _, num := range nums {
		hashmap[num]++
	}
	log.Println(hashmap)

	var frequent []int
	for key, value := range hashmap {
		if len(frequent) < k {
			frequent = append(frequent, key)
		} else {
			for index := range frequent {
				if hashmap[frequent[index]] < value {
					frequent[index] = key
					break
				}
			}
		}
	}

	return frequent
}

func main() {
	nums := []int{5, -3, 9, 1, 7, 7, 9, 10, 2, 2, 10, 10, 3, -1, 3, 7, -9, -1, 3, 3}
	resp := topKFrequent(nums, 2)
	log.Println(resp)
}

package main

import "log"

func groupAnagram(wordsList []string) [][]string {
	hashmap := make(map[[26]int][]string)

	for _, word := range wordsList {
		var count [26]int
		for _, letter := range word {
			count[letter-'a']++
		}
		hashmap[count] = append(hashmap[count], word)
		log.Println(hashmap[count])

	}
	result := make([][]string, len(hashmap))

	index := 0
	for _, v := range hashmap {
		result[index] = v
		index++
	}
	return result
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	output := groupAnagram(strs)
	log.Println(output)

}

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello, World!")
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	// [["bat"],["nat","tan"],["ate","eat","tea"]] // output
	result := groupAnagrams(strs)
	fmt.Println(result)
}

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)
	for _, str := range strs {
		chars := []byte(str)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})

		key := string(chars)

		anagramMap[key] = append(anagramMap[key], str)
	}
	result := [][]string{}

	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result

}

func IsAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	count := make(map[rune]int)
	// Count the frequency of each character in s1
	// count = map[rune]int{'a': 1, 'e': 1, 't': 1}
	for _, c := range s1 {
		count[c]++
	}
	// Check if s2 has the same frequency of characters
	for _, c := range s2 {
		// If the character is not in the map or its count is zero, then s2 is not an anagram of s1
		if count[c] == 0 {
			return false
		}
		// count[c] is decremented to account for the character in s2
		count[c]--
	}
	// If all counts are zero, then s1 and s2 are anagrams
	return true
}

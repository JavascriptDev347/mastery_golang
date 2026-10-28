package main

import (
	"fmt"
	"sort"
)

func main() {

	// contains duplicate
	nums := []int{1, 2, 3, 4, 1}
	re := containsDuplicate2(nums)
	fmt.Println("Result:", re)
	// fmt.Println("Hello, World!")
	// strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	// // [["bat"],["nat","tan"],["ate","eat","tea"]] // output
	// result := groupAnagrams(strs)
	// fmt.Println(result)
}

// nums1 = [1,2,2,1], nums2 = [2,2]
func intersection(nums1 []int, nums2 []int) []int {
	count := make(map[int]struct{}, len(nums1))
	result := make([]int, 0)
	for _, v := range nums1 {
		count[v] = struct{}{}
	}

	for _, v := range nums2 {
		if _, ok := count[v]; ok {
			result = append(result, v)
			delete(count, v)
			// if !slices.Contains(result, v) {
			// 	result = append(result, v)
			// }
		}
	}
	return result
}

// 217 contains duplicate
func containsDuplicate(nums []int) bool {

	m := map[int]int{}
	for _, v := range nums {
		_, ok := m[v]
		if ok {
			return true
		}
		m[v]++
	}

	return false
}

func containsDuplicate2(nums []int) bool {
	seen := make(map[int]struct{})

	for _, v := range nums {
		_, ok := seen[v]
		if ok {
			return true
		}

		seen[v] = struct{}{}
	}

	return false
}

// 242 valid anagram
func isAnagram(s1, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	count := map[rune]int{}

	for _, v := range s1 {
		count[v]++
	}

	for _, v := range s2 {
		if count[v] == 0 {
			return false
		}

		count[v]--
	}

	return true
}
func IsAnagram2(s1, s2 string) bool {
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

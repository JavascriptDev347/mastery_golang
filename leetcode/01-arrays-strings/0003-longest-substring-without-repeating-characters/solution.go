// Package longestsubstring solves LeetCode 3. Longest Substring Without
// Repeating Characters.
package longestsubstring

// 1. What is being asked ?
// Find the length of the longest contiguous portion (substring) of s that contains no character more than once.
// 2. What is the input ?
// s is a string
// Important section about the input is that string will be empty. If s is empty, return 0.
//3. What is the output ?
// A single integer: the length of the longest substring without any repeated character
// We need only the length, not the actual substring itself
// 4. Step-by-step example analysis
// Example 1: s = "abcabcbb"
// "a" => 1 👍
// "b" => 2 👍
// "c" => 3 👍
// "a" => 3 repeated 👎
// "b" => 3 👍
// "c" => 3 repeated 👎
// "b" => 3 👍
// "b" => 3 repeated 👎
// "b" => 3 👍
// The answer is 3

// Approach 1: Brute force
// Algorithm (step-by-step):
//
//	1.outer loop: start index i=0 to n-1
//	2.inner loop: start index j=i to n-1
//	3.for substring s[i...j, check the all characters are unique
//	4.if unique, update max length maxLen = max(maxLen, j-i+1)
//	5.If not unique → break inner loop (any extension will also repeat)
func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	n := len(s)
	for i := 0; i < n; i++ {
		seen := make(map[byte]bool)
		for j := i; j < n; j++ {
			if seen[s[j]] {
				break //duplicate found
			}
			seen[s[j]] = true
			if j-i+1 > maxLen {
				maxLen = j - i + 1
			}
		}
	}
	return maxLen
}

// Approach 2: Sliding window
// What is the problem with Brute force?
// We're re-examining the same characters over and over. Optimization: maintain a moving window [left, right]. Expand right; when a duplicate appears, shrink from the left until the duplicate is removed.
// Algorithm (step-by-step):
//  1. left = 0, right = 0, seen = set()
//  2. For right from 0 to n-1: a. While s[right] in seen: remove s[left], increment left b. Add s[right] to seen
//     c. maxLen = max(maxLen, right - left + 1)
//  3. Return maxLen
func lengthOfLongestSubstring2(s string) int {
	seen := make(map[byte]bool)
	left, maxLen := 0, 0
	for right := 0; right < len(s); right++ {
		for seen[s[right]] {
			delete(seen, s[left])
			left++
		}
		seen[s[right]] = true
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

// Approach 3: Optimal (Sliding Window with Last-Seen Index Map)
// What is the problem with the previous approach?
// When a duplicate is found, Approach 2 shrinks left one step at a time until the duplicate is removed. For a string like "aXXXXXa", finding the second 'a' forces left to step through all the Xs. We can do better: jump left directly to lastSeen['a'] + 1.
// Optimization idea: Replace the set with a map that stores the most recent index of each character. When a duplicate ch is found at right, set left = max(left, lastSeen[ch] + 1). The max ensures we never move left backwards (the duplicate might be to the left of the current window).
// Algorithm (step-by-step):
// 1.left = 0, maxLen = 0, lastSeen = {}
//  2. For right from 0 to n-1: a. ch = s[right] b. If ch in lastSeen AND lastSeen[ch] >= left:
//     left = lastSeen[ch] + 1 (jump past the previous occurrence)
//     lastSeen[ch] = right
//     maxLen = max(maxLen, right - left + 1)
//  3. Return maxLen
func lengthOfLongestSubstring3(s string) int {
	lastSeen := make(map[byte]int)
	maxLen, left := 0, 0
	for right := 0; right < len(s); right++ {
		ch := s[right]
		if idx, ok := lastSeen[ch]; ok && idx >= left {
			left = idx + 1
		}
		lastSeen[ch] = right
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

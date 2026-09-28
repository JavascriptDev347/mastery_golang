# 3. Longest Substring Without Repeating Characters

[LeetCode link](https://leetcode.com/problems/longest-substring-without-repeating-characters/) · Difficulty: Medium

## Problem

Given a string `s`, find the length of the longest substring without
repeating characters.

**Example 1:**
```
Input:  s = "abcabcbb"
Output: 3          ("abc")
```

**Example 2:**
```
Input:  s = "bbbbb"
Output: 1          ("b")
```

**Example 3:**
```
Input:  s = "pwwkew"
Output: 3          ("wke" — "pwke" is a subsequence, not a substring)
```

**Constraints:**
- 0 <= s.length <= 5 * 10^4
- `s` consists of English letters, digits, symbols and spaces.

## Approaches

| Approach | Function | Time | Space | Notes |
|----------|----------|------|-------|-------|
| Brute Force | `lengthOfLongestSubstring` | O(n²) | O(n) | Checks every substring `s[i..j]` from scratch |
| Sliding Window | `lengthOfLongestSubstring2` | O(n) | O(n) | Shrinks the window one step at a time |
| Sliding Window + Last-Seen Index | `lengthOfLongestSubstring3` | O(n) | O(n) | Jumps `left` straight past the duplicate |

### Brute Force

For every start index `i`, extend `j` forward and track which characters
have been seen in the current `s[i..j]` window with a fresh `map[byte]bool`.
The moment a repeat shows up, stop extending — any further extension would
still contain that repeat, so there's no point continuing the inner loop.

```go
func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	n := len(s)
	for i := 0; i < n; i++ {
		seen := make(map[byte]bool)
		for j := i; j < n; j++ {
			if seen[s[j]] {
				break // duplicate found
			}
			seen[s[j]] = true
			if j-i+1 > maxLen {
				maxLen = j - i + 1
			}
		}
	}
	return maxLen
}
```

**Why it's slow:** every time `i` advances by one, the whole `seen` map is
rebuilt from scratch and the inner loop re-walks characters the outer loop
has already looked at once — that repeated work is exactly what the sliding
window below eliminates.

### Sliding Window

Same idea, but instead of restarting `seen` for every `i`, keep **one**
window `[left, right]` alive across the whole string. Expand `right`; if
`s[right]` is already in the window, shrink `left` one character at a time
— removing characters from `seen` as they leave the window — until the
duplicate is gone.

```go
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
```

Each pointer (`left`, `right`) only ever moves forward, and each moves at
most `n` times total across the whole run — that's the amortized argument
for why this is O(n) despite the nested loop shape (see
[03-sliding-window](../../03-sliding-window)).

### Sliding Window + Last-Seen Index (optimal)

The remaining inefficiency: on a string like `"aXXXXXa"`, finding the
second `a` still forces `left` to creep past every `X` one step at a time.
Fix: instead of a `seen` set, keep a `map[byte]int` of each character's
**most recent index**. When a duplicate is found, jump `left` directly to
`lastSeen[ch] + 1` — no per-character shrinking loop needed.

```go
func lengthOfLongestSubstring3(s string) int {
	lastSeen := make(map[byte]int)
	maxLen, left := 0, 0
	for right := 0; right < len(s); right++ {
		ch := s[right]
		if idx, ok := lastSeen[ch]; ok && idx >= left {
			left = idx + 1 // jump past the previous occurrence
		}
		lastSeen[ch] = right
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}
```

The `idx >= left` check matters: without it, a stale `lastSeen[ch]` entry
from *before* the current window (already pushed out by an earlier shrink)
would incorrectly yank `left` backwards.

## Tests

`solution_test.go` runs the same table of cases against all three
functions, so every approach is checked against the same inputs:

```go
{"mixed repeats", "abcabcbb", 3},
{"all same char", "bbbbb", 1},
{"substring not subsequence", "pwwkew", 3},
{"empty string", "", 0},
{"single char", "a", 1},
{"all unique", "abcdef", 6},
{"repeat at the end", "abba", 2},
{"spaces and symbols", "a !@# a", 5},
```

Run them with:

```
go test ./leetcode/01-arrays-strings/0003-longest-substring-without-repeating-characters/... -v
```

All 3 functions × 8 cases = 24 subtests, all passing.

## Key takeaway

Brute force re-derives the same "have I seen this?" fact repeatedly by
restarting from scratch at every `i` — the sliding window fixes that by
keeping *one* window's state alive and only ever moving `left`/`right`
forward. The last-seen-index variant is the next level of the same idea:
instead of shrinking the window step by step, jump `left` directly using
information already stored in the map — turning a solution that's O(n) in
theory but has repeated per-character shrinking into one with no wasted
steps at all.

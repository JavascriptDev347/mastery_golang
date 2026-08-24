package longestsubstring

import "testing"

func cases() []struct {
	name string
	s    string
	want int
} {
	return []struct {
		name string
		s    string
		want int
	}{
		{"mixed repeats", "abcabcbb", 3},
		{"all same char", "bbbbb", 1},
		{"substring not subsequence", "pwwkew", 3},
		{"empty string", "", 0},
		{"single char", "a", 1},
		{"all unique", "abcdef", 6},
		{"repeat at the end", "abba", 2},
		{"spaces and symbols", "a !@# a", 5},
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, tc := range cases() {
		t.Run(tc.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tc.s); got != tc.want {
				t.Errorf("lengthOfLongestSubstring(%q) = %d, want %d", tc.s, got, tc.want)
			}
		})
	}
}

func TestLengthOfLongestSubstring2(t *testing.T) {
	for _, tc := range cases() {
		t.Run(tc.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring2(tc.s); got != tc.want {
				t.Errorf("lengthOfLongestSubstring2(%q) = %d, want %d", tc.s, got, tc.want)
			}
		})
	}
}

func TestLengthOfLongestSubstring3(t *testing.T) {
	for _, tc := range cases() {
		t.Run(tc.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring3(tc.s); got != tc.want {
				t.Errorf("lengthOfLongestSubstring3(%q) = %d, want %d", tc.s, got, tc.want)
			}
		})
	}
}

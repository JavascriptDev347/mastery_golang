package twosum

import "testing"

func TestBruteForce(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		target int
		want   bool
	}{
		{"pair exists", []int{2, 7, 11, 15}, 9, true},
		{"pair missing", []int{3, 2, 4}, 100, false},
		{"needs same value twice but only one copy", []int{3, 5}, 6, false},
		{"empty slice", []int{}, 0, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := BruteForce(tc.nums, tc.target); got != tc.want {
				t.Errorf("BruteForce(%v, %d) = %v, want %v", tc.nums, tc.target, got, tc.want)
			}
		})
	}
}

func TestTwoPointers(t *testing.T) {
	cases := []struct {
		name       string
		sortedNums []int
		target     int
		want       bool
	}{
		{"pair exists", []int{1, 2, 3, 4, 5, 7, 8, 9}, 12, true},
		{"pair missing", []int{1, 2, 3, 4, 5}, 100, false},
		{"two elements match", []int{2, 7}, 9, true},
		{"empty slice", []int{}, 0, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := TwoPointers(tc.sortedNums, tc.target); got != tc.want {
				t.Errorf("TwoPointers(%v, %d) = %v, want %v", tc.sortedNums, tc.target, got, tc.want)
			}
		})
	}
}

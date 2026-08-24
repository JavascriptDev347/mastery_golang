// Package twosum solves LeetCode 1. Two Sum.
package twosum

// BruteForce checks every pair with a nested loop.
// Time: O(n^2), Space: O(1).
func BruteForce(nums []int, target int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return true
			}
		}
	}
	return false
}

// TwoPointers closes in from both ends of a sorted slice.
// Only correct on already-sorted input (it doesn't sort internally, since
// sorting here would discard the original indices).
// Time: O(n), Space: O(1).
func TwoPointers(sortedNums []int, target int) bool {
	left, right := 0, len(sortedNums)-1

	for left < right {
		currentSum := sortedNums[left] + sortedNums[right]

		if currentSum == target {
			return true
		}

		if currentSum < target {
			left++
		} else {
			right--
		}
	}

	return false
}

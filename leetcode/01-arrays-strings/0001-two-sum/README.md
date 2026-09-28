# 1. Two Sum

[LeetCode link](https://leetcode.com/problems/two-sum/) · Difficulty: Easy

## Problem

Given an array of integers `nums` and an integer `target`, return true if
two numbers in the array add up to `target`.

**Example:**
```
Input:  nums = [2, 7, 11, 15], target = 9
Output: true          (2 + 7 == 9)
```

**Constraints:**
- 2 <= nums.length <= 10^4
- Each input has exactly one solution; the same element can't be used twice.

> Note: the canonical LeetCode version asks for the *indices* of the two
> numbers, not just a bool. The functions here currently answer "does such a
> pair exist" — kept as originally written; returning indices (and adding a
> hash-map approach for that) is a natural next revisit.

## Approaches

| Approach | Time | Space | Notes |
|----------|------|-------|-------|
| Brute Force | O(n^2) | O(1) | Check every pair directly |
| Two Pointers | O(n) | O(1) | Requires the input already sorted |

### Brute Force

Nested loop over every `(i, j)` pair, `i < j`. Always correct, always the
baseline to compare against — the question is always "can I avoid the
inner loop."

### Two Pointers

Once the array is sorted, start pointers at both ends: if the sum is too
small, move `left` right (increase the sum); if too big, move `right` left
(decrease the sum). Each step eliminates one candidate, so the whole scan is
O(n). The catch: this only tells you a pair *exists*, not the original
indices, since sorting scrambles index order — see the note above.

### Not yet implemented: HashMap

O(n) time, O(n) space, and — unlike two-pointers — naturally returns the
*original* indices without needing the array sorted: for each `nums[i]`,
check whether `target - nums[i]` was already seen (store value → index as
you go). Worth adding here as the third approach.

## Key takeaway

Two Pointers trades the "must be sorted" requirement for O(1) extra space;
HashMap trades O(n) space to avoid needing sorted input *and* to keep
original indices. When indices/order matter, reach for the hash map, not
two pointers.

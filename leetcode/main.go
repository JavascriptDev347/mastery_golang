package main

// two sum
func twoSum(nums []int, target int) bool {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return true
			}
		}
	}
	return false
}

func twoSum2(nums []int, target int) bool {
	left, right := 0, len(nums)-1

	for left < right {
		currentSum := nums[left] + nums[right]

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

func main() {
	nums := []int{1, 2, 3, 4, 5, 7, 8, 9}
	target := 12
	result := twoSum2(nums, target)
	println(result)
}

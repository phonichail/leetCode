package twosum

// TwoSum returns the indexes of two numbers in an array which sum up to the target.
// No result returns {0, 0}
func TwoSum(nums []int, target int) []int {
	result := []int{0, 0}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result[0] = i
				result[1] = j
			}
		}
	}
	return result
}

// https://leetcode.com/problems/search-insert-position/

// Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

// You must write an algorithm with O(log n) runtime complexity.

func searchInsert(nums []int, target int) int {
    if len(nums) == 0 {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	left, right := 0, len(nums)-1
	mid := 0
	for left <= right {
		if nums[mid] == target {
			return mid
		}
		if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
        
        
	}
    
	if nums[mid] > target {
		return mid
        
	}
	return mid + 1
}

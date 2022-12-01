// https://leetcode.com/problems/search-in-rotated-sorted-array/



func search(nums []int, target int) int {
    
    left := 0
    right := len(nums) - 1
    
    for left <= right {
        mid := left+(right-left)/2
        if nums[mid] == target{
            return mid
        }
        
        if nums[left] <= nums[mid]{
            if nums[left] > target || nums[mid] < target {
                left = mid + 1
            } else {
                right = mid - 1
            }
        } else {
            if nums[left] > target && nums[mid] < target {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }
    
    return -1
    
}




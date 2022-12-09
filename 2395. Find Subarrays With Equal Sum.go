// https://leetcode.com/problems/find-subarrays-with-equal-sum/description/

// Given a 0-indexed integer array nums, determine whether there exist two subarrays of length 2 with equal sum. Note that the two subarrays must begin at different indices.

// Return true if these subarrays exist, and false otherwise.

// A subarray is a contiguous non-empty sequence of elements within an array.

func findSubarrays(nums []int) bool {
    
    hash:= make(map[int]int)
    size := len(nums) - 1
    
    for i:= 0 ; i < size ; i++ {
        sum := nums[i] + nums[i+1]
        hash[sum]++
        if hash[sum] >= 2 {
            return true
        }
    }

    return false
}


func findSubarrays(nums []int) bool {
    var count int = 0
    size := len(nums) - 1
    
    for i:= 0 ; i < size ; i++ {
        for j := i+1 ; j < size ; j++ {
            if ((nums[i] + nums[i+1]) == (nums[j] + nums[j+1])) {
                count++
            }
        }
    }
    return count > 0 
}

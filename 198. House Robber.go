// https://leetcode.com/problems/house-robber/description/

// You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

// Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.


func rob(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    
    if len(nums) == 2 {
        return max(nums[0], nums[1])
    }


     robber := make([]int, len(nums))
    robber[0], robber[1] = nums[0], max(nums[0], nums[1])
    for i := 2; i < len(nums); i++ {
        robber[i] = max(robber[i-2] + nums[i], robber[i-1])
    }
    return robber[len(nums)-1]
}



func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}


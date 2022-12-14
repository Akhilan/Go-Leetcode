// https://leetcode.com/problems/maximum-subarray/


// Given an integer array nums, find the subarray which has the largest sum and return its sum.

func maxSubArray(nums []int) int {
	curSum, maxSum := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		curSum = max(curSum+nums[i], nums[i])
		maxSum = max(curSum, maxSum)
	}
	return maxSum


	}

func max(a,b int) int {
    if a>b { 
        return a
    }
    return b
    }

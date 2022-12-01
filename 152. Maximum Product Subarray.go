// https://leetcode.com/problems/maximum-product-subarray/


// Given an integer array nums, find a subarray that has the largest product, and return the product.

// The test cases are generated so that the answer will fit in a 32-bit integer.


func maxProduct(nums []int) int {
	res, curMin, curMax := nums[0], nums[0], nums[0]
	for _, v := range nums[1:] {
        if v < 0 {
            curMin, curMax = curMax, curMin
        }
        
        curMin = min(curMin*v, v)
        curMax = max(curMax*v, v)
		res = max(res, curMax)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
    
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
    
	return b
}

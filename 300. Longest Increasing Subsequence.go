// https://leetcode.com/problems/longest-increasing-subsequence/

func lengthOfLIS(nums []int) int {
    if len(nums) == 0 {
		return 0
	}
    dp := make([]int, len(nums))
	res := math.MinInt32
    
    
    for i := 0; i < len(nums); i++ {
		currentMax := math.MinInt
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				currentMax = max(currentMax, dp[j])
			}
		}

		if currentMax == math.MinInt {
			dp[i] = 1
		} else {
			dp[i] = currentMax + 1
		}

		res = max(res, dp[i])
	}

	return res
}



func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

// https://leetcode.com/problems/coin-change/

func coinChange(coins []int, amount int) int {

    dp := make([]int, amount + 1)
    
    
    for i := 1; i <= amount; i++ {
        minCoin := math.MaxInt32
        for _, coin := range coins {
            if i-coin >= 0 && dp[i-coin] != -1 {
               minCoin = min(minCoin, dp[i-coin]+1)
            }
        }
        if minCoin == math.MaxInt32 {
            dp[i] = -1
        } else {
            dp[i] = minCoin
        }
    }
    
    return dp[amount]
        
}



func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

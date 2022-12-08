// https://leetcode.com/problems/min-cost-climbing-stairs/description/

// You are given an integer array cost where cost[i] is the cost of ith step on a staircase. Once you pay the cost, you can either climb one or two steps.

// You can either start from the step with index 0, or the step with index 1.

// Return the minimum cost to reach the top of the floor.

func minCostClimbingStairs(cost []int) int {
    
    cost = append(cost, 0)
    for i := len(cost) - 3 ; i >= 0 ; i -- {
        
        cost[i] += min(cost[i+2], cost[i+1])
        
    }
    return min(cost[0], cost[1])
}


func min(a, b int) int {
    
    if a<b {
        return a
    }
    
    return b
}

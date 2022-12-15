// https://leetcode.com/problems/jump-game-ii/description/

func jump(nums []int) int {
    max, res, curMax := 0, 0, 0
    for i := 0; i < len(nums) - 1; i++ {
        max = Max(max, nums[i] + i)
        if curMax == i {
            res++
            curMax = max
        }
    }
    return res
}

func Max(a, b int) int {
    if a > b { return a }
    return b
}

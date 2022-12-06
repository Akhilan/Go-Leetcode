// https://leetcode.com/problems/number-of-distinct-averages/description/


func distinctAverages(nums []int) int {
    
    res, sum := 0, map[int]bool{}
    sort.Ints(nums)
    for i, j := 0, len(nums) - 1; i <= j; i, j = i + 1, j - 1 {
        if a := nums[i] + nums[j]; !sum[a] {
            sum[a] = true
            res++
        }
    }
    return res
}



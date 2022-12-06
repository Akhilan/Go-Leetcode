// https://leetcode.com/problems/apply-operations-to-an-array/description/

func applyOperations(nums []int) []int {
    
    for i := 0; i < len(nums) - 1; i++ {
        if nums[i] == nums[i + 1] {
            nums[i] *= 2
            nums[i + 1] = 0
        }
    }
    ans, idx := make([]int, len(nums)), 0
    for _, num := range nums {
        if num != 0 {
            ans[idx] = num
            idx++
        }
    }
    return ans
    
}

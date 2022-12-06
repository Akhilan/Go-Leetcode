// https://leetcode.com/problems/number-of-unequal-triplets-in-array/description/

func unequalTriplets(nums []int) int {
    ans := 0
    for i := 0; i < len(nums) - 2; i++ {
        for j := i + 1; j < len(nums) - 1; j++ {
            for k := j + 1; k < len(nums); k++ {
                if nums[i] != nums[j] && nums[i] != nums[k] && nums[j] != nums[k] {
                    ans++
                }
            }
        }
    }
    return ans
}

// https://leetcode.com/problems/3sum/

// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

// Notice that the solution set must not contain duplicate triplets.



func threeSum(nums []int) [][]int {
    
    res:=[][]int{}
    sort.Ints(nums)
    
    for i := range nums {
        if i>0 && nums[i]==nums[i-1] {
            continue 
        }
        j, k := i+1, len(nums)-1
        for j<k {
            if j > i+1 && nums[j] == nums[j-1] {
                j++
                continue 
            }
            if k < len(nums) - 1 && nums[k] == nums[k+1] {
                k--
                continue
            }
            sum := nums[i] + nums[j] + nums[k]                   
            if sum==0 {
                        res = append(res, []int{nums[i], nums[j], nums[k]})
                j++
                
                k--
                
            } else if sum > 0 {
                k--
            } else {
                j++
            }
            
            
        }
    }
    return res
}

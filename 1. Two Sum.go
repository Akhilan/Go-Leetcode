// https://leetcode.com/problems/two-sum/

//Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

//You may assume that each input would have exactly one solution, and you may not use the same element twice.

//You can return the answer in any order.




func twoSum(nums []int, target int) []int {
    
    record := make(map[int]int)
    result := []int{}
    
    for k, v := range nums {
        if val, ok := record[v]; ok {
            result = append(result, k)
            result = append(result, val)
            
            return result
        }
        
        record[target - v] = k
    }
    
    return result
}






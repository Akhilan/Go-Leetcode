// https://leetcode.com/problems/largest-positive-integer-that-exists-with-its-negative/description/

// Given an integer array nums that does not contain any zeros, find the largest positive integer k such that -k also exists in the array.

// Return the positive integer k. If there is no such integer, return -1.

func findMaxK(nums []int) int {
    ans, tmp := -1, make(map[int]bool)
    for _, num := range nums {
        if num < 0 {
            tmp[num] = true
        }
    }
    for _, num := range nums {
        if num > 0 {
            if _, ok := tmp[-num]; ok {
                ans = max(ans, num)
            }
        }
    }
    return ans
}

func max(a, b int) int {
    if a > b { return a }
    return b
}


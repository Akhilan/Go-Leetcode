// https://leetcode.com/problems/most-frequent-even-element/description/

// Given an integer array nums, return the most frequent even element.

// If there is a tie, return the smallest one. If there is no such element, return -1.

func mostFrequentEven(nums []int) int {
    m:=make(map[int]int)
    res:=0
    max:=0
    for _,n:= range nums{
        if n%2==0{
             m[n]++
        }
    }
    for i, n := range m{
        if n > max || (n == max && res > i){
            max = n
            res = i
        }
    }
    if max == 0{
        return -1
    }
    return res
}


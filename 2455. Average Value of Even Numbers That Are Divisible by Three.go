// https://leetcode.com/problems/average-value-of-even-numbers-that-are-divisible-by-three/description/

func averageValue(nums []int) int {
    sum :=0
    count :=0
    for _, i := range nums {
        if i % 3 == 0 && i %2 ==0 {
            sum += i
            count++
        }
    }
    if count != 0 {
        return sum/count
    } 
    return 0
}

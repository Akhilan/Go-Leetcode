// https://leetcode.com/problems/longest-subsequence-with-limited-sum/

func answerQueries(nums []int, queries []int) []int {
    sort.Ints(nums)
    
    res := make([]int, len(queries))
    for i, query := range queries {
        for _, num := range nums {
            if query - num >= 0 {
                query -= num
                res[i]++   
            } else {
                break
            }
        }
    }
	
    return res
}

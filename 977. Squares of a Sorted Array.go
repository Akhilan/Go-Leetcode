// https://leetcode.com/problems/squares-of-a-sorted-array/description/

func sortedSquares(A []int) []int {
	ret := make([]int, len(A))
	left, right := 0, len(A)-1
	for i := len(A)-1; i >= 0; i-- {
        if math.Abs(float64(A[left])) < math.Abs(float64(A[right])) {
			ret[i] = A[right] * A[right]
            right--
		} else {
			ret[i] = A[left] * A[left]
            left++
		}
	}
	return ret
}



func sortedSquares(nums []int) []int {
    out := make([]int, 0)
    for _, v := range nums {
        out = append(out, v*v)
    }
    sort.Ints(out)
    return out
}

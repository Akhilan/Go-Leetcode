// https://leetcode.com/problems/minimum-falling-path-sum/description/

func minFallingPathSum(A [][]int) int {
    res := math.MaxInt32
    for i := 1; i < len(A); i++ {
        for j := 0; j < len(A); j++ {
            A[i][j] += min(A[i-1][j], min(A[i-1][max(0, j-1)], A[i-1][min(len(A)-1, j+1)]))
        }
    }
    for _, v := range A[len(A)-1] {
        res = min(res, v)
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

    


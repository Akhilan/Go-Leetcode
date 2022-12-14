// https://leetcode.com/problems/unique-paths/description/

func uniquePaths(m int, n int) int {
    
    ways := make([][]int, m)
    for r := range ways {
        ways[r] = make([]int, n)
        ways[r][0] = 1 
    }
    for c := range ways[0] { 
        ways[0][c] = 1 
    }
    
    for r := 1; r < m; r++ {
        for c := 1; c < n; c++ {
            ways[r][c] = ways[r-1][c] + ways[r][c-1]
        }
    }
    return ways[m-1][n-1]
    
}

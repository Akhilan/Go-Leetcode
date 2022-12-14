// https://leetcode.com/problems/rotting-oranges/

func orangesRotting(grid [][]int) int {
    var rotting [][]int
    n := len(grid)
    m := len(grid[0])


    for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if grid[r][c] == 2 {
				rotting = append(rotting, []int{r, c})
			}
		}
	}


    var queue [][]int
	times:=0
	for _, r := range rotting {
		queue = append(queue, []int{r[0], r[1], times})
	}

    for len(queue) != 0 {
		node := queue[0]
		times = node[2]
		if node[0]-1 >= 0 && grid[node[0]-1][node[1]] == 1 {
			grid[node[0]-1][node[1]] = 2
			queue = append(queue, []int{node[0] - 1, node[1], times + 1})
		}
		if node[0]+1 < n && grid[node[0]+1][node[1]] == 1 {
			grid[node[0]+1][node[1]] = 2
			queue = append(queue, []int{node[0] + 1, node[1], times + 1})
		}
		if node[1]-1 >= 0 && grid[node[0]][node[1]-1] == 1 {
			grid[node[0]][node[1]-1] = 2
			queue = append(queue, []int{node[0], node[1] - 1, times + 1})
		}
		if node[1]+1 < m && grid[node[0]][node[1]+1] == 1 {
			grid[node[0]][node[1]+1] = 2
			queue = append(queue, []int{node[0], node[1] + 1, times + 1})
		}
		queue = queue[1:]
	}
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if grid[r][c] == 1 {
				times = -1
				goto result
			}
		}
	}
result:
	return times
}


// https://leetcode.com/problems/candy/description/

func candy(ratings []int) int {
	length := len(ratings)
	res := make([]int, length)
	for i := 0; i < length; i++ {
		res[i] = 1
	}
	for i := 1; i < length; i++ {
		if ratings[i] > ratings[i-1] {
		    res[i] = res[i-1] + 1
	    }
	}
	for i := length - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
		res[i-1] = max(res[i]+1, res[i-1])
	}
	}
	count := 0
	for _, item := range res {
		count += item
	}
	return count
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}




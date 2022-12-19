// https://leetcode.com/problems/daily-temperatures/description/

func dailyTemperatures(temperatures []int) []int {
	var arr []int
	var count int
	for i, val := range temperatures {
		flag := true
		for j := i + 1; j < len(temperatures); j++ {
			count = 0
			if temperatures[j] > val {
				flag = false
				count = j - i
				break
			}
		}
		if flag {
			arr = append(arr, 0)
		} else {
			arr = append(arr, count)
		}

	}
	return arr
}



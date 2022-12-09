// https://leetcode.com/problems/sort-the-people/description/

// You are given an array of strings names, and an array heights that consists of distinct positive integers. Both arrays are of length n.

// For each index i, names[i] and heights[i] denote the name and height of the ith person.

// Return names sorted in descending order by the people's heights.

func sortPeople(names []string, heights []int) []string {
    
    
    dict:= make(map[int]string)
	ans:= []string{}
    for i,val:= range heights{
		dict[val] = names[i]
	}
	sort.Ints(heights)
	fmt.Println(heights)
	for i:= range heights{
		ans =append(ans, dict[heights[len(heights)-i-1]])
	}
	return ans
    
}

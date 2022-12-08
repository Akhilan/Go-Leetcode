// https://leetcode.com/problems/determine-if-two-strings-are-close/description/

func closeStrings(word1 string, word2 string) bool {
    
    if len(word1) != len(word2) {
		return false
	}

	counter1, counter2 := [26]int{}, [26]int{}
	for i := range word1 {
		counter1[word1[i]-'a']++
		counter2[word2[i]-'a']++
	}

    
	for i := range counter1 {
		if counter1[i] != 0 && counter2[i] == 0 {
			return false
		}
	}

    
	m := map[int]int{}
	for i := range counter1 {
		m[counter1[i]]++
		m[counter2[i]]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}



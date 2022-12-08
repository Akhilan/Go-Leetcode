// https://leetcode.com/problems/sort-characters-by-frequency/description/

// Given a string s, sort it in decreasing order based on the frequency of the characters. The frequency of a character is the number of times it appears in the string.

// Return the sorted string. If there are multiple answers, return any of them.

func frequencySort(s string) string {
    var t []byte
    freq := make(map[byte]int)
    
    for i:= 0; i < len(s); i++ {
        freq[s[i]]++
    }
    
    num_unique_letters := len(freq)
    for i := 0; i < num_unique_letters; i++ {
        var max_letter byte
        max_count := 0
        for letter, count := range freq {
            if count >= max_count {
                max_count = count
                max_letter = letter
            }
        }
        delete(freq, max_letter)
        for j:=0; j < max_count; j++ {
            t = append(t, max_letter)
        }
    }
    
    return string(t)
}


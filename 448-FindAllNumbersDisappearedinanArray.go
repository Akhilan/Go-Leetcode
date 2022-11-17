func findDisappearedNumbers(nums []int) []int {
    
    var output []int
	hash := make(map[int]bool)

    
    //If present, setting the hashvalues of corresponding numbers to true
	for _, num := range nums {
		hash[num] = true
	}

    //Appending the non present values in array to output array
	for i := 1; i <= len(nums); i++ {
		if !hash[i] {
			output = append(output, i)
		}
	}

	return output
    }

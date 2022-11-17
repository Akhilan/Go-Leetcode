func canPlaceFlowers(flowerbed []int, n int) bool {
    
    
    count, result, size := 0, 0, len(flowerbed)
    
    for i := 0; i <= size; i++ {
        
        //Counting number of empty positions
        if i < size && flowerbed[i] == 0 {
            count++
            
            //Adding edge cases in both side, since no need to check one side in both cases
            if i == 0 {
				count++
			}
			if i == size-1 {
				count++
			}
        } else {
            // checking number of flowers can be planted
            result += (count - 1) / 2
            
            //checking the given value is equal or less than result - Atleast n flowers can be planted
			if result >= n {
				return true
             }
            
        //Resetting count value to zero, to return in last count
			count = 0
        }
        }

	return false
}
    

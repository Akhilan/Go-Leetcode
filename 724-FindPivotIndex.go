func pivotIndex(nums []int) int {
    
    size :=len(nums)
    sumL := 0
    sumR := 0
    
    for i:= 0; i<size ; i++ {
        sumR += nums[i]
    }
    

    for i:= 0; i<size ; i++ {
        sumR -= nums[i]
        
        if sumL == sumR {
		return i
        }
        
        sumL += nums[i]
    }
      return -1
}

// https://leetcode.com/problems/jump-game/description/

func canJump(nums []int) bool {
	length := len(nums)
	if length == 1 {
		return true
	} else if nums[0] == 0 && length > 1 {
		return false
	}

	maxJump := nums[0]
    
	for i := 1; i < length-1; i++ {
		next := nums[i] + i
		if next > maxJump {
			maxJump = next
		}

		if nums[i] == 0 && i >= maxJump {
			return false
		}
	}

	return true
}

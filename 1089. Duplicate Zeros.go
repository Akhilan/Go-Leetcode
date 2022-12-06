// https://leetcode.com/problems/duplicate-zeros/description/


// Given a fixed-length integer array arr, duplicate each occurrence of zero, shifting the remaining elements to the right.

// Note that elements beyond the length of the original array are not written. Do the above modifications to the input array in place and do not return anything.


func duplicateZeros(arr []int)  {
    for i := 0; i < len(arr); i++ {
        if arr[i] == 0 {
            copy(arr[i + 1:], arr[i:])
            arr[i] = 0
            i++
        }
    }
}



func duplicateZeros(arr []int)  {
    for index := 0; index < len(arr); index++ {
		if arr[index] == 0 && index+1 < len(arr) {
			arr = append(arr[:index+1], arr[index:len(arr)-1]...)
			index++
		}
	}
}


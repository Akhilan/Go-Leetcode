// https://leetcode.com/problems/merge-sorted-array/


func merge(nums1 []int, m int, nums2 []int, n int)  {
    j := 0
  
    for i := m; i < m+n; i++ {
        nums1[i] = nums2[j]
        j++
    }
    
    sort.Ints(nums1)
    
}

// https://leetcode.com/problems/first-bad-version/

/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    l, r := 1, n
    for l  < r {
        mid := l + (r - l) / 2;
            if isBadVersion(mid) {
            r = mid
            } else {
            l = mid + 1
    }
    }
     return l
}


func firstBadVersion(n int) int {
    i, j := 1, n
    for i + 1 < j {
        mid := (i + j)/2
        if isBadVersion(mid) {
            j = mid
        } else {
            i = mid
        }
    }
    if isBadVersion(i) {
        return i
    }
    return j
}

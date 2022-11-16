/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    begin, last := 1, n
	for begin < last {
		mid := begin + (last-begin)/2
		switch guess(mid) {
		case 0:
			return mid
		case 1:
			begin = mid + 1
		case -1:
			last = mid
		}
	}

	return last
}
    

// https://leetcode.com/problems/determine-if-two-events-have-conflict/description/

func haveConflict(event1 []string, event2 []string) bool {
    event1start, event1end := to_second(event1)
	event2start, event2end := to_second(event2)
    
    sep := event1start > event2end || event2start > event1end
	return !sep
}


func to_second(event []string) (int, int) {
	h1, _ := strconv.Atoi(event[0][:2])
	m1, _ := strconv.Atoi(event[0][3:])
	h2, _ := strconv.Atoi(event[1][:2])
	m2, _ := strconv.Atoi(event[1][3:])
	return h1*60 + m1, h2*60 + m2
}

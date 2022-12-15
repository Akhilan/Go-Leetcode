// https://leetcode.com/problems/gas-station/description/

func canCompleteCircuit(gas []int, cost []int) int {
    index, left, g, c := 0, 0, 0, 0
    for i := range gas {
        if gas[i] + left < cost[i] {
            index = i + 1
            left = 0
        } else {
            left += gas[i] - cost[i]
        }
        g += gas[i]; c += cost[i]
    }
    if g < c { return -1 }
    return index
}



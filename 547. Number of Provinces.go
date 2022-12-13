// https://leetcode.com/problems/number-of-provinces/

func findCircleNum(isConnected [][]int) int {
    visited := make(map[int]bool)
    provinces := 0

    for index := range isConnected {
        if _, ok := visited[index]; ok {
            continue
        }
        exploreCity(isConnected, visited, index)
         provinces++
    }
    return provinces
}


func exploreCity(isConnected [][]int, visited map[int]bool, city int) {
    visited[city] = true
    for cIndex, connectedCity := range isConnected[city] {
        if connectedCity == 0 {
            continue
        }
        if _, ok := visited[cIndex]; ok {
            continue
        }
        exploreCity(isConnected, visited, cIndex)
    }
}

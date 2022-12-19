// https://leetcode.com/problems/find-if-path-exists-in-graph/description/


func validPath(n int, edges [][]int, start int, end int) bool {
    graph := buildGraph(edges, n);
    visited := make([]bool, n);
    for i := range visited { visited[i] = false; }
    
    dfs(start, graph, visited);
    
    return (visited[end]);
}

func dfs(node int, graph [][]int, visited []bool) {
    if visited[node] { return }
    visited[node] = true;
    
    for _, neigh := range graph[node] {
        dfs(neigh, graph, visited);
    }
}

func buildGraph(edges [][]int, n int) [][]int {
    graph := make([][]int, n);
    
    for _, edge := range edges {
        graph[edge[0]] = append(graph[edge[0]], edge[1]);
        graph[edge[1]] = append(graph[edge[1]], edge[0]);
    }
    
    return graph;
}

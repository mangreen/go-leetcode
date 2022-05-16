package q323_NumberOfConnectedComponentsInAnUndirectedGraph

func countComponents(n int, edges [][]int) int {
    g := make(map[int][]int, n)
    for _, e := range edges {
        g[e[0]] = append(g[e[0]], e[1])
        g[e[1]] = append(g[e[1]], e[0])
    }

    visited := make(map[int]bool)

    count := 0

    for i:=0; i<n; i++ {
        if !visited[i] {
            count++

            DFS(g, visited, i)
        }
    }

    return count
}

func DFS(g map[int][]int, v map[int]bool, cur int) {
    if v[cur] {
        return
    }

    v[cur] = true

    for _, next := range g[cur] {
        DFS(g, v, next)
    }
}
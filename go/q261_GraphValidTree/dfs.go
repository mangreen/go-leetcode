package q261_GraphValidTree

func validTree(n int, edges [][]int) bool {
    g := make(map[int][]int, n)
    for _, e := range edges {
        g[e[0]] = append(g[e[0]], e[1])
        g[e[1]] = append(g[e[1]], e[0])
    }

    visited := make(map[int]bool, n)

    if !DFS(g, visited, 0, -1) {
        return false
    }

    for _, v := range visited {
        if !v {
            return false
        }
    }

    return true
}

// visited twice, means it has a circle, not tree, return false
func DFS(g map[int][]int, v map[int]bool, cur, pre int) bool {
    if v[cur] {
        return false
    }

    v[cur] = true

    for _, next := range g[cur] {
        if next != pre && !DFS(g, v, next, cur) {
            return false
        }
    }

    return true
}
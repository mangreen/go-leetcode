package q323_NumberOfConnectedComponentsInAnUndirectedGraph

func countComponents2(n int, edges [][]int) int {
    g := make(map[int][]int, n)
    for _, e := range edges {
        g[e[0]] = append(g[e[0]], e[1])
        g[e[1]] = append(g[e[1]], e[0])
    }

    visited := make(map[int]bool, n)

    count := 0

    for i:=0; i<n; i++ {
        if !visited[i] {
            count++

            q := []int{ i }

            for len(q) > 0 {
                for _, n1 := range q {
                    q = q[1:]

                    visited[n1] = true

                    for _, n2 := range g[n1] {
                        if !visited[n2] {
                            q = append(q, n2)
                        }
                    }
                }
            }

        }
    }

    return count
}
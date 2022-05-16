package q323_NumberOfConnectedComponentsInAnUndirectedGraph

func countComponents3(n int, edges [][]int) int {
    ans := n

    root := make([]int, n)
    for i:=0; i<n; i++ {
        root[i] = -1
    }

    for _, e := range edges {
        x := find(root, e[0])
        y := find(root, e[1])

        if x != y {
            ans--
            root[y] = x
        }
    }

    return ans
}

func find(root []int, i int) int {
    for root[i] != -1 {
        i = root[i]
    }

    return i
}
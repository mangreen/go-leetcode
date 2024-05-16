package knapsack

func complete(items [][]int, W int) int {
    dp := make([]int, W+1)
    
    for _, item := range items {
        for j:=item[0]; j<=W; j++ {
            dp[j] = MAX(dp[j], dp[j-item[0]]+item[1])
        }
    }
    
    return dp[W]
}

func MAX(a, b int) int {
    if a > b {
        return a
    }
    
    return b
}
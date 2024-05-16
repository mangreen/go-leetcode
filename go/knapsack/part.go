package knapsack

func part(items [][]int, W int) int {
    dp := make([]int, W+1)
    
    for _, item := range items {
        count := make([]int, W+1)
        
        for j:=item[0]; j<=W; j++ {
            if count[j-item[0]] < item[2] {
                dp[j] = MAX(dp[j], dp[j-item[0]]+item[1])
            
                if dp[j] == dp[j-item[0]] + item[1] {
                    count[j] = count[j-item[0]] + 1
                }
            }
            
            if dp[j] == 0 || dp[j]<dp[j-1] {
                dp[j] = dp[j-1]
                count[j] = count[j-1]
            }
        }
    }
    
    return dp[W]
}
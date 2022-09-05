package q694_NumberOfDistinctIslands

import "strconv"

func numDistinctIslands(grid [][]int) int {
	m := make(map[string]bool)

	for i, row := range grid {
		for j := range row {
			g := ""
			
			DFS(grid, i, j, i, j, &g)

			if g != "" {
				m[g] = true
			}
		}
	}

	return len(m)
}

func DFS(grid [][]int, x0, y0 int, i, j int, g *string) {
	if i < 0 || len(grid)-1 < i || j < 0 || len(grid[0])-1 < j {
		return
	}

	if grid[i][j] != 1 {
		return
	}

	grid[i][j] = 0

	*g += strconv.Itoa(i-x0) + "_" + strconv.Itoa(j-y0) + "_"

	DFS(grid, x0, y0, i-1, j, g)
	DFS(grid, x0, y0, i+1, j, g)
	DFS(grid, x0, y0, i, j-1, g)
	DFS(grid, x0, y0, i, j+1, g)
}

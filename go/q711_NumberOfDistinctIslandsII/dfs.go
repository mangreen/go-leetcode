package q711_NumberOfDistinctIslandsII

import (
	"fmt"
	"sort"
	"strconv"
)

func numDistinctIslands2(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	ans := make(map[string]bool)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				g := [][]int{}
				DFS(grid, i, j, &g)
				ans[normalize(g)] = true
			}
		}
	}

	fmt.Println(ans)

	return len(ans)
}

func DFS(grid [][]int, i, j int, g *[][]int) {
	if i < 0 || len(grid)-1 < i || j < 0 || len(grid[0])-1 < j {
		return
	}

	if grid[i][j] != 1 {
		return
	}

	grid[i][j] *= -1

	*g = append(*g, []int{i, j})

	DFS(grid, i-1, j, g)
	DFS(grid, i+1, j, g)
	DFS(grid, i, j-1, g)
	DFS(grid, i, j+1, g)
}

func normalize(graph [][]int) string {
	shapes := make([][][]int, 8)
	for _, g := range graph {
		x := g[0]
		y := g[1]

		shapes[0] = append(shapes[0], []int{x, y})
		shapes[1] = append(shapes[1], []int{x, -y})
		shapes[2] = append(shapes[2], []int{-x, y})
		shapes[3] = append(shapes[3], []int{-x, -y})
		shapes[4] = append(shapes[4], []int{y, x})
		shapes[5] = append(shapes[5], []int{y, -y})
		shapes[6] = append(shapes[6], []int{-y, x})
		shapes[7] = append(shapes[7], []int{-y, -x})
	}

	for _, s := range shapes {
		sort.Slice(s, func(i, j int) bool {
			return s[i][1] < s[j][1]
		})
		sort.Slice(s, func(i, j int) bool {
			return s[i][0] < s[j][0]
		})

		for i:=len(s)-1; i>=0; i--  {
			s[i][0] -= s[0][0]
			s[i][1] -= s[0][1]
		}
	}

	fmt.Println("ssssss => ", shapes)

	sort.Slice(shapes, func(a, b int) bool {
		return shapes[a][0][0] < shapes[b][0][0]
	})

	res := ""
	for _, s0 := range shapes[0] {
		res += strconv.Itoa(s0[0]) + "," + strconv.Itoa(s0[1]) + "_"
	}

	return res
}

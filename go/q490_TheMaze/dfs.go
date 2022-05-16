package q490_TheMaze

var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func hasPath2(maze [][]int, start, goal []int) bool {
	return DFS(&maze, start, goal)
}

func DFS(maze *[][]int, cur, goal []int) bool {
	if cur[0] == goal[0] && cur[1] == goal[1] {
		return true
	}

	for _, d := range directions {
		newX := cur[0]
		newY := cur[1]

		for isValid2(*maze, newX+d[0], newY+d[1]) {
			newX += d[0]
			newY += d[1]
		}

		if (*maze)[newX][newY] != 2 {
			(*maze)[newX][newY] = 2

			if DFS(maze, []int{newX, newY}, goal) {
				return true
			}
		}
	}

	return false
}

func isValid2(maze [][]int, x, y int) bool {
	return 0 <= x && x <= len(maze)-1 && 0 <= y && y <= len(maze[0])-1 && maze[x][y] != 1
}

type Node struct {
	X, Y int
}

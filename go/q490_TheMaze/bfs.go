package q490_TheMaze

func hasPath(maze [][]int, start, goal []int) bool {
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	q := []Point{{X: start[0], Y: start[1]}}

	for len(q) > 0 {
		p := q[0]
		q = q[1:]

		if p.X == goal[0] && p.Y == goal[1] {
			return true
		}

		for _, d := range directions {
			newX := p.X
			newY := p.Y

			for isValid(maze, newX+d[0], newY+d[1]) {
				newX += d[0]
				newY += d[1]
			}

			if maze[newX][newY] != 2 {
				q = append(q, Point{X: newX, Y: newY})
				maze[newX][newY] = 2
			}
		}
	}

	return false
}

func isValid(maze [][]int, x, y int) bool {
	return 0 <= x && x <= len(maze)-1 && 0 <= y && y <= len(maze[0])-1 && maze[x][y] != 1
}

type Point struct {
	X, Y int
}

package q505_theMazeII

func shortestDistance(maze [][]int, start, goal []int) int {
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	q := []Point{{X: start[0], Y: start[1]}}

	for len(q) > 0 {
		p := q[0]
		q = q[1:]

		if p.X == goal[0] && p.Y == goal[1] {
			return p.cost
		}

		for _, d := range directions {
			newX := p.X
			newY := p.Y
			newCost := p.cost

			for isValid(maze, newX+d[0], newY+d[1]) {
				newX += d[0]
				newY += d[1]
				newCost++
			}

			if maze[newX][newY] != 2 {
				q = append(q, Point{X: newX, Y: newY, cost: newCost})
				maze[newX][newY] = 2
			}
		}
	}

	return -1
}

func isValid(maze [][]int, x, y int) bool {
	return 0 <= x && x <= len(maze)-1 && 0 <= y && y <= len(maze[0])-1 && maze[x][y] != 1
}

type Point struct {
	X, Y int
	cost int
}

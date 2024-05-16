package q683_KEmptySlots

func findLonelyPixel(picture [][]byte) int {
	rowCnt := make([]int, len(picture))
	colCnt := make([]int, len(picture[0]))
	for i, row := range picture {
		for j := range row {
			if picture[i][j] == 'B' {
				rowCnt[i]++
				colCnt[j]++
			}
		}
	}

	ans := 0
	for i, row := range picture {
		for j := range row {
			if picture[i][j] == 'B' && (rowCnt[i] == 1 && colCnt[j] == 1) {
					ans++
			}
		}
	}

	return ans
}

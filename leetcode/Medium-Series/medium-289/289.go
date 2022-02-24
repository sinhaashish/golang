package main

func gameOfLife(board [][]int) [][]int {

	m := len(board)
	n := len(board[0])
	temp := make([][]int, m)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			temp[i] = make([]int, n)
			temp[i][j] = board[i][j]
		}
	}

	dx := []int{1, 1, 0, -1, -1, -1, 0, 1}
	dy := []int{0, 1, 1, 1, 0, -1, -1, -1}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			liveCount := 0
			for k := 0; k < 8; k++ {
				if isSafe(i+dx[k], j+dy[k], m, n) && temp[i+dx[k]][j+dy[k]] == 1 {
					liveCount++
				}
			}
			if temp[i][j] == 0 && liveCount == 3 {
				board[i][j] = 1
			}
			if temp[i][j] == 1 && (liveCount < 2 || liveCount > 3) {
				board[i][j] = 0
			}
		}
	}
	return board

}

func isSafe(x, y, M, N int) bool {
	return x >= 0 && x < M && y >= 0 && y < N
}

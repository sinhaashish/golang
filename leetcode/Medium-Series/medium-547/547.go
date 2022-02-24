package main

func findCircleNum(isConnected [][]int) int {
	count := 0
	for i := 0; i < len(isConnected); i++ {
		for j := 0; j < len(isConnected[0]); j++ {
			if isConnected[i][j] == 1 {
				count++
				populate(isConnected, i, j)
			}
		}
	}
	return count
}

func populate(isConnected [][]int, i, j int) {
	if i < 0 || i > len(isConnected)-1 || j < 0 || j > len(isConnected[0])-1 || isConnected[i][j] == 0 {
		return
	}
	isConnected[i][j] = 0
	populate(isConnected, i+1, j)
	populate(isConnected, i, j+1)
	populate(isConnected, i-1, j)
	populate(isConnected, i, j+1)

}

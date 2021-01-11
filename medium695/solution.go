package medium695

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	var dfs func(i, j int)
	var count int
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return
		}
		count++
		grid[i][j] = 0
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}

	var max int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
			}
			if count > max {
				max = count
			}
			count = 0
		}
	}
	return max
}

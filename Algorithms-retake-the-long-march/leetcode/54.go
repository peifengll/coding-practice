package leetcode

func spiralOrder(matrix [][]int) (res []int) {
	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var dfs func(x, y, d int)
	dfs = func(x, y, d int) {
		res = append(res, matrix[x][y])
		matrix[x][y] = -1000
		if x+dir[d][0] < len(matrix) && x+dir[d][0] >= 0 && y+dir[d][1] < len(matrix[0]) && y+dir[d][1] >= 0 && matrix[x+dir[d][0]][y+dir[d][1]] != -1000 {
			dfs(x+dir[d][0], y+dir[d][1], d)
		} else {
			d = (d + 1) % 4
			if x+dir[d][0] < len(matrix) && x+dir[d][0] >= 0 && y+dir[d][1] < len(matrix[0]) && y+dir[d][1] >= 0 && matrix[x+dir[d][0]][y+dir[d][1]] != -1000 {
				dfs(x+dir[d][0], y+dir[d][1], d)
			}
		}
	}
	dfs(0, 0, 0)
	return
}

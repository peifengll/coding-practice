package leetcode

import (
	"fmt"
	"testing"
)

func gameOfLife(board [][]int) {
	var res [30][30]int
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			res[i][j] = board[i][j]
		}
	}
	judge := func(x, y, status int) int {
		//	 返回存活还是死亡
		dir := [8][2]int{
			{0, 1},
			{1, 0},
			{0, -1},
			{-1, 0},
			{1, 1},
			{1, -1},
			{-1, 1},
			{-1, -1},
		}
		count := 0
		for _, v := range dir {
			if x+v[0] < 0 || x+v[0] >= len(board) || y+v[1] < 0 || y+v[1] >= len(board[0]) {
				continue
			}
			if res[x+v[0]][y+v[1]] == 1 {
				count++
			}
		}
		if count < 2 {
			return 0
		} else if count == 3 {
			return 1
		} else if count == 2 {
			return status
		} else if count > 3 {
			return 0
		}
		return status
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			board[i][j] = judge(i, j, res[i][j])
		}
	}
}

func TestGameOfLife(t *testing.T) {
	board := [][]int{
		{1, 1},
		{1, 0},
	}
	gameOfLife(board)
	fmt.Println(board)
}

package leetcode

import "fmt"

//https://leetcode.cn/problems/container-with-most-water/

/*
多半在最大的那几根上
左边的右边
右边的往左
单调的进行碰撞
*/

type pair struct {
	x, y int
}

func maxArea(height []int) int {
	left := make([]int, 1)
	left[0] = 0

	for i := 1; i < len(height); i++ {
		//单增就加进去
		if height[i] > height[left[len(left)-1]] {
			left = append(left, i)
		}
	}
	fmt.Println("letf: ", left)
	right := make([]int, 1)
	right[0] = len(height) - 1
	for i := len(height) - 2; i >= 0; i-- {
		if height[i] > height[right[len(right)-1]] {
			right = append(right, i)
		}
	}
	res := 0
	for i := 0; i < len(left); i++ {
		for j := 0; j < len(right); j++ {
			x := left[i]
			y := right[j]
			if x > y {
				continue
			}
			res = maxx(res, minx(height[x], height[y])*(y-x))
		}
	}
	println(left)
	fmt.Println(right)
	return res
}

func minx(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func maxx(a, b int) int {
	if a > b {
		return a
	}
	return b
}

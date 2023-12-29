package coding_practice

import (
	"fmt"
	"math"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Printf("%.0f", math.MaxFloat32)
}

func TestSpace(t *testing.T) {
	c := ' '
	k := "blue"
	r := k + string(c) + k
	fmt.Println(r)
}

func reverseWords(s string) string {
	temp := ""
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if temp != "" {
				if res == "" {
					res += temp
				} else {
					res += " " + temp
				}
				temp = ""
			}
			continue
		}
		temp = string(s[i]) + temp
	}
	if temp != "" {
		if res != "" {
			res += " "
		}
		res += temp
	}

	return res
}

func maxSlidingWindow(nums []int, k int) []int {
	// 申明一个切片，当作队列使用
	// 要维护的是一个单调递减的队列， 注意：q中放的值是nums中的下标，
	//而不是具体的值，这样才能保最大值是否是在k位之中
	q := []int{}
	// 定义push方法
	push := func(i int) {
		// 维护的是单调递减序列，队首是最大值， 如过 要加入的大于屁股上的，就弹出屁股那的
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		// 这是往屁股后边加入
		q = append(q, i)
	}
	//先加 k个进去，懂得都懂
	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	//这个是答案
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		// 入队
		push(i)
		for q[0] <= i-k {
			// 判断队首是不是 不合法（不和法值得是是不是在这k位之中）
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}

	return ans
}

func TestSlice(t *testing.T) {
	q := []int{1, 2, 3, 4, 5}
	q = q[:len(q)-1]
	fmt.Println(q)

}

package leetcode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var findTree func(now *TreeNode) int
	res := &TreeNode{}
	res = nil
	findTree = func(now *TreeNode) int {
		n := 0
		if now == nil {
			return 0
		} else if now == p || now == q {
			n += 1
		}
		a := findTree(now.Left)
		if (a == 2 || a+n == 2) && res == nil {
			res = now
		}
		b := findTree(now.Right)
		if (b == 2 || b+n == 2) && res == nil {
			res = now
		}

		if a+b == 2 && res == nil {
			res = now
		}
		return a + b + n
	}
	findTree(root)
	if res == nil {
		res = root
	}
	return res
}

func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func maxProfit(k int, prices []int) int {
	f := make([][][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		f[i] = make([][]int, k+1)
		for j := 0; j <= k; j++ {
			f[i][j] = make([]int, 2)
		}
	}
	f[0][0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {

		for j := 1; j <= k; j++ {
			//卖出才算是完成一笔交易
			f[i][j][1] = max(f[i][j][1], f[i-1][j][0]-prices[i])
			//				"自己 "    		"跟上一天的记录保持"	"把上一天交际记录为j-1的，完成交易"
			f[i][j][0] = max(f[i][j][0], max(f[i-1][j][0], f[i-1][j-1][1]+prices[i])) // f[i-1][j][0],f[i-1][j-1][1]
		}
	}
	res := 0
	for i := 0; i <= k; i++ {
		res = max(res, f[len(prices)-1][i][0])
	}
	return res
	//	0是没有，1是持有
	//	 f[i][j][1]=f[i-1][j-1][]
}

// https://leetcode.cn/problems/add-binary/
func addBinary(a string, b string) string {
	getRe := func(c string) string {
		//返回一个倒的
		ru := []rune(c)
		for i, j := 0, len(c)-1; i < j; i, j = i+1, j-1 {
			ru[i], ru[j] = ru[j], ru[i]
		}
		return string(ru)
	}
	a = getRe(a)
	b = getRe(b)
	//从头开始加
	s := ""
	bl := false //要不要进位
	if len(a) < len(b) {
		a, b = b, a
	}
	for i := 0; i < len(a) || bl; i++ {
		//这是两个都不正常情况
		if i >= len(a) {
			if bl {
				s += "1"
				break
			}
		} else if i >= len(b) {
			// 这是有一个没完的情况
			if a[i] == '0' {
				if bl {
					s += "1"
					bl = false
				} else {
					s += "0"
				}
			} else if a[i] == '1' {
				if bl {
					s += "0"
					bl = true
				} else {
					s += "1"
				}
			}

		} else if i < len(b) {
			// 这是正常情况
			if a[i] == '0' && b[i] == '0' {
				if bl {
					s += "1"
					bl = false
				} else {
					s += "0"
				}
			} else if a[i] == '1' && b[i] == '0' {
				if bl {
					s += "0"
					bl = true
				} else {
					s += "1"
				}
			} else if a[i] == '0' && b[i] == '1' {
				if bl {
					s += "0"
					bl = true
				} else {
					s += "1"
				}
			} else if a[i] == '1' && b[i] == '1' {
				if bl {
					s += "1"
					bl = true
				} else {
					s += "0"
				}
			}
		}
	}
	return getRe(s)
}

// https://leetcode.cn/problems/set-matrix-zeroes/
func setZeroes(matrix [][]int) {
	// 有几行
	x := make([]bool, len(matrix))
	// 有几列
	y := make([]bool, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				x[i] = true
				y[j] = true
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		// 这一行都归0
		if x[i] {
			for j := 0; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	for i := 0; i < len(matrix[0]); i++ {
		if y[i] {
			for j := 0; j < len(matrix); j++ {
				matrix[j][i] = 0
			}
		}
	}

}

// https://leetcode.cn/problems/rotate-array/
func rotate(nums []int, k int) {
	k = k % len(nums)
	C := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	for i := 0; i < len(nums); i++ {
		nums[i] = C[i]
	}
}

// https://leetcode.cn/problems/valid-parentheses/
func isValid(s string) bool {
	k := []rune{'(', ')', '{', '}', '[', ']'}
	// 奇数-1，偶数+1
	mp := make(map[rune]rune)
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			mp[k[i]] = k[i+1]
		} else {
			mp[k[i]] = k[i-1]
		}
	}
	qu := make([]rune, len(s))
	l := -1
	// 如果上一个不是，那直接false
	for _, v := range s {
		if v == k[1] || v == k[3] || v == k[5] {
			// 收
			if l < 0 || qu[l] != mp[v] {
				return false
			}
			l--
		} else {
			l++
			qu[l] = v
		}
	}
	if l >= 0 {
		return false
	}
	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/minimum-absolute-difference-in-bst/
func getMinimumDifference(root *TreeNode) int {
	//	二叉搜索树，中序，左根右
	var res = int(1e5)
	last := res
	var dfs func(now *TreeNode)
	dfs = func(now *TreeNode) {
		if now == nil {
			return
		}
		dfs(now.Left)
		res = min(res, abs(last, now.Val))
		last = now.Val
		dfs(now.Right)
	}
	dfs(root)
	return res
}

// 2652. 倍数求和 https://leetcode.cn/problems/sum-multiples/description/
func sumOfMultiples(n int) int {
	sum := 0
	st := make([]bool, 10002)
	for i := 1; i <= n; i++ {
		st[3*i] = true
		st[5*i] = true
		st[7*i] = true
		if 3*i > n {
			break
		}
	}
	for i := 3; i <= n; i++ {
		if st[i] {
			sum += i
		}
	}
	return sum
}

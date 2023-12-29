package leetcode

func removeDuplicates(nums []int) int {
	j := 0
	// 把相同的找完了再放
	for l := 0; l < len(nums); l++ {
		r := l + 1
		for r < len(nums) {
			if nums[l] == nums[r] {
				r++
			} else {
				break
			}
		}
		k := nums[l]
		if r-l >= 2 {
			nums[j] = k
			j++
		}
		nums[j] = k
		j++
		l = r - 1
	}
	return j
}

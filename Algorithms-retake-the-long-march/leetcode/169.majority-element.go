package leetcode

import "sort"

func majorityElement(nums []int) int {
	sort.Ints(nums)
	// 在一半里绝对能找出来
	lenc := len(nums)/2 + 1
	for i := 0; i < lenc; i++ {
		//一样的都挨着
		r := i
		for r < len(nums) {
			if nums[r] != nums[i] {
				break
			}
			r++
		}
		if r-i >= lenc {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}

package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	fmt.Println(search(nums, target))
}

// search 有序nums -> 找target的下标 -> 优先想到二分查找
// 时间复杂度：O(logn)，其中 n 为 nums 的长度。
// 空间复杂度：O(1)，仅用到若干额外变量。
func search(nums []int, target int) int {
	if len(nums) == 1 {
		if target == nums[0] {
			return 0
		}
		return -1
	}

	i, j := 0, len(nums)-1
	middle := 0
	for i <= j {
		middle = i + (j-i)/2 // (i+j)/2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			j = middle - 1
			continue
		} else {
			i = middle + 1
			continue
		}
	}
	return -1
}

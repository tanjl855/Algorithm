package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}

func removeElement(nums []int, val int) int {
	//k := 0
	//for i := 0; i < len(nums); i++ {
	//	if nums[i] != val {
	//		nums[k] = nums[i]
	//		k++
	//	}
	//}
	//return k
	i, j := 0, len(nums)
	for i < j {
		if nums[i] == val {
			nums[i] = nums[j-1]
			j--
			continue
		}
		i++
	}
	return i
}

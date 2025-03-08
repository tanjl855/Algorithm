package main

import "fmt"

func main() {
	fmt.Println(maxVowels("abcdeeiiiiesaasdfououuo", 10))
}

// maxVowels https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/ (【套路】教你解决定长滑窗！适用于所有定长滑窗题目！)
// 2024-12-22
func maxVowels(s string, k int) (ans int) {
	/*
		入：下标为 i 的元素进入窗口，更新相关统计量。如果 i<k−1 则重复第一步。
		更新：更新答案。一般是更新最大值/最小值。
		出：下标为 i−k+1 的元素离开窗口，更新相关统计量。
	*/
	vowel := 0
	for i, in := range s {
		if ans == k {
			return
		}
		//进入窗口计算
		if in == 'a' || in == 'e' || in == 'i' || in == 'o' || in == 'u' {
			vowel++
		}
		if i < k-1 { //校验窗口是否满
			continue
		}
		//更新
		ans = max(ans, vowel)
		out := s[i-k+1]
		//离开窗口不计算入内
		if out == 'a' || out == 'e' || out == 'i' || out == 'o' || out == 'u' {
			vowel--
		}
	}
	return
}

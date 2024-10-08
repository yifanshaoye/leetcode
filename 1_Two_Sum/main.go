package main

import "fmt"

func twoSum1(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == i {
				continue
			}
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	hash := make(map[int]int)
	for i, num := range nums {
		hash[num] = i
	}

	for i, num := range nums {
		if j, ok := hash[target-num]; ok && i != j {
			return []int{i, j}
		}
	}
	return []int{}

}

func twoSum3(nums []int, target int) []int {
	hash := make(map[int]int)

	for i, num := range nums {
		if j, ok := hash[target-num]; ok {
			return []int{i, j}
		} else {
			hash[num] = i
		}
	}
	return []int{}

}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum1(nums, target))
	fmt.Println(twoSum2(nums, target))
	fmt.Println(twoSum3(nums, target))
}

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	target := 6

	res := twoSum(nums, target)
	fmt.Println(res)

}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		residue := target - nums[i]
		if _, ok := m[residue]; ok {
			return []int{m[residue], i}
		}
		m[nums[i]] = i
	}

	return nil
}

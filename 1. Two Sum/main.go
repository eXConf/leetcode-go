package main

import "fmt"

func main() {
	fmt.Printf("%v\n", twoSum([]int{2, 7, 11, 15}, 9)) // [0, 1]
	fmt.Printf("%v\n", twoSum([]int{3, 2, 4}, 6))      // [1, 2]
	fmt.Printf("%v\n", twoSum([]int{3, 3}, 6))         // [0, 1]
}

func twoSum(nums []int, target int) []int {
	diffsList := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		current := nums[i]

		val, exists := diffsList[current]
		if exists {
			return []int{val, i}
		}

		diffsList[diff] = i
	}

	return []int{0, 0}
}

//

// func twoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		for j := 0; j < len(nums); j++ {
// 			if j == i {
// 				continue
// 			}
// 			if nums[j]+nums[i] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}

// 	return []int{0, 0}
// }

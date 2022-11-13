package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// map to store value and index of nums
	// e.g [2, 7] as m[2] = 0 and m[7] as 1
	m := make(map[int]int, len(nums))

	// iterates nums while checking the target and storing in map
	// e.g [2, 7, 11, 15] target = 9
	for i, num := range nums {
		// when meet 2nd index (num=7), the map already has:
		// m[2] = 0, because the target is 9,
		// then target-num => 9-7 = 2 key is on exist on map
		// immediately return the index from map value and
		// current iteration index
		if val, ok := m[target-num]; ok {
			return []int{val, i}
		}

		// always store if the index still not found yet
		m[num] = i
	}

	// return nil if no index matches with the target
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

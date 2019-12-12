package program

import "sort"

func TwoNumberSumBrute(array []int, target int) []int {
	// Write your code here.

	for i := 0; i < len(array)-1; i++ {
		for j := 1; j < len(array); j++ {
			if array[i]+array[j] == target && i != j {
				if array[i] > array[j] {
					return []int{array[j], array[i]}
				}
				return []int{array[j], array[i]}
			}
		}
	}

	return []int{}
}

func TwoNumberSumHash(array []int, target int) []int {
	// Write your code here.

	m := make(map[int]int)

	for i := 0; i < len(array); i++ {
		v, isPresent := m[target-array[i]]
		if isPresent {
			return []int{v, array[i]}
		}

		m[array[i]] = array[i]
	}

	return []int{}
}

func TwoNumberSum(array []int, target int) []int {
	// Write your code here.

	sort.Ints(array)

	l := 0
	r := len(array) - 1

	for l < r {
		currentSum := array[l] + array[r]

		if currentSum == target {
			return []int{array[l], array[r]}
		} else if currentSum < target {
			l++
		} else if currentSum > target {
			r--
		}

	}

	return []int{}
}

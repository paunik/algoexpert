package twonumbers

import "sort"

// TwoNumberSumBrute O(n^2) time | O(1) space
func TwoNumberSumBrute(array []int, target int) []int {
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

// TwoNumberSumHash  with hashing O(n) time  | O(n) space
func TwoNumberSumHash(array []int, target int) []int {
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

// TwoNumberSum with presort O(n log(n)) time | O(1) space
func TwoNumberSum(array []int, target int) []int {
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

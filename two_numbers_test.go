// Add, edit, or remove tests in this file.
// Treat it as your playground!

package program

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	expected := []int{4, 6}
	output := TwoNumberSum([]int{4, 6}, 10)
	require.ElementsMatch(t, expected, output)
}

func TestCase2(t *testing.T) {
	expected := []int{1, 4}
	output := TwoNumberSum([]int{4, 6, 1}, 5)
	require.ElementsMatch(t, expected, output)
}

func TestCase3(t *testing.T) {
	expected := []int{-3, 6}
	output := TwoNumberSum([]int{4, 6, 1, -3}, 3)
	require.ElementsMatch(t, expected, output)
}

func TestCase4(t *testing.T) {
	expected := []int{-1, 11}
	output := TwoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10)
	require.ElementsMatch(t, expected, output)
}

func TestCase5(t *testing.T) {
	expected := []int{8, 9}
	output := TwoNumberSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 17)
	require.ElementsMatch(t, expected, output)
}

func TestCase6(t *testing.T) {
	expected := []int{3, 15}
	output := TwoNumberSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15}, 18)
	require.ElementsMatch(t, expected, output)
}

func TestCase7(t *testing.T) {
	expected := []int{-5, 0}
	output := TwoNumberSum([]int{-7, -5, -3, -1, 0, 1, 3, 5, 7}, -5)
	require.ElementsMatch(t, expected, output)
}

func TestCase8(t *testing.T) {
	expected := []int{-47, 210}
	output := TwoNumberSum([]int{-21, 301, 12, 4, 65, 56, 210, 356, 9, -47}, 163)
	require.ElementsMatch(t, expected, output)
}

func TestCase9(t *testing.T) {
	expected := []int{}
	output := TwoNumberSum([]int{-21, 301, 12, 4, 65, 56, 210, 356, 9, -47}, 164)
	require.ElementsMatch(t, expected, output)
}

func TestCase10(t *testing.T) {
	expected := []int{}
	output := TwoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 15)
	require.ElementsMatch(t, expected, output)
}

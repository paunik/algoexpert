package twonumbers

var testCases = []struct {
	name   string
	input  []int
	target int
	want   []int
}{
	{
		"TestCase1",
		[]int{4, 6},
		10,
		[]int{4, 6},
	},
	{
		"TestCase2",
		[]int{4, 6, 1},
		5,
		[]int{1, 4},
	},

	{
		"TestCase3",
		[]int{4, 6, 1, -3},
		3,
		[]int{-3, 6},
	},
	{
		"TestCase4",
		[]int{3, 5, -4, 8, 11, 1, -1, 6},
		10,
		[]int{-1, 11},
	},
	{
		"TestCase5",
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		17,
		[]int{8, 9},
	},
	{
		"TestCase6",
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15},
		18,
		[]int{3, 15},
	},
	{
		"TestCase7",
		[]int{-7, -5, -3, -1, 0, 1, 3, 5, 7},
		-5,
		[]int{-5, 0},
	},
	{
		"TestCase8",
		[]int{-21, 301, 12, 4, 65, 56, 210, 356, 9, -47},
		163,
		[]int{-47, 210},
	},
	{
		"TestCase9",
		[]int{-21, 301, 12, 4, 65, 56, 210, 356, 9, -47},
		164,
		[]int{},
	},
	{
		"TestCase10",
		[]int{3, 5, -4, 8, 11, 1, -1, 6},
		15,
		[]int{},
	},
}

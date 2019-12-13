// Add, edit, or remove tests in this file.
// Treat it as your playground!

package twonumbers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoNumbers(t *testing.T) {
	for _, tc := range testCases {
		got := TwoNumberSum(tc.input, tc.target)
		require.ElementsMatch(t, tc.want, got, tc.name)
	}
}

func BenchmarkHamming(b *testing.B) {
	// bench combined time to run through all test cases
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			// ignoring errors and results because we're just timing function execution
			_ = TwoNumberSum(tc.input, tc.target)
		}
	}
}

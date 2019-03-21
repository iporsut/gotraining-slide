package calc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		a, b   int
		expect int
	}{
		{10, 0, 10},
		{0, 10, 10},
		{10, 20, 30},
		{10, -20, -10},
		{-10, 20, 10},
		{-10, -20, -3},
	}

	for _, tc := range testCases {
		tcName := fmt.Sprintf("Add(%d,%d)", tc.a, tc.b)
		t.Run(tcName, func(t *testing.T) {
			assert.Equal(t, tc.expect, Add(tc.a, tc.b))
		})
	}
}

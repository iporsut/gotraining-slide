package calc

import (
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
		{-10, -20, -30},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expect, Add(tc.a, tc.b))
	}
}

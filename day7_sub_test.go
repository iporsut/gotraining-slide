package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalc(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		assert.Equal(t, 30, Add(10, 20))
	})
	t.Run("Sub", func(t *testing.T) {
		assert.Equal(t, 30, Sub(50, 20))
	})
	t.Run("Mul", func(t *testing.T) {
		assert.Equal(t, 30, Mul(15, 2))
	})
	t.Run("Div", func(t *testing.T) {
		assert.Equal(t, 30, Div(60, 2))
	})
}

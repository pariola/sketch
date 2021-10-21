package canvas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

	w, h := 10, 3 // 10x3

	c := New(w, h)

	for _, row := range c.matrix {
		assert.Equal(t, w, len(row), "invalid canvas width")
	}

	assert.Equal(t, h, len(c.matrix), "invalid canvas height")

}

func TestBoundary(t *testing.T) {

	w, h := 10, 3 // 10x3

	c := New(w, h)

	boundX, boundY := c.boundary()

	assert.Equal(t, h, boundY, "invalid x-axis boundary")
	assert.Equal(t, w, boundX, "invalid y-axis boundary")
}

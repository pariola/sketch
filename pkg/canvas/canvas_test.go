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

func TestCanvas_boundary(t *testing.T) {

	w, h := 10, 3 // 10x3

	c := New(w, h)

	boundX, boundY := c.boundary()

	assert.Equal(t, h, boundY, "invalid x-axis boundary")
	assert.Equal(t, w, boundX, "invalid y-axis boundary")
}

func TestCanvas_DrawWithFill(t *testing.T) {

	c := New(12, 12) // 12x12

	// rectangle at [2,2] width: 5, height: 5, fill: 'X', outline: none
	r := NewRectangle(2, 2, 5, 5, "X", "")

	c.Draw(*r)

	// confirm all matrix cells
	for y := r.posY; y < r.posY+r.height; y++ {
		for x := r.posX; x < r.posX+r.width; x++ {
			assert.Equalf(t, r.fillChar, c.matrix[y][x], "invalid cell [%d, %d]", x, y)
		}
	}
}

func TestCanvas_DrawWithOutline(t *testing.T) {

	c := New(12, 12) // 12x12

	// rectangle at [2,2] width: 5, height: 5, fill: none, outline: 'X'
	r := NewRectangle(2, 2, 5, 5, "", "X")

	c.Draw(*r)

	// confirm only the rectangle boundaries
	for y := r.posY; y < r.posY+r.height; y++ {
		for x := r.posX; x < r.posX+r.width; x++ {
			if inYBoundary(x, y, *r) || inXBoundary(x, y, *r) {
				assert.Equalf(t, r.outlineChar, c.matrix[y][x], "invalid cell [%d, %d]", x, y)
			}
		}
	}
}

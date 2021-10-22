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
	r := NewRectangle(2, 2, 5, 5, 'X', empty)

	c.Draw(*r)

	// confirm all rectangle cells
	for y := r.posY; y < r.posY+r.height; y++ {
		for x := r.posX; x < r.posX+r.width; x++ {
			assert.Equalf(t, r.fillChar, c.matrix[y][x], "invalid cell [%d, %d]", x, y)
		}
	}
}

func TestCanvas_DrawWithOutline(t *testing.T) {

	c := New(12, 12) // 12x12

	// rectangle at [2,2] width: 5, height: 5, fill: none, outline: 'X'
	r := NewRectangle(2, 2, 5, 5, empty, 'X')

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

func TestCanvas_DrawWithFillOutline(t *testing.T) {

	c := New(12, 12) // 12x12

	// rectangle at [2,2] width: 5, height: 5, fill: none, outline: 'X'
	r := NewRectangle(2, 2, 5, 5, 'O', 'X')

	c.Draw(*r)

	for y := r.posY; y < r.posY+r.height; y++ {
		for x := r.posX; x < r.posX+r.width; x++ {
			if inYBoundary(x, y, *r) || inXBoundary(x, y, *r) {
				// boundary
				assert.Equalf(t, r.outlineChar, c.matrix[y][x], "invalid cell [%d, %d]", x, y)
			} else {
				// insides
				assert.Equalf(t, r.fillChar, c.matrix[y][x], "invalid cell [%d, %d]", x, y)
			}
		}
	}
}

func TestCanvas_expand(t *testing.T) {

	c := New(7, 5) // 7x5

	// draw on canvas
	r := NewRectangle(0, 0, 5, 3, 'X', empty)

	c.Draw(*r)

	// expand to 10x10
	c.expand(10, 10)

	boundX, boundY := c.boundary()

	assert.Equal(t, 10, boundX, "width should increase")
	assert.Equal(t, 10, boundY, "height should increase")

	// confirm all rectangle cells
	for y := r.posY; y < r.posY+r.height; y++ {
		for x := r.posX; x < r.posX+r.width; x++ {
			assert.Equalf(t, r.fillChar, c.matrix[y][x], "invalid cell [%d, %d]", x, y)
		}
	}
}

func TestCanvas_FloodFillInside(t *testing.T) {

	c := New(7, 5) // 7x5

	// draw on canvas
	r := NewRectangle(0, 0, 5, 3, 'X', '@')

	c.Draw(*r)

	expected := [][]byte{
		{'@', '@', '@', '@', '@', empty, empty},
		{'@', 'O', 'O', 'O', '@', empty, empty},
		{'@', '@', '@', '@', '@', empty, empty},
		{empty, empty, empty, empty, empty, empty, empty},
		{empty, empty, empty, empty, empty, empty, empty},
	}

	// perform flood fill operation on boundary
	c.FloodFill(1, 1, 'O')

	// confirm if resulting matrix matches the expected matrix
	assert.Equal(t, expected, c.matrix, "result does not match expected matrix")
}

func TestCanvas_FloodFillCanvas(t *testing.T) {

	c := New(7, 5) // 7x5

	// draw on canvas
	r := NewRectangle(0, 1, 5, 3, 'X', '@')

	c.Draw(*r)

	expected := [][]byte{
		{'-', '-', '-', '-', '-', '-', '-'},
		{'@', '@', '@', '@', '@', '-', '-'},
		{'@', 'X', 'X', 'X', '@', '-', '-'},
		{'@', '@', '@', '@', '@', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-'},
	}

	// perform flood fill operation on boundary
	c.FloodFill(0, 0, '-')

	// confirm if resulting matrix matches the expected matrix
	assert.Equal(t, expected, c.matrix, "result does not match expected matrix")
}

func TestCanvas_FloodFillBoundary(t *testing.T) {

	c := New(7, 5) // 7x5

	// draw on canvas
	r := NewRectangle(0, 0, 5, 3, 'X', '@')

	c.Draw(*r)

	expected := [][]byte{
		{'-', '-', '-', '-', '-', empty, empty},
		{'-', 'X', 'X', 'X', '-', empty, empty},
		{'-', '-', '-', '-', '-', empty, empty},
		{empty, empty, empty, empty, empty, empty, empty},
		{empty, empty, empty, empty, empty, empty, empty},
	}

	// perform flood fill operation on boundary
	c.FloodFill(3, 2, '-')

	// confirm if resulting matrix matches the expected matrix
	assert.Equal(t, expected, c.matrix, "result does not match expected matrix")
}

func TestCanvas_Print(t *testing.T) {

	expected := "---------------.......------ ---------------.......------ ---------------.......------ OOOOOOOO-------.......------ O O-------.......------ O XXXXX-----.......------ OOOOOXXXXX------------------ -----XXXXX------------------ ---------------------------- ---------------------------- ---------------------------- ----------------------------"

	c := New(28, 12)

	c.Draw(*NewRectangle(15, 0, 7, 6, '.', empty))
	c.Draw(*NewRectangle(0, 3, 8, 4, empty, 'O'))
	c.Draw(*NewRectangle(5, 5, 5, 3, 'X', 'X'))

	c.FloodFill(0, 0, '-')

	assert.Equal(t, expected, c.Print())
}

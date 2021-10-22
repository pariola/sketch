package canvas

import (
	"bytes"
	"sync"
)

var (
	empty = byte(0)
)

// Canvas represents a real-life drawable canvas
type Canvas struct {
	m sync.RWMutex

	// 2D array to represent the points on a Canvas starting from (0,0) top left
	matrix [][]byte
}

// New returns an instance of the Canvas with the required size
func New(width, height int) *Canvas {

	matrix := make([][]byte, height)

	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]byte, width)
	}

	return &Canvas{
		matrix: matrix,
	}
}

// write puts value v into the underlying matrix at position [y][x]
func (c *Canvas) write(x, y int, v byte) {

	c.m.Lock()
	defer c.m.Unlock()

	c.matrix[y][x] = v
}

// read returns the value at position [y][x] from the underlying matrix
func (c *Canvas) read(x, y int) byte {

	c.m.RLock()
	defer c.m.RUnlock()

	return c.matrix[y][x]
}

// boundary returns the edges/boundaries (x-axis, y-axis) of the Canvas
func (c *Canvas) boundary() (int, int) {

	var x, y int

	y = len(c.matrix)

	if y > 0 {
		x = len(c.matrix[0])
	}

	return x, y
}

// expand increases the size of the Canvas to specified size if current size is smaller
func (c *Canvas) expand(x, y int) {

	boundX, boundY := c.boundary()

	// expand y-axis
	if y > boundY {

		// add blank rows
		for i := 0; i < (y - boundY); i++ {
			c.matrix = append(c.matrix, make([]byte, boundX))
		}

		boundY = y
	}

	// expand x-axis
	if x > boundX {
		boundX = x

		// expand all rows by copying
		for i := 0; i < boundY; i++ {
			row := make([]byte, boundX)
			copy(row, c.matrix[i])
			c.matrix[i] = row
		}
	}
}

// Draw draws a rectangle unto the Canvas
func (c *Canvas) Draw(r Rectangle) {

	// expand Canvas to contain rectangle
	c.expand(r.posX+r.width, r.posY+r.height)

	for y := r.posY; y < r.posY+r.height; y++ {
		for x := r.posX; x < r.posX+r.width; x++ {

			// fill
			c.write(x, y, r.fillChar)

			// paint outline
			// confirm if point is part of the rectangle boundary
			if r.outlineChar != byte(0) && (inYBoundary(x, y, r) || inXBoundary(x, y, r)) {
				c.write(x, y, r.outlineChar)
			}
		}
	}
}

// FloodFill wrapper for floodFill
func (c *Canvas) FloodFill(x, y int, fillChar byte) {

	// check canvas bounds
	boundX, boundY := c.boundary()

	if !inBounds(x, y, boundX, boundY) {
		return
	}

	target := c.read(x, y)

	// stop unnecessary fill-ing
	if target == fillChar {
		return
	}

	c.floodFill(x, y, target, fillChar)
}

// floodFill draws the fill character to the start coordinate, and
// continues to attempt drawing the character around (up, down, left, right)
// in each direction from the position it was drawn at, as long as a different
// character, or a border of the Canvas, is not reached.
func (c *Canvas) floodFill(x, y int, targetChar, fillChar byte) {

	// check canvas bounds
	boundX, boundY := c.boundary()

	if !inBounds(x, y, boundX, boundY) {
		return
	}

	cur := c.read(x, y)

	// stop if it is a different character
	// avoid stack overflow by not revisiting
	if cur != targetChar || cur == fillChar {
		return
	}

	// fill start coordinate
	c.write(x, y, fillChar)

	// draw around start point
	c.floodFill(x, y-1, targetChar, fillChar) // up
	c.floodFill(x, y+1, targetChar, fillChar) // down
	c.floodFill(x-1, y, targetChar, fillChar) // left
	c.floodFill(x+1, y, targetChar, fillChar) // right
}

// Print returns the string representation of the Canvas
func (c *Canvas) Print() string {

	var hasBlank bool
	var buf bytes.Buffer

	boundX, boundY := c.boundary()

	// read all cells
	for y := 0; y < boundY; y++ {
		for x := 0; x < boundX; x++ {

			cell := c.read(x, y)

			// skip contiguous blank cells
			if cell == byte(0) {
				hasBlank = true
				continue
			}

			// hasBlank and current cell isn't blank, add a space
			if hasBlank {
				hasBlank = false // reset

				// at least buffer isn't empty
				if buf.Len() > 0 {
					buf.WriteString(" ")
				}
			}

			buf.WriteByte(cell)
		}

		// not last row and doesn't have blank spaces before
		if !hasBlank && y != boundY-1 {
			buf.WriteString(" ")
		}
	}

	return buf.String()
}

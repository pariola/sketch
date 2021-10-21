package canvas

import (
	"bytes"
)

// canvas
type canvas struct {

	// 2D array to represent the points on a canvas starting from (0,0) top left
	matrix [][]string
}

// New returns an instance of the canvas with the required size
func New(width, height int) *canvas {

	matrix := make([][]string, height)

	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]string, width)
	}

	return &canvas{
		matrix: matrix,
	}
}

// boundary returns the edges/boundaries (x-axis, y-axis) of the canvas
func (c *canvas) boundary() (int, int) {

	var x, y int

	y = len(c.matrix)

	if y > 0 {
		x = len(c.matrix[0])
	}

	return x, y
}

// expand increases the size of the Canvas to specified size if current size is smaller
func (c *canvas) expand(x, y int) {

	boundX, boundY := c.boundary()

	// expand y-axis
	if y > boundY {

		// add blank rows
		for i := 0; i < (y - boundY); i++ {
			c.matrix = append(c.matrix, make([]string, boundX))
		}

		boundY = y
	}

	// expand x-axis
	if x > boundX {
		boundX = x

		// expand all rows by copying
		for i := 0; i < boundY; i++ {
			row := make([]string, boundX)
			copy(row, c.matrix[i])
			c.matrix[i] = row
		}
	}
}

// Draw draws a rectangle unto the canvas
func (c *canvas) Draw(r Rectangle) {

	// expand canvas to contain rectangle
	c.expand(r.posX+r.width, r.posY+r.height)

	for y := r.posY; y < r.posY+r.height; y++ {
		for x := r.posX; x < r.posX+r.width; x++ {

			// fill
			c.matrix[y][x] = r.fillChar

			// paint outline
			// confirm if point is part of the rectangle boundary
			if r.outlineChar != "" && (inYBoundary(x, y, r) || inXBoundary(x, y, r)) {
				c.matrix[y][x] = r.outlineChar
			}
		}
	}
}

// Print returns the string representation of the canvas
func (c *canvas) Print() string {

	var buf bytes.Buffer

	boundX, boundY := c.boundary()

	// read all cells
	for y := 0; y < boundY; y++ {
		for x := 0; x < boundX; x++ {
			cell := c.matrix[y][x]
			buf.WriteString(cell)
		}
		buf.WriteString(" ")
	}

	return buf.String()
}

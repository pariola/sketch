package canvas

import (
	"bytes"
)

// canvas
type canvas struct {

	// 2D array
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

// Draw draws a rectangle unto the canvas
func (c *canvas) Draw(r Rectangle) {

	// todo: expand canvas if rectangle out of bounds

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

func inXBoundary(x, y int, r Rectangle) bool {
	return x >= r.posX && x < r.posX+r.width && (y == r.posY || y == r.posY+r.height-1)
}

func inYBoundary(x, y int, r Rectangle) bool {
	return y >= r.posY && y < r.posY+r.height && (x == r.posX || x == r.posX+r.width-1)
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

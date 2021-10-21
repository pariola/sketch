package canvas

// canvas
type canvas struct {

	// 2D array
	matrix [][]string
}

// New returns an instance of the canvas
func New() *canvas {
	return &canvas{
		matrix: make([][]string, 0),
	}
}

// edges returns the edges of x & y-axis
func (c *canvas) edges() (int, int) {

	var x, y int

	y = len(c.matrix)

	if y > 0 {
		x = len(c.matrix[0])
	}

	return x, y
}

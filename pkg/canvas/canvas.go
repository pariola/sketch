package canvas

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

// edges returns the edges of x & y-axis
func (c *canvas) edges() (int, int) {

	var x, y int

	y = len(c.matrix)

	if y > 0 {
		x = len(c.matrix[0])
	}

	return x, y
}

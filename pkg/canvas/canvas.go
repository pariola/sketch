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

package canvas

// Rectangle represents a rectangle to be drawn
type Rectangle struct {

	// posX, posY represents the upper left coordinates to draw the Rectangle
	posX, posY int

	// width, height represents the dimensions of the Rectangle
	width, height int

	// fillChar represents an optional fill character
	fillChar byte

	// outlineChar represents an optional outline character
	outlineChar byte
}

// NewRectangle returns an instance of Rectangle with the specified parameters
func NewRectangle(posX, posY, width, height int, fill, outline byte) *Rectangle {

	if posX < 0 || posY < 0 || width < 0 || height < 0 {
		return nil
	}

	// one of either fill or outline should always be present
	if outline == empty && fill == empty {
		return nil
	}

	return &Rectangle{
		posX:        posX,
		posY:        posY,
		width:       width,
		height:      height,
		fillChar:    fill,
		outlineChar: outline,
	}
}

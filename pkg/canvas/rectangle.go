package canvas

// Rectangle represents a rectangle to be drawn
type Rectangle struct {

	// posX, posY represents the upper left coordinates to draw the Rectangle
	posX, posY int

	// width, height represents the dimensions of the Rectangle
	width, height int

	// todo: use rune instead
	// fillChar represents an optional fill character
	fillChar string

	// todo: use rune instead
	// outlineChar represents an optional outline character
	outlineChar string
}

// NewRectangle returns an instance of Rectangle with the specified parameters
func NewRectangle(posX, posY, width, height int, fill, outline string) *Rectangle {

	if posX < 0 || posY < 0 || width < 0 || height < 0 {
		return nil
	}

	// one of either fill or outline should always be present
	if outline == "" && fill == "" {
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

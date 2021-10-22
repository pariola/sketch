package canvas

// inBounds checks if position (x,y) is within bounds
func inBounds(x, y, boundX, boundY int) bool {
	return !(x < 0 || x >= boundX || y < 0 || y >= boundY)
}

// inXBoundary checks if position (x,y) is on the X-axis boundary of the rectangle
func inXBoundary(x, y int, r Rectangle) bool {
	return x >= r.posX && x < r.posX+r.width && (y == r.posY || y == r.posY+r.height-1)
}

// inYBoundary checks if position (x,y) is on the Y-axis boundary of the rectangle
func inYBoundary(x, y int, r Rectangle) bool {
	return y >= r.posY && y < r.posY+r.height && (x == r.posX || x == r.posX+r.width-1)
}

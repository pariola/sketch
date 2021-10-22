package service

// DrawRectangleRequest represents a draw rectangle operation on the canvas
type DrawRectangleRequest struct {
	PosX int `json:"pos_x"`
	PosY int `json:"pos_y"`

	Width  int `json:"width"`
	Height int `json:"height"`

	Fill    string `json:"fill"`
	Outline string `json:"outline"`
}

// FloodFillRequest represents a flood fill operation on the canvas
type FloodFillRequest struct {
	PosX int `json:"pos_x"`
	PosY int `json:"pos_y"`

	Fill string `json:"fill"`
}

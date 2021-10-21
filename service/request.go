package service

// DrawRequest represents a draw rectangle operation on the canvas
type DrawRequest struct {
	PosX int `json:"pos_x"`
	PosY int `json:"pos_y"`

	Width  int `json:"width"`
	Height int `json:"height"`

	Fill    string `json:"fill"`
	Outline string `json:"outline"`
}
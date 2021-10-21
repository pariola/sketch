package handler

import (
	"sketch/service"
)

// SketchService is an interface that wraps the SketchService functionality
type SketchService interface {
	DrawRectangle(string, service.DrawRectangleRequest) (string, error)
}

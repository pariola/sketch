package handler

import (
	"sketch/service"
)

// SketchService is an interface that wraps the SketchService functionality
type SketchService interface {
	Draw(string, service.DrawRequest) (string, error)
}

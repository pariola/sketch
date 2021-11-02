package storage

import (
	"sketch/pkg/canvas"
)

type Store interface {
	SaveState(map[string]*canvas.Canvas) error
	LoadState() (map[string]*canvas.Canvas, error)
}

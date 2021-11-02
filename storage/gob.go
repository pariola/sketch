package storage

import (
	"encoding/gob"
	"io"
	"os"

	"sketch/pkg/canvas"
)

type gobStore struct {
	path string
}

func New(path string) Store {
	return &gobStore{path: path}
}

func (s gobStore) LoadState() (map[string]*canvas.Canvas, error) {

	f, err := os.OpenFile(s.path, os.O_CREATE|os.O_RDONLY, 0644)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	var state map[string]*canvas.Canvas

	err = gob.NewDecoder(f).Decode(&state)

	// error would be EOF for newly created stores
	if err == io.EOF {
		return make(map[string]*canvas.Canvas, 0), nil
	}

	return state, err
}

func (s gobStore) SaveState(state map[string]*canvas.Canvas) error {

	f, err := os.OpenFile(s.path, os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	defer f.Close()

	err = gob.NewEncoder(f).Encode(state)

	if err != nil {
		return err
	}

	return nil
}

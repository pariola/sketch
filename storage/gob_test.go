package storage

import (
	"os"
	"testing"

	"sketch/pkg/canvas"

	"github.com/stretchr/testify/assert"
)

func TestGobStore_LoadState(t *testing.T) {

	stateFilePath := "dummy.gob"

	// cleanup prior state file
	_ = os.Remove(stateFilePath)

	store := &gobStore{
		path: stateFilePath,
	}

	state, err := store.LoadState()

	assert.Nil(t, err, "LoadState should not fail", err)
	assert.NotNil(t, state, "state should be a non-nil map")
	assert.FileExists(t, stateFilePath, "state file should be created")

	// cleanup state file
	_ = os.Remove(stateFilePath)
}

func TestGobStore(t *testing.T) {

	stateFilePath := "dummy.gob"

	// cleanup
	defer os.Remove(stateFilePath)

	// write dummy data
	c := canvas.New(7, 5)

	// draw on canvas
	c.Draw(
		*canvas.NewRectangle(0, 1, 5, 3, 'X', '@'),
	)

	// initialize state
	state := map[string]*canvas.Canvas{
		"dummy": c,
	}

	store := &gobStore{
		path: stateFilePath,
	}

	err := store.SaveState(state)

	assert.Nil(t, err, "SaveState should not fail", err)
	assert.FileExists(t, stateFilePath, "state file should be saved")

	// load state
	loadedState, err := store.LoadState()

	assert.Nil(t, err, "LoadState should not fail", err)

	// loaded state should contain only one key "dummy"
	assert.Len(t, loadedState, 1)

	dummyCanvas, ok := loadedState["dummy"]

	if assert.True(t, ok, "loaded state should have `dummy` key") {

		// compare underlying matrix
		assert.Equal(t, c.Matrix, dummyCanvas.Matrix, "result matrix does not match initial matrix")

		// compare canvas string representations
		assert.Equal(t, c.Print(), dummyCanvas.Print(), "result print does not match initial canvas print")
	}
}

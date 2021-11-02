# ASCII Canvas

This service implements an ASCII art drawing canvas as a web service.

The `Canvas` has a default size of `28x12`, it can automatically expand when trying to draw a `Rectangle` that goes beyond its original size. 

This service persists state across launches by encoding all the canvas' state with `encoding/gob` in a file (configured as `sketch.gob`) at the current directory.

### Run

Run the command below to start the Sketch Service on port 5000.

```sh
go run main.go
```

To run tests on the project:

```sh
go test ./...
```

### Client View (read-only)

uses a websocket to provide near realtime changes to the Canvas. Should be viewed with a browser.

#### Create Canvas

`http://localhost:5000/` - creates a new canvas and redirects to the canvas view.

#### View Canvas

`http://localhost:5000/[CANVAS ID]` - displays the canvas.

### API Endpoints

_NOTE:_ URL parameter `id` refers to the ID of the canvas.

#### Create Canvas

`GET /api/canvas` Creates a new canvas then returns the ID.

Request:

```sh
curl -XGET 'http://localhost:5000/api/canvas'
```

Response:

```text
78587113-74e0-4997-852c-2b0d469e5194
```

#### Draw Rectangle

`POST /api/canvas/:id/draw` Draws a Rectangle on the canvas then prints it.

Parameters:

- `pos_x`: coordinates for the upper-left corner (X-axis)
- `pos_y`: coordinates for the upper-left corner (Y-axis)
- `width`: width of the rectangle
- `height`: height of the rectangle
- `fill`: (Optional) a character to fill the rectangle with
- `outline`: (Optional) a character to outline the rectangle with

NOTE: One of either `fill` or `outline` should always be present.

Request:

```sh
curl -XPOST -H "Content-type: application/json" -d '{
  "pos_x": 3,
  "pos_y": 2,
  "width": 5,
  "height": 3,
  "fill": "X",
  "outline": "@"
}' 'http://localhost:5000/api/canvas/78587113-74e0-4997-852c-2b0d469e5194/draw'
```

Response:

```text
@@@@@ @XXX@ @@@@@
```

#### Flood Fill

`POST /api/canvas/:id/floodfill` Performs the flood fill operation on the canvas and then prints it.

Parameters:

- `pos_x`: start coordinates (X-axis)
- `pos_y`: start coordinates (Y-axis)
- `fill`:  a character to use in the operation

Request:

```sh
curl -XPOST -H "Content-type: application/json" -d '{
  "pos_x": 4,
  "pos_y": 3,
  "fill": "-"
}' 'http://localhost:5000/api/canvas/78587113-74e0-4997-852c-2b0d469e5194/floodfill'
```

Response:

```text
@@@@@ @---@ @@@@@
```

#### Print Canvas

`GET /api/canvas/:id` Returns the content of the canvas with the referenced `id`.

Request:

```sh
curl -XGET 'http://localhost:5000/api/canvas/78587113-74e0-4997-852c-2b0d469e5194'
```

Response:

```text
@@@@@ @---@ @@@@@
```

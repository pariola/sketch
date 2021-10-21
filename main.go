package main

import (
	"log"
	"net/http"

	"sketch/pkg/canvas"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	c := canvas.New(10, 12)

	e.POST("/draw", func(ctx echo.Context) error {

		// rectangle at [2,2] width: 5, height: 5, fill: 'X', outline: none
		r := canvas.NewRectangle(2, 2, 5, 5, "X", "0")

		c.Draw(*r)

		return ctx.String(http.StatusOK, c.Print())
	})

	// start server
	log.Fatalln(e.Start(":5000"))
}

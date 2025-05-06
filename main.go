package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		file, err := os.ReadFile("./index.html")
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error reading file")
		}
		return c.HTML(http.StatusOK, string(file))
	})
	e.Logger.Fatal(e.Start(":8080"))
}

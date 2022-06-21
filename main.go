package main

import (
	"gozzafadillah/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.POST("/file", handler.FileUpload)
	e.POST("/remote", handler.RemoteUpload)

	e.Logger.Fatal(e.Start(":8080"))
}

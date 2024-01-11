package main

import (
	"fmt"

	"github.com/jschue85/bujo-command/htmlrenderer"
	"github.com/jschue85/bujo-command/viewmodels"
	"github.com/labstack/echo/v4"
)

const PORT = 8080
const EXTERNAL_LOOPBACK = "172.30.32.1"

func main() {
	tr := &htmlrenderer.Template{
		Templates: htmlrenderer.ConfigTemplatePath("public/views/*.html"),
	}
	e := echo.New()

	e.POST("/", viewmodels.SelectJournal)

	e.Renderer = tr

	e.GET("/", viewmodels.Index)
	e.GET("/daily", viewmodels.Daily)

	fmt.Printf("https://%s:%d", EXTERNAL_LOOPBACK, PORT)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", PORT)))
}

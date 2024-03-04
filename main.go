package main

import (
	"fmt"

	"github.com/jschue85/bujo-command/htmlrenderer"
	"github.com/jschue85/bujo-command/viewmodels"
	"github.com/labstack/echo/v4"
)

const PORT = 8080
const EXTERNAL_LOOPBACK = "172.21.144.1"

func main() {
	fmt.Println("Starting...")
	tr := &htmlrenderer.Template{
		Templates: htmlrenderer.ConfigTemplatePath("public/views/*.html"),
	}
	e := echo.New()
	e.Static("/static", "public/views/static")
	e.Renderer = tr
	e.GET("/journal/:id", viewmodels.Journal)
	e.POST("/journal", viewmodels.AddJournal)
	e.DELETE("/journal/:id", viewmodels.DeleteJournal)
	e.GET("/journal/:id/daily/:month/:day", viewmodels.Daily)
	e.GET("/login", viewmodels.LoginForm)
	e.POST("/login", viewmodels.Login)
	e.GET("/", viewmodels.Index)
	//e.GET("/daily", viewmodels.Daily)

	fmt.Printf("http://%s:%d", EXTERNAL_LOOPBACK, PORT)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", PORT)))
}

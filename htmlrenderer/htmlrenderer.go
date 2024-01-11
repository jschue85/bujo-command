package htmlrenderer

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

func ConfigTemplatePath(paths string) *template.Template {
	return template.Must(template.ParseGlob(paths))
}

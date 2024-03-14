package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Count struct {
	Count int
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	count := Count{Count: 0}
	e.Renderer = NewTemplates()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index.html", count)
	})
	e.POST("/count", func(c echo.Context) error {
		count.Count++
		return c.Render(200, "index.html", count)
	})
	e.Logger.Fatal(e.Start(":42069"))
}

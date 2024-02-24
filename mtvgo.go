
package main

import (
	// "database/sql"
	"flag"
	// "fmt"
	"html/template"
	"io"
	"net/http"

	
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

type Template struct {
	templates *template.Template
}



func init() {
	godotenv.Load("mtvgo.env")
}

func main() {
	p := flag.String("port", "80", "Port to run on")
	port := ":" + *p
	flag.Parse()

	
		e := echo.New()
		e.Use(middleware.CORS())
		e.Use(middleware.Gzip())
		e.Use(middleware.Recover())
		t := &Template{
			templates: template.Must(template.ParseGlob("AtsTemplates/*")),
		}
		e.Renderer = t

		e.GET("/", index)
		e.GET("/tv", tv)
		// e.GET("/comments", ats_comments)
		// e.GET("/estimates", ats_estimates)
		// e.GET("/images", ats_images)
		// e.GET("/services", ats_services)
		
		e.Static("/assets", "assets")
		e.Logger.Fatal(e.Start(port))
	
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "WORKED")
}

func tv(c echo.Context) error {
	return c.Render(http.StatusOK, "tv", "WORKED")
}






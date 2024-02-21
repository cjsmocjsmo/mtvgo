
package mtvgosetup

import (
	// "crypto/tls"
	// "golang.org/x/crypto/acme"
	// "database/sql"
	"flag"
	// "fmt"
	"html/template"
	"io"
	// "log"
	// "mime/multipart"
	"net/http"
	// "net/mail"
	// "os"
	// "regexp"
	// "strconv"
	// "time"
	// "unicode"
	// "strings"
	

	// "github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "github.com/mailjet/mailjet-apiv3-go/v4"
	_ "github.com/mattn/go-sqlite3"
	// "golang.org/x/crypto/acme/autocert"
	// "github.com/cjsmocjsmo/mtvgosetup"
)

type Template struct {
	templates *template.Template
}

// func createMoviesDB(db_path string) {
// 	db, err := sql.Open("sqlite3", db_path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
// 	sqlStmt := `
// 	CREATE TABLE IF NOT EXISTS movies (
// 		id INTEGER PRIMARY KEY,
//             name TEXT NOT NULL,
//             year TEXT NOT NULL,
//             posteraddr TEXT NOT NULL,
//             size TEXT NOT NULL,
//             path TEXT NOT NULL,
//             idx TEXT NOT NULL,
//             movid TEXT NOT NULL UNIQUE,
//             catagory TEXT NOT NULL,
//             httpthumbpath TEXT NOT NULL
// 	);
// 	`
// 	_, err = db.Exec(sqlStmt)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func createTVShowsDB(db_path string) {
// 	db, err := sql.Open("sqlite3", db_path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
// 	sqlStmt := `
// 	CREATE TABLE IF NOT EXISTS tvshows (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
//             tvid TEXT NOT NULL UNIQUE,
//             size TEXT NOT NULL,
//             catagory TEXT NOT NULL,
//             name TEXT NOT NULL,
//             season TEXT NOT NULL,
//             episode TEXT NOT NULL,
//             path TEXT NOT NULL,
//             idx TEXT NOT NULL
// 	);
// 	`
// 	_, err = db.Exec(sqlStmt)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func createImagesDB(db_path string) {
// 	db, err := sql.Open("sqlite3", db_path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
// 	sqlStmt := `
// 	CREATE TABLE IF NOT EXISTS images (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
//             imgid TEXT NOT NULL UNIQUE,
//             path TEXT NOT NULL,
//             imgpath TEXT NOT NULL,
//             size TEXT NOT NULL,
//             name TEXT NOT NULL,
//             thumbpath TEXT NOT NULL,
//             idx INTEGER NOT NULL,
//             httpthumbpath TEXT NOT NULL
// 	);
// 	`
// 	_, err = db.Exec(sqlStmt)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func init() {
	godotenv.Load("mtvgo.env")
	MtvGoSetup()
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

		e.GET("/", ats_index)
		// e.GET("/about", ats_about)
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

func ats_index(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_index", "WORKED")
}





package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/codegangsta/martini"
	"github.com/pilu/fresh/runner/runnerutils"
)

func runnerMiddleware(w http.ResponseWriter, r *http.Request) {
	if runnerutils.HasErrors() {
		runnerutils.RenderError(w)
	}
}

func main() {
	m := martini.Classic()

	if os.Getenv("MARTINI_ENV") != "production" {
		m.Use(runnerMiddleware)
	}
	m.Use(martini.Static("public/app"))

	m.Get("/", func(r http.ResponseWriter) {
		t, err := template.ParseFiles("public/app/index.html")
		if err != nil {
			panic(err.Error())
		}
		t.Execute(r, nil)
	})
	m.Run()
}

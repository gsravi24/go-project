package main

import (
	"net/http"
	"github.com/ravindrags24/sto-repo/go-project/pkg/render"
)

// Home is the home page handler.
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler.
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

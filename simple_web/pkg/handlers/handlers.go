package handlers

import (
	"github.com/keshavrkaranth/simple_web/pkg/render"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

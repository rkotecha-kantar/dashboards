package controllers

import (
	"io/fs"
	"net/http"

	"github.com/radha-kantar/dashboard/backend/internal/models"
)

type PageController struct {
	Store    *models.WidgetStore
	Dist     fs.FS
	Renderer *Renderer
}

func (c *PageController) Register(mux *http.ServeMux) {
	mux.Handle("GET /assets/", http.FileServerFS(c.Dist))
	mux.HandleFunc("GET /{$}", c.dashboard)
}

func (c *PageController) dashboard(w http.ResponseWriter, r *http.Request) {
	c.Renderer.Render(w, "Dashboard", Props{
		"widgets":  c.Store.All(),
		"userName": "Asha",
		"appName":  "Dashboard",
	})
}
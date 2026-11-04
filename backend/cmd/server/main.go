package main

import (
	"io/fs"
	"log"
	"net/http"

	"github.com/radha-kantar/dashboard/backend/internal/controllers"
	"github.com/radha-kantar/dashboard/backend/internal/models"
	"github.com/radha-kantar/dashboard/backend/internal/web"
)

func main() {
	dist, err := fs.Sub(web.Dist, "dist")
	if err != nil {
		log.Fatal(err)
	}
	
	store := models.NewWidgetStore()
	mux := http.NewServeMux()

	(&controllers.WidgetController{Store: store}).Register(mux)
	(&controllers.PageController{
		Store:    store,
		Dist:     dist,
		Renderer: &controllers.Renderer{Dist: dist},
	}).Register(mux)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
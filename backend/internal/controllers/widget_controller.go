package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/radha-kantar/dashboard/backend/internal/models"
)

type WidgetController struct {
	Store *models.WidgetStore
}

func (c *WidgetController) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/widgets", c.list)
	mux.HandleFunc("POST /api/widgets", c.create)
	mux.HandleFunc("DELETE /api/widgets/{id}", c.remove)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func (c *WidgetController) list(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, c.Store.All())
}

func (c *WidgetController) create(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Title string `json:"title"`
		Value int    `json:"value"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil || strings.TrimSpace(in.Title) == "" {
		http.Error(w, "title required", http.StatusBadRequest)
		return
	}
	writeJSON(w, http.StatusCreated, c.Store.Create(in.Title, in.Value))
}

func (c *WidgetController) remove(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "bad id", http.StatusBadRequest)
		return
	}
	if !c.Store.Delete(id) {
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
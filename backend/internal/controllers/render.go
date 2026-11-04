package controllers

import (
	"encoding/json"
	"io/fs"
	"net/http"
	"strings"
)

type Props map[string]any

type Renderer struct {
	Dist fs.FS
}

func (rd *Renderer) Render(w http.ResponseWriter, component string, props Props) {
	index, err := fs.ReadFile(rd.Dist, "index.html")
	if err != nil {
		http.Error(w, "view not built: run `make build-frontend`", http.StatusInternalServerError)
		return
	}

	payload, err := json.Marshal(map[string]any{
		"component": component,
		"props":     props,
	})
	if err != nil {
		http.Error(w, "could not encode props", http.StatusInternalServerError)
		return
	}

	// json.Marshal escapes <, > and &, so the payload cannot break out of the script tag
	page := strings.Replace(string(index), "__PAGE_DATA__", string(payload), 1)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(page))
}
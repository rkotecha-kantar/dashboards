package models

import "sync"

type Widget struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Value int    `json:"value"`
}

type WidgetStore struct {
	mu      sync.Mutex
	nextID  int
	widgets []Widget
}

func NewWidgetStore() *WidgetStore {
	s := &WidgetStore{nextID: 1}
	s.Create("Active users", 1280)
	s.Create("Open tickets", 42)
	s.Create("Revenue (k)", 96)
	return s
}

func (s *WidgetStore) All() []Widget {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]Widget, len(s.widgets))
	copy(out, s.widgets)
	return out
}

func (s *WidgetStore) Create(title string, value int) Widget {
	s.mu.Lock()
	defer s.mu.Unlock()
	w := Widget{ID: s.nextID, Title: title, Value: value}
	s.nextID++
	s.widgets = append(s.widgets, w)
	return w
}

func (s *WidgetStore) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.widgets {
		if s.widgets[i].ID == id {
			s.widgets = append(s.widgets[:i], s.widgets[i+1:]...)
			return true
		}
	}
	return false
}
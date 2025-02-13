package server

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (s *Server) SearchBookByIDHandler(w http.ResponseWriter, r *http.Request){
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 || pathParts[2] == "" {
		http.Error(w, "Book ID is required", http.StatusBadRequest)
		return
	}
	id := pathParts[2]

	book, err := s.client.GetBookByID(id)
	if err != nil {
		http.Error(w, "Bad API call", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Bad API call", http.StatusConflict)
	}

	status, err := w.Write(data)
	if err != nil {
		http.Error(w, "Bad API call", status)
	}
	w.WriteHeader(status)
}
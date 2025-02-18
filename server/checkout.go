package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"Github.com/Ekaloi/library/db"
)

func (s *Server) CheckoutHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	id := parts[2]
	book, err := s.client.GetBookByID(id)
	if err != nil {
		http.Error(w, "Bad API call", http.StatusBadRequest)
		return
	}

	db.InsertBook(s.dynamo, book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusBadRequest)
		return
	}
	w.Write(data)
}

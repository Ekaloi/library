package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) SearchBookHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	query := fmt.Sprint(title)

	books, err := s.client.SearchBook(query)
	if err != nil {
		http.Error(w, `{"error": "server err"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(books)
	if err != nil {
		http.Error(w, `{"error": "server err"}`, http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

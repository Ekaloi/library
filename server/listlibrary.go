package server

import (
	"encoding/json"
	"net/http"

	"Github.com/Ekaloi/library/db"
)

func (s *Server) ListCheckedOutBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := db.GetBooks(s.dynamo)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	data, err := json.Marshal(books)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

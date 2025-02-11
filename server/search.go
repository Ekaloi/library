package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func (s *Server) SearchBookHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Query()["title"]
	query := fmt.Sprint(title)

	books, err := s.client.SearchBook(query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"error": "server err"})
		return
	}

	w.Header().Set("Content-Type","application/json")

	data, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"error": "server err"})
		return
	}
	w.WriteHeader(200)
	w.Write(data)
}

package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func (c *Client) GetBookByID(id string) (BookItem, error) {

	googleapiKey := os.Getenv("google_api_key")
	url := baseURL + "/" + id
	url = url + "?key=" + googleapiKey
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return BookItem{}, fmt.Errorf("error req %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return BookItem{}, fmt.Errorf("error server %w", err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var book BookItem
	err = decoder.Decode(&book)
	if err != nil {
		return BookItem{}, fmt.Errorf("error return resp %w", err)
	}

	return book, nil
}

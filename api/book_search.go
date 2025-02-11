package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func (c *Client) SearchBook(query string) (BookResponse, error) {
	googleapiKey := os.Getenv("google_api_key")
	fields := strings.Fields(query)
	url := baseURL + "?q="

	for i, v := range fields {
		if i == 0 {
			url = url + v
		} else {
			url = url + "+" + v
		}

	}
	url = url + "&key=" + googleapiKey
	req, err := http.NewRequest("GET", url, nil)
	fmt.Println(url)
	if err != nil {
		return BookResponse{}, fmt.Errorf("error with the request %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return BookResponse{}, fmt.Errorf("error with the api call %w", err)
	}

	bookResp := BookResponse{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&bookResp)

	if err != nil {
		return BookResponse{}, fmt.Errorf("decoding resp %w", err)
	}

	return bookResp, nil

}

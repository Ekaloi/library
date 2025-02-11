package api

import "net/http"

type Client struct {
	httpClient http.Client
}

const baseURL string = "https://www.googleapis.com/books/v1/volumes"

func NewClient() *Client {
	return &Client{httpClient: http.Client{}}
}
 
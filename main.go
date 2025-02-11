package main

import (
	"fmt"
	"log"
	"Github.com/Ekaloi/library/api"
	"github.com/joho/godotenv"
)

func main(){
	if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

	query := "Dresden Files"
	client := api.NewClient()

	bookResp, err := client.SearchBook(query)
	if err != nil {
		fmt.Printf("Error %v", err)
	}

	for i := 0; i < 10 ; i++{
		fmt.Printf("Title %s\n", bookResp.Items[i].VolumeInfo.Title)
	}
	
}
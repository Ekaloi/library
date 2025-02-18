package db

import (
	"context"
	"time"

	"Github.com/Ekaloi/library/api"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)


func InsertBook(client *dynamodb.Client, book api.BookItem) error {
	var data Book = Book{
		BookID: book.ID,
		Title: book.VolumeInfo.Title,
		Author: book.VolumeInfo.Authors[0],
		Genre: book.VolumeInfo.MaturityRating,
		PublishedYear: book.VolumeInfo.PublishedDate,
		ReadStatus: true,
		DateAdded: time.Now().GoString(),
	}

	item, err := attributevalue.MarshalMap(data)
	if err != nil {
		return err
	}

	_, err = client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("Books"),
		Item:      item,
	})
	return err
}

func GetBook(client *dynamodb.Client, bookID string) (*Book, error) {
	out, err := client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("Books"),
		Key: map[string]types.AttributeValue{
			"BookID": &types.AttributeValueMemberS{Value: bookID},
		},
	})
	if err != nil {
		return nil, err
	}

	var book Book
	err = attributevalue.UnmarshalMap(out.Item, &book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func GetBooks(client *dynamodb.Client)([]Book, error){
	out, err := client.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("Books"),
	})
	if err != nil {
		return nil, err
	}

	var books []Book
	err = attributevalue.UnmarshalListOfMaps(out.Items,&books)
	if err != nil {
		return nil, err
	}

	return books, nil
}


type Book struct {
    BookID        string `json:"bookid"`
    Title         string `json:"title"`
    Author        string `json:"author"`
    Genre         string `json:"genre"`
    PublishedYear string    `json:"publishedYear"`
    ReadStatus    bool   `json:"readStatus"`
    DateAdded     string `json:"dateAdded"`
}

package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/nurislam03/Bookshop/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// BookDataStore ...
type BookDataStore struct {
	URL string
}

// GetAllBooks ...
func (store *BookDataStore) GetAllBooks() ([]*model.Book, error) {
	// setting timer for this request to make this call active
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connecting database
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(store.URL))
	if err != nil {
		return nil, err
	}

	// getting the collection from the connected database
	collection := client.Database("bookshop").Collection("books")

	// book model for result value
	var resBook []*model.Book

	// setting cursor to expected collection
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal("Error on Finding all Book documents", err)
	}

	// decoding each book and putting them into result book array
	for cur.Next(ctx) {
		var tempBook model.Book
		err = cur.Decode(&tempBook)
		if err != nil {
			log.Fatal("Error on Decoding Book document", err)
		}
		resBook = append(resBook, &tempBook)
	}

	log.Println("connecting db using ", store.URL)
	return resBook, err
}

// GetBookByID ...
func (store *BookDataStore) GetBookByID(id string) (*model.Book, error) {
	// setting timer for this request to make this call active
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connecting database
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(store.URL))
	if err != nil {
		return nil, err
	}

	// getting the collection from the connected database
	collection := client.Database("bookshop").Collection("books")

	// book model for result value
	var resBook *model.Book

	filter := bson.M{"id": id}

	documentReturned := collection.FindOne(ctx, filter)
	err = documentReturned.Decode(&resBook)
	if err != nil {
		log.Fatal("Error on Decoding a single Book document", err)
	}
	log.Println("connecting db using ", store.URL)
	return resBook, err
}

// AddNewBook ...
func (store *BookDataStore) AddNewBook(book *model.Book) (*model.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(store.URL))
	if err != nil {
		return nil, err
	}
	client.Database("bookshop").Collection("books").InsertOne(ctx, book)
	log.Println("connecting db using ", store.URL)
	return book, nil
}

// UpdateBookByID ...
func (store *BookDataStore) UpdateBookByID(id string, book *model.Book) (*model.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(store.URL))
	if err != nil {
		return nil, err
	}

	filter := bson.M{"id": book.ID}
	newData := bson.D{{Key: "$set", Value: book}}

	collection := client.Database("bookshop").Collection("books")

	// updatedResult, err := collection.UpdateOne(ctx, filter, newData)
	_, err = collection.UpdateOne(ctx, filter, newData)
	if err != nil {
		log.Fatal("Error on updating one Hero", err)
	}

	// var resBook *model.Book
	// err = updatedResult.Decode(&resBook)
	// if err != nil {
	// 	log.Fatal("Error on Decoding a single Book document", err)
	// }

	log.Println("connecting db using ", store.URL)
	return book, err
}

// DeleteBookByID ...
func (store *BookDataStore) DeleteBookByID(id string) error {
	// setting timer for this request to make this call active
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connecting database
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(store.URL))
	if err != nil {
		return err
	}

	// getting the collection from the connected database
	collection := client.Database("bookshop").Collection("books")

	filter := bson.M{"id": id}
	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal("Error on deleting one book", err)
	}
	// return deleteResult, err
	return err
}

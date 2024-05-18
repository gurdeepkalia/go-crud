package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db = getMongoClient()

/*
here getMongoClient() would have been called even before init(). Order goes like this, variables initialise -> init() -> main()
Additionally, you can have multiple init() functions per package; they will be executed in the order they show up in the file
(after all variables are initialized of course). If they span multiple files, they will be executed in lexical file name order
 func init() {
 	err := godotenv.Load()
 	if err != nil {
 		log.Fatal(err)
 	}
 }
*/

func GetCollection(dbName string, collectionName string) *mongo.Collection {
	return db.Database(dbName).Collection(collectionName)
}

func getMongoClient() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_CONNECTION_STRING")))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to mongodb")
	return client
}

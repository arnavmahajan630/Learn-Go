package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
  serverAPI := options.ServerAPI(options.ServerAPIVersion1)
  dbUrl := os.Getenv("DATABASE_URL")
  if(dbUrl == "") {
	log.Fatal("Error Fetching The Database URL")
  }
  opts := options.Client().ApplyURI(dbUrl).SetServerAPIOptions(serverAPI)

  client, err := mongo.Connect(context.Background(), opts)
  if err != nil {
    panic(err)
  }

  defer func() {
    if err = client.Disconnect(context.TODO()); err != nil {
      panic(err)
    }
  }()

  // Send a ping to confirm a successful connection
  if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
    panic(err)
  }
  log.Println("Pinged the Deployment. Successfully connected to MongoDB!")
}

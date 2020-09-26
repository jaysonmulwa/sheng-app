package repository

import (
	
	"log"
	"os"
	"context"
	"fmt"
	"time"
	"github.com/jaysonmulwa/sheng-app/graph/model"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WordRepository interface{
	Save(word *model.Word)
	FindAll() []model.Word
}

type database struct {
	client *mongo.Client
}

const (
	DATABASE = "library"
	COLLECTION = "words"
)

func Mongo_Key() string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	api_key := os.Getenv("MONGODB")

	return api_key

}

func New() WordRepository {

	//mongodb+srv://USERNAME:PASSWORD@HOST:PORT
	//MONGODB := OS.Getenv("MONGODB")
	//MONGODB:= "mongodb+srv://jay:jayjay123@sheng.v3ijf.mongodb.net/<dbname>?retryWrites=true&w=majority"
	
	clientOptions := options.Client().ApplyURI(Mongo_Key())
	clientOptions = clientOptions.SetMaxPoolSize(50)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	dbClient, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}


	fmt.Println("Conected toMongoDB")


	return &database{
		client: dbClient,
	}
}

func (db *database) Save(word *model.Word)  {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	_, err := collection.InsertOne(context.TODO(), word)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *database) FindAll() []*model.Word  {

	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.TODO())

	var result []*model.Word
	for cursor.Next(context.TODO()){
		var v *model.Word
		err := cursor.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, v)

	}

	return result
	
}
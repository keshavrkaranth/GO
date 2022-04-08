package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/keshavrkaranth/mongoDb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

var connectionString = "mongodb+srv://keshavrkaranth:gF5q5VzlVXN2Opcr@cluster0.ei3ug.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
var dbName = "netflix"
var colName = "watchList"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	connect, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo DB connection Success")
	collection = connect.Database(dbName).Collection(colName)
	fmt.Println("Collection is ready", collection.Name())

}

func insertOneMovie(movie model.Netflix) {
	success, err := collection.InsertOne(context.Background(), movie)
	fmt.Println("MOVIE", movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie with ID", success.InsertedID)

}

func updateOneMovie(movie string) {
	id, _ := primitive.ObjectIDFromHex(movie)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	success, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count", success.ModifiedCount)
}

func deleteOneRecord(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted movie delete count", deleteCount)
}

func deleteManyRecord() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of movies deleted", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

func getAllMovies() []primitive.M {
	curr, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
	for curr.Next(context.Background()) {
		var movie bson.M
		err := curr.Decode(&movies)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)

	}
	defer func(curr *mongo.Cursor, ctx context.Context) {
		err := curr.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(curr, context.Background())
	return movies
}

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	_ = json.NewEncoder(w).Encode(allMovies)

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	fmt.Println("DATA", r.Body)
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	_ = json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	_ = json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOneRecord(params["id"])
	_ = json.NewEncoder(w).Encode(params["id"])

}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	count := deleteManyRecord()
	_ = json.NewEncoder(w).Encode(count)

}
